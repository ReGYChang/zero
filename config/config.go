package config

type appConfig struct {
	// General configuration
	Env      *string
	LogLevel *string

	// Database configuration
	DatabaseDSN *string

	// HTTP/gRPC configuration
	Type *string
	Port *int

	// Token configuration
	TokenSigningKey         *string
	TokenExpiryDurationHour *int
	TokenIssuer             *string
}

var (
	AppConfig = &appConfig{}
)
