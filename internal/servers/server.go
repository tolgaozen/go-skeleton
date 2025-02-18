package servers

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net"
	"net/http"
	"net/http/pprof"
	"time"

	grpcAuth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
	grpcRecovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/logging"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/ratelimit"
	grpcValidator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/rs/cors"
	"go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"
	health "google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"
	"google.golang.org/protobuf/encoding/protojson"

	"github.com/tolgaOzen/go-skeleton/internal/authn/preshared"
	"github.com/tolgaOzen/go-skeleton/internal/config"
	"github.com/tolgaOzen/go-skeleton/internal/middleware"
	"github.com/tolgaOzen/go-skeleton/internal/storage"
	grpcV1 "github.com/tolgaOzen/go-skeleton/pkg/pb/base/v1"
)

// Container is a struct that holds the invoker and various storage
// for permission-related operations. It serves as a central point of access
// for interacting with the underlying data and services.
type Container struct {
	// DataReader for reading data from storage
	DR storage.DataReader
	// DataWriter for writing data to storage
	DW storage.DataWriter
}

func NewContainer(dr storage.DataReader, dw storage.DataWriter) *Container {
	return &Container{
		DR: dr,
		DW: dw,
	}
}

func (s *Container) Run(
	ctx context.Context,
	srv *config.Server,
	logger *slog.Logger,
	authentication *config.Authn,
	profiler *config.Profiler,
) error {
	var err error

	limiter := middleware.NewRateLimiter(srv.RateLimit) // for example 1000 req/sec

	lopts := []logging.Option{
		logging.WithLogOnEvents(logging.StartCall, logging.FinishCall),
	}

	unaryInterceptors := []grpc.UnaryServerInterceptor{
		grpcValidator.UnaryServerInterceptor(),
		grpcRecovery.UnaryServerInterceptor(),
		ratelimit.UnaryServerInterceptor(limiter),
		logging.UnaryServerInterceptor(InterceptorLogger(logger), lopts...),
	}

	streamingInterceptors := []grpc.StreamServerInterceptor{
		grpcValidator.StreamServerInterceptor(),
		grpcRecovery.StreamServerInterceptor(),
		ratelimit.StreamServerInterceptor(limiter),
		logging.StreamServerInterceptor(InterceptorLogger(logger), lopts...),
	}

	// Configure authentication based on the provided method.
	// Add the appropriate interceptors to the unary and streaming interceptors.
	if authentication != nil && authentication.Enabled {
		switch authentication.Method {
		case "preshared":
			var authenticator *preshared.KeyAuthn
			authenticator, err = preshared.NewKeyAuthn(ctx, authentication.Preshared)
			if err != nil {
				return err
			}
			unaryInterceptors = append(unaryInterceptors, grpcAuth.UnaryServerInterceptor(middleware.AuthFunc(authenticator)))
			streamingInterceptors = append(streamingInterceptors, grpcAuth.StreamServerInterceptor(middleware.AuthFunc(authenticator)))
		default:
			return fmt.Errorf("unknown authentication method: '%s'", authentication.Method)
		}
	}

	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(unaryInterceptors...),
		grpc.ChainStreamInterceptor(streamingInterceptors...),
		grpc.StatsHandler(otelgrpc.NewServerHandler()),
	}

	if srv.GRPC.TLSConfig.Enabled {
		var c credentials.TransportCredentials
		c, err = credentials.NewServerTLSFromFile(srv.GRPC.TLSConfig.CertPath, srv.GRPC.TLSConfig.KeyPath)
		if err != nil {
			return err
		}
		opts = append(opts, grpc.Creds(c))
	}

	// Create a new gRPC server instance with the provided options.
	grpcServer := grpc.NewServer(opts...)

	// Register various gRPC services to the server.
	grpcV1.RegisterUserServiceServer(grpcServer, NewUserServer(s.DR, s.DW))

	// Register health check and reflection services for gRPC.
	health.RegisterHealthServer(grpcServer, NewHealthServer())
	reflection.Register(grpcServer)

	// If profiling is enabled, set up the profiler using the net/http package.
	if profiler.Enabled {
		// Create a new HTTP ServeMux to register pprof routes.
		mux := http.NewServeMux()
		mux.HandleFunc("/debug/pprof/", pprof.Index)
		mux.HandleFunc("/debug/pprof/cmdline", pprof.Cmdline)
		mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
		mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
		mux.HandleFunc("/debug/pprof/trace", pprof.Trace)

		// Run the profiler server in a separate goroutine.
		go func() {
			// Log a message indicating the profiler server's start status and port.
			slog.Info(fmt.Sprintf("🚀 profiler server successfully started: %s", profiler.Port))

			// Define the HTTP server with timeouts and the mux handler for pprof routes.
			pprofserver := &http.Server{
				Addr:         ":" + profiler.Port,
				Handler:      mux,
				ReadTimeout:  20 * time.Second,
				WriteTimeout: 20 * time.Second,
				IdleTimeout:  15 * time.Second,
			}

			// Start the profiler server.
			if err := pprofserver.ListenAndServe(); err != nil {
				// Check if the error was due to the server being closed, and log it.
				if errors.Is(err, http.ErrServerClosed) {
					slog.Error("failed to start profiler", slog.Any("error", err))
				}
			}
		}()
	}

	var lis net.Listener
	lis, err = net.Listen("tcp", ":"+srv.GRPC.Port)
	if err != nil {
		return err
	}

	// Start the gRPC server.
	go func() {
		if err := grpcServer.Serve(lis); err != nil {
			slog.Error("failed to start grpc server", slog.Any("error", err))
		}
	}()

	slog.Info(fmt.Sprintf("🚀 grpc server successfully started: %s", srv.GRPC.Port))

	var httpServer *http.Server

	// Start the optional HTTP server with CORS and optional TLS configurations.
	// Connect to the gRPC server and register the HTTP handlers for each service.
	if srv.HTTP.Enabled {
		options := []grpc.DialOption{
			grpc.WithBlock(),
			grpc.WithUnaryInterceptor(otelgrpc.UnaryClientInterceptor()),
		}
		if srv.GRPC.TLSConfig.Enabled {
			c, err := credentials.NewClientTLSFromFile(srv.GRPC.TLSConfig.CertPath, srv.NameOverride)
			if err != nil {
				return err
			}
			options = append(options, grpc.WithTransportCredentials(c))
		} else {
			options = append(options, grpc.WithTransportCredentials(insecure.NewCredentials()))
		}

		timeoutCtx, cancel := context.WithTimeout(ctx, 3*time.Second)
		defer cancel()

		conn, err := grpc.DialContext(timeoutCtx, ":"+srv.GRPC.Port, options...)
		if err != nil {
			return err
		}
		defer func() {
			if err = conn.Close(); err != nil {
				slog.Error("Failed to close gRPC connection", slog.Any("error", err))
			}
		}()

		healthClient := health.NewHealthClient(conn)
		muxOpts := []runtime.ServeMuxOption{
			runtime.WithHealthzEndpoint(healthClient),
			runtime.WithMarshalerOption(runtime.MIMEWildcard, &runtime.HTTPBodyMarshaler{
				Marshaler: &runtime.JSONPb{
					MarshalOptions: protojson.MarshalOptions{
						UseProtoNames:   true,
						EmitUnpopulated: true,
					},
					UnmarshalOptions: protojson.UnmarshalOptions{
						DiscardUnknown: true,
					},
				},
			}),
		}

		mux := runtime.NewServeMux(muxOpts...)

		if err = grpcV1.RegisterUserServiceHandler(ctx, mux, conn); err != nil {
			return err
		}

		httpServer = &http.Server{
			Addr: ":" + srv.HTTP.Port,
			Handler: cors.New(cors.Options{
				AllowCredentials: true,
				AllowedOrigins:   srv.HTTP.CORSAllowedOrigins,
				AllowedHeaders:   srv.HTTP.CORSAllowedHeaders,
				AllowedMethods: []string{
					http.MethodGet, http.MethodPost,
					http.MethodHead, http.MethodPatch, http.MethodDelete, http.MethodPut,
				},
			}).Handler(mux),
			ReadHeaderTimeout: 5 * time.Second,
		}

		// Start the HTTP server with TLS if enabled, otherwise without TLS.
		go func() {
			var err error
			if srv.HTTP.TLSConfig.Enabled {
				err = httpServer.ListenAndServeTLS(srv.HTTP.TLSConfig.CertPath, srv.HTTP.TLSConfig.KeyPath)
			} else {
				err = httpServer.ListenAndServe()
			}
			if !errors.Is(err, http.ErrServerClosed) {
				slog.Error(err.Error())
			}
		}()

		slog.Info(fmt.Sprintf("🚀 http server successfully started: %s", srv.HTTP.Port))
	}

	// Wait for the context to be canceled (e.g., due to a signal).
	<-ctx.Done()

	// Shutdown the servers gracefully.
	ctxShutdown, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	if httpServer != nil {
		if err := httpServer.Shutdown(ctxShutdown); err != nil {
			slog.Error(err.Error())
			return err
		}
	}

	// Gracefully stop the gRPC server.
	grpcServer.GracefulStop()

	slog.Info("gracefully shutting down")

	return nil
}

// InterceptorLogger adapts slog logger to interceptor logger.
func InterceptorLogger(l *slog.Logger) logging.Logger {
	return logging.LoggerFunc(func(ctx context.Context, lvl logging.Level, msg string, fields ...any) {
		l.Log(ctx, slog.Level(lvl), msg, fields...)
	})
}
