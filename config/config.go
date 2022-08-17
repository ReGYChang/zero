package config

type entrypoint struct {
	// HTTP/gRPC configuration
	Type string
	Port int
}

var (
	Entrypoint = &entrypoint{}
)
