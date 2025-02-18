package flags

import (
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

// RegisterServeFlags - Define and registers CLI flags
func RegisterServeFlags(flags *pflag.FlagSet) {
	var err error

	// Config File
	if err = viper.BindPFlag("config.file", flags.Lookup("config")); err != nil {
		panic(err)
	}

	// Server
	if err = viper.BindPFlag("server.rate_limit", flags.Lookup("server-rate-limit")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.rate_limit", "SKELETON_RATE_LIMIT"); err != nil {
		panic(err)
	}

	// GRPC Server
	if err = viper.BindPFlag("server.grpc.port", flags.Lookup("grpc-port")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.grpc.port", "SKELETON_GRPC_PORT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.grpc.tls.enabled", flags.Lookup("grpc-tls-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.grpc.tls.enabled", "SKELETON_GRPC_TLS_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.grpc.tls.key", flags.Lookup("grpc-tls-key-path")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.grpc.tls.key", "SKELETON_GRPC_TLS_KEY_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.grpc.tls.cert", flags.Lookup("grpc-tls-cert-path")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.grpc.tls.cert", "SKELETON_GRPC_TLS_CERT_PATH"); err != nil {
		panic(err)
	}

	// HTTP Server
	if err = viper.BindPFlag("server.http.enabled", flags.Lookup("http-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.enabled", "SKELETON_HTTP_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.port", flags.Lookup("http-port")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.port", "SKELETON_HTTP_PORT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.tls.enabled", flags.Lookup("http-tls-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.tls.enabled", "SKELETON_HTTP_TLS_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.tls.key", flags.Lookup("http-tls-key-path")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.tls.key", "SKELETON_HTTP_TLS_KEY_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.tls.cert", flags.Lookup("http-tls-cert-path")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.tls.cert", "SKELETON_HTTP_TLS_CERT_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.cors_allowed_origins", flags.Lookup("http-cors-allowed-origins")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.cors_allowed_origins", "SKELETON_HTTP_CORS_ALLOWED_ORIGINS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("server.http.cors_allowed_headers", flags.Lookup("http-cors-allowed-headers")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("server.http.cors_allowed_headers", "SKELETON_HTTP_CORS_ALLOWED_HEADERS"); err != nil {
		panic(err)
	}

	// PROFILER
	if err = viper.BindPFlag("profiler.enabled", flags.Lookup("profiler-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("profiler.enabled", "SKELETON_PROFILER_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("profiler.port", flags.Lookup("profiler-port")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("profiler.port", "SKELETON_PROFILER_PORT"); err != nil {
		panic(err)
	}

	// LOG
	if err = viper.BindPFlag("logger.level", flags.Lookup("log-level")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.level", "SKELETON_LOG_LEVEL"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.output", flags.Lookup("log-output")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.output", "SKELETON_LOG_OUTPUT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.enabled", flags.Lookup("log-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.enabled", "SKELETON_LOG_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.exporter", flags.Lookup("log-exporter")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.exporter", "SKELETON_LOG_EXPORTER"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.endpoint", flags.Lookup("log-endpoint")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.endpoint", "SKELETON_LOG_ENDPOINT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.insecure", flags.Lookup("log-insecure")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.insecure", "SKELETON_LOG_INSECURE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.urlpath", flags.Lookup("log-urlpath")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.urlpath", "SKELETON_LOG_URL_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.headers", flags.Lookup("log-headers")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.headers", "SKELETON_LOG_HEADERS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("logger.protocol", flags.Lookup("log-protocol")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("logger.protocol", "SKELETON_LOG_PROTOCOL"); err != nil {
		panic(err)
	}

	// AUTHN
	if err = viper.BindPFlag("authn.enabled", flags.Lookup("authn-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.enabled", "SKELETON_AUTHN_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.method", flags.Lookup("authn-method")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.method", "SKELETON_AUTHN_METHOD"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.preshared.keys", flags.Lookup("authn-preshared-keys")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.preshared.keys", "SKELETON_AUTHN_PRESHARED_KEYS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.issuer", flags.Lookup("authn-oidc-issuer")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.issuer", "SKELETON_AUTHN_OIDC_ISSUER"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.audience", flags.Lookup("authn-oidc-audience")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.audience", "SKELETON_AUTHN_OIDC_AUDIENCE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.refresh_interval", flags.Lookup("authn-oidc-refresh-interval")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.refresh_interval", "SKELETON_AUTHN_OIDC_REFRESH_INTERVAL"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.backoff_interval", flags.Lookup("authn-oidc-backoff-interval")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.backoff_interval", "SKELETON_AUTHN_OIDC_BACKOFF_INTERVAL"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.backoff_max_retries", flags.Lookup("authn-oidc-backoff-max-retries")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.backoff_max_retries", "SKELETON_AUTHN_OIDC_BACKOFF_RETRIES"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.backoff_frequency", flags.Lookup("authn-oidc-backoff-frequency")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.backoff_frequency", "SKELETON_AUTHN_OIDC_BACKOFF_FREQUENCY"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("authn.oidc.valid_methods", flags.Lookup("authn-oidc-valid-methods")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("authn.oidc.valid_methods", "SKELETON_AUTHN_OIDC_VALID_METHODS"); err != nil {
		panic(err)
	}

	// TRACER
	if err = viper.BindPFlag("tracer.enabled", flags.Lookup("tracer-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.enabled", "SKELETON_TRACER_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.exporter", flags.Lookup("tracer-exporter")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.exporter", "SKELETON_TRACER_EXPORTER"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.endpoint", flags.Lookup("tracer-endpoint")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.endpoint", "SKELETON_TRACER_ENDPOINT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.insecure", flags.Lookup("tracer-insecure")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.insecure", "SKELETON_TRACER_INSECURE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.urlpath", flags.Lookup("tracer-urlpath")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.urlpath", "SKELETON_TRACER_URL_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.headers", flags.Lookup("tracer-headers")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.headers", "SKELETON_TRACER_HEADERS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("tracer.protocol", flags.Lookup("tracer-protocol")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("tracer.protocol", "SKELETON_TRACER_PROTOCOL"); err != nil {
		panic(err)
	}

	// METER
	if err = viper.BindPFlag("meter.enabled", flags.Lookup("meter-enabled")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.enabled", "SKELETON_METER_ENABLED"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.exporter", flags.Lookup("meter-exporter")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.exporter", "SKELETON_METER_EXPORTER"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.endpoint", flags.Lookup("meter-endpoint")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.endpoint", "SKELETON_METER_ENDPOINT"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.insecure", flags.Lookup("meter-insecure")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.insecure", "SKELETON_METER_INSECURE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.urlpath", flags.Lookup("meter-urlpath")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.urlpath", "SKELETON_METER_URL_PATH"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.headers", flags.Lookup("meter-headers")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.headers", "SKELETON_METER_HEADERS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.interval", flags.Lookup("meter-interval")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.interval", "SKELETON_METER_INTERVAL"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("meter.protocol", flags.Lookup("meter-protocol")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("meter.protocol", "SKELETON_METER_PROTOCOL"); err != nil {
		panic(err)
	}

	// SERVICE
	if err = viper.BindPFlag("service.circuit_breaker", flags.Lookup("service-circuit-breaker")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("service.circuit_breaker", "SKELETON_SERVICE_CIRCUIT_BREAKER"); err != nil {
		panic(err)
	}

	// DATABASE
	if err = viper.BindPFlag("database.engine", flags.Lookup("database-engine")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.engine", "SKELETON_DATABASE_ENGINE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.uri", flags.Lookup("database-uri")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.uri", "SKELETON_DATABASE_URI"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.writer.uri", flags.Lookup("database-writer-uri")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.writer.uri", "SKELETON_DATABASE_WRITER_URI"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.reader.uri", flags.Lookup("database-reader-uri")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.reader.uri", "SKELETON_DATABASE_READER_URI"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.auto_migrate", flags.Lookup("database-auto-migrate")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.auto_migrate", "SKELETON_DATABASE_AUTO_MIGRATE"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.max_open_connections", flags.Lookup("database-max-open-connections")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.max_open_connections", "SKELETON_DATABASE_MAX_OPEN_CONNECTIONS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.max_idle_connections", flags.Lookup("database-max-idle-connections")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.max_idle_connections", "SKELETON_DATABASE_MAX_IDLE_CONNECTIONS"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.max_connection_lifetime", flags.Lookup("database-max-connection-lifetime")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.max_connection_lifetime", "SKELETON_DATABASE_MAX_CONNECTION_LIFETIME"); err != nil {
		panic(err)
	}

	if err = viper.BindPFlag("database.max_connection_idle_time", flags.Lookup("database-max-connection-idle-time")); err != nil {
		panic(err)
	}
	if err = viper.BindEnv("database.max_connection_idle_time", "SKELETON_DATABASE_MAX_CONNECTION_IDLE_TIME"); err != nil {
		panic(err)
	}
}
