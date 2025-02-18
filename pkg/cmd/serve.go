package cmd

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"

	"github.com/sony/gobreaker"
	"go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/trace"
	"golang.org/x/sync/errgroup"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/tolgaOzen/go-skeleton/internal"
	"github.com/tolgaOzen/go-skeleton/internal/config"
	"github.com/tolgaOzen/go-skeleton/internal/factories"
	"github.com/tolgaOzen/go-skeleton/internal/servers"
	"github.com/tolgaOzen/go-skeleton/internal/storage"
	"github.com/tolgaOzen/go-skeleton/internal/storage/decorators/circuitBreaker"
	"github.com/tolgaOzen/go-skeleton/pkg/cmd/flags"
	"github.com/tolgaOzen/go-skeleton/pkg/telemetry"
	"github.com/tolgaOzen/go-skeleton/pkg/telemetry/meterexporters"
	"github.com/tolgaOzen/go-skeleton/pkg/telemetry/tracerexporters"
)

func NewServeCommand() *cobra.Command {
	command := &cobra.Command{
		Use:   "serve",
		Short: "serve the server",
		RunE:  serve(),
		Args:  cobra.NoArgs,
	}

	conf := config.DefaultConfig()
	f := command.Flags()
	f.StringP("config", "c", "", "config file (default is $HOME/.config.yaml)")
	f.Bool("http-enabled", conf.Server.HTTP.Enabled, "switch option for HTTP server")
	f.Int64("server-rate-limit", conf.Server.RateLimit, "the maximum number of requests the server should handle per second")
	f.String("server-name-override", conf.Server.NameOverride, "server name override")
	f.String("grpc-port", conf.Server.GRPC.Port, "port that GRPC server run on")
	f.Bool("grpc-tls-enabled", conf.Server.GRPC.TLSConfig.Enabled, "switch option for GRPC tls server")
	f.String("grpc-tls-key-path", conf.Server.GRPC.TLSConfig.KeyPath, "GRPC tls key path")
	f.String("grpc-tls-cert-path", conf.Server.GRPC.TLSConfig.CertPath, "GRPC tls certificate path")
	f.String("http-port", conf.Server.HTTP.Port, "HTTP port address")
	f.Bool("http-tls-enabled", conf.Server.HTTP.TLSConfig.Enabled, "switch option for HTTP tls server")
	f.String("http-tls-key-path", conf.Server.HTTP.TLSConfig.KeyPath, "HTTP tls key path")
	f.String("http-tls-cert-path", conf.Server.HTTP.TLSConfig.CertPath, "HTTP tls certificate path")
	f.StringSlice("http-cors-allowed-origins", conf.Server.HTTP.CORSAllowedOrigins, "CORS allowed origins for http gateway")
	f.StringSlice("http-cors-allowed-headers", conf.Server.HTTP.CORSAllowedHeaders, "CORS allowed headers for http gateway")
	f.Bool("profiler-enabled", conf.Profiler.Enabled, "switch option for profiler")
	f.String("profiler-port", conf.Profiler.Port, "profiler port address")
	f.String("log-level", conf.Log.Level, "set log verbosity ('info', 'debug', 'error', 'warning')")
	f.String("log-output", conf.Log.Output, "logger output valid values json, text")
	f.Bool("authn-enabled", conf.Authn.Enabled, "enable server authentication")
	f.String("authn-method", conf.Authn.Method, "server authentication method")
	f.StringSlice("authn-preshared-keys", conf.Authn.Preshared.Keys, "preshared key/keys for server authentication")
	f.Bool("tracer-enabled", conf.Tracer.Enabled, "switch option for tracing")
	f.String("tracer-exporter", conf.Tracer.Exporter, "can be; jaeger, signoz, zipkin or otlp. (integrated tracing tools)")
	f.String("tracer-endpoint", conf.Tracer.Endpoint, "export uri for tracing data")
	f.Bool("tracer-insecure", conf.Tracer.Insecure, "use https or http for tracer data, only used for otlp exporter or signoz")
	f.String("tracer-urlpath", conf.Tracer.URLPath, "allow to set url path for otlp exporter")
	f.StringSlice("tracer-headers", conf.Tracer.Headers, "allows setting custom headers for the tracer exporter in key-value pairs")
	f.String("tracer-protocol", conf.Tracer.Protocol, "allows setting the communication protocol for the tracer exporter, with options http or grpc")
	f.Bool("meter-enabled", conf.Meter.Enabled, "switch option for metric")
	f.String("meter-exporter", conf.Meter.Exporter, "can be; otlp. (integrated metric tools)")
	f.String("meter-endpoint", conf.Meter.Endpoint, "export uri for metric data")
	f.Bool("meter-insecure", conf.Meter.Insecure, "use https or http for metric data")
	f.String("meter-urlpath", conf.Meter.URLPath, "allow to set url path for otlp exporter")
	f.StringSlice("meter-headers", conf.Meter.Headers, "allows setting custom headers for the metric exporter in key-value pairs")
	f.Int("meter-interval", conf.Meter.Interval, "allows to set metrics to be pushed in certain time interval")
	f.String("meter-protocol", conf.Meter.Protocol, "allows setting the communication protocol for the meter exporter, with options http or grpc")
	f.Bool("service-circuit-breaker", conf.Service.CircuitBreaker, "switch option for service circuit breaker")
	f.String("database-engine", conf.Database.Engine, "data source. e.g. postgres, memory")
	f.String("database-uri", conf.Database.URI, "uri of your data source to store relation tuples and schema")
	f.String("database-writer-uri", conf.Database.Writer.URI, "writer uri of your data source to store relation tuples and schema")
	f.String("database-reader-uri", conf.Database.Reader.URI, "reader uri of your data source to store relation tuples and schema")
	f.Bool("database-auto-migrate", conf.Database.AutoMigrate, "auto migrate database tables")
	f.Int("database-max-open-connections", conf.Database.MaxOpenConnections, "maximum number of parallel connections that can be made to the database at any time")
	f.Int("database-max-idle-connections", conf.Database.MaxIdleConnections, "maximum number of idle connections that can be made to the database at any time")
	f.Duration("database-max-connection-lifetime", conf.Database.MaxConnectionLifetime, "maximum amount of time a connection may be reused")
	f.Duration("database-max-connection-idle-time", conf.Database.MaxConnectionIdleTime, "maximum amount of time a connection may be idle")

	// SilenceUsage is set to true to suppress usage when an error occurs
	command.SilenceUsage = true

	command.PreRun = func(cmd *cobra.Command, args []string) {
		flags.RegisterServeFlags(f)
	}

	return command
}

func serve() func(cmd *cobra.Command, args []string) error {
	return func(cmd *cobra.Command, args []string) error {
		var cfg *config.Config
		var err error
		cfgFile := viper.GetString("config.file")
		if cfgFile != "" {
			cfg, err = config.NewConfigWithFile(cfgFile)
			if err != nil {
				return fmt.Errorf("failed to create new config: %w", err)
			}

			if err = viper.Unmarshal(cfg); err != nil {
				return fmt.Errorf("failed to unmarshal config: %w", err)
			}
		} else {
			// Load configuration
			cfg, err = config.NewConfig()
			if err != nil {
				return fmt.Errorf("failed to create new config: %w", err)
			}

			if err = viper.Unmarshal(cfg); err != nil {
				return fmt.Errorf("failed to unmarshal config: %w", err)
			}
		}

		// Print banner and initialize logger
		internal.PrintBanner()

		// Set up context and signal handling
		ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
		defer stop()

		var logger *slog.Logger
		var handler slog.Handler

		switch cfg.Log.Output {
		case "json":
			handler = slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
				Level: getLogLevel(cfg.Log.Level),
			})
		case "text":
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: getLogLevel(cfg.Log.Level),
			})
		default:
			handler = slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
				Level: getLogLevel(cfg.Log.Level),
			})
		}

		logger = slog.New(handler)
		slog.SetDefault(logger)

		slog.Info("🚀 starting service...")

		// Run database migration if enabled
		if cfg.Database.AutoMigrate {
			err = storage.Migrate(cfg.Database)
			if err != nil {
				slog.Error("failed to migrate database", slog.Any("error", err))
				return err
			}
		}

		// Initialize database
		db, err := factories.DatabaseFactory(cfg.Database)
		if err != nil {
			slog.Error("failed to initialize database", slog.Any("error", err))
			return err
		}
		defer func() {
			if err = db.Close(); err != nil {
				slog.Error("failed to close database", slog.Any("error", err))
			}
		}()

		// Tracing
		if cfg.Tracer.Enabled {
			headers := map[string]string{}
			for _, header := range cfg.Tracer.Headers {
				h := strings.Split(header, ":")
				if len(h) != 2 {
					return errors.New("invalid header format; expected 'key:value'")
				}
				headers[h[0]] = h[1]
			}

			var exporter trace.SpanExporter
			exporter, err = tracerexporters.ExporterFactory(
				cfg.Tracer.Exporter,
				cfg.Tracer.Endpoint,
				cfg.Tracer.Insecure,
				cfg.Tracer.URLPath,
				headers,
				cfg.Tracer.Protocol,
			)
			if err != nil {
				slog.Error(err.Error())
			}

			shutdown := telemetry.NewTracer(exporter)

			defer func() {
				if err = shutdown(ctx); err != nil {
					slog.Error(err.Error())
				}
			}()
		}

		// Meter
		if cfg.Meter.Enabled {
			headers := map[string]string{}
			for _, header := range cfg.Meter.Headers {
				h := strings.Split(header, ":")
				if len(h) != 2 {
					return errors.New("invalid header format; expected 'key:value'")
				}
				headers[h[0]] = h[1]
			}

			var exporter metric.Exporter
			exporter, err = meterexporters.ExporterFactory(
				cfg.Meter.Exporter,
				cfg.Meter.Endpoint,
				cfg.Meter.Insecure,
				cfg.Meter.URLPath,
				headers,
				cfg.Meter.Protocol,
			)
			if err != nil {
				slog.Error(err.Error())
			}

			shutdown := telemetry.NewMeter(exporter, time.Duration(cfg.Meter.Interval)*time.Second)

			defer func() {
				if err = shutdown(ctx); err != nil {
					slog.Error(err.Error())
				}
			}()
		}

		// Initialize the storage with factory methods
		dataReader := factories.DataReaderFactory(db)
		dataWriter := factories.DataWriterFactory(db)

		if cfg.Service.CircuitBreaker {
			var cb *gobreaker.CircuitBreaker
			var st gobreaker.Settings
			st.Name = "storage"
			st.ReadyToTrip = func(counts gobreaker.Counts) bool {
				failureRatio := float64(counts.TotalFailures) / float64(counts.Requests)
				return counts.Requests >= 10 && failureRatio >= 0.6
			}

			cb = gobreaker.NewCircuitBreaker(st)

			// Wrap the dataReader with circuit breaker
			dataReader = circuitBreaker.NewDataReader(dataReader, cb)
		}

		// Initialize the container which brings together multiple components such as the invoker, data readers/writers, and schema handlers.
		container := servers.NewContainer(
			dataReader,
			dataWriter,
		)

		// Create an error group with the provided context
		var g *errgroup.Group
		g, ctx = errgroup.WithContext(ctx)

		// Add the container.Run function to the error group
		g.Go(func() error {
			return container.Run(
				ctx,
				&cfg.Server,
				logger,
				&cfg.Authn,
				&cfg.Profiler,
			)
		})

		// Wait for the error group to finish and log any errors
		if err = g.Wait(); err != nil {
			slog.Error(err.Error())
		}

		return nil
	}
}

// getLogLevel converts a string representation of log level to its corresponding slog.Level value.
func getLogLevel(level string) slog.Level {
	switch level {
	case "info":
		return slog.LevelInfo // Return Info level
	case "warn":
		return slog.LevelWarn // Return Warning level
	case "error":
		return slog.LevelError // Return Error level
	case "debug":
		return slog.LevelDebug // Return Debug level
	default:
		return slog.LevelInfo // Default to Info level if unrecognized
	}
}
