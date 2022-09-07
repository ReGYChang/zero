package logger

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"
)

func InitRootLogger(levelStr, appName, env string) zerolog.Logger {
	// Set global log level
	level, err := zerolog.ParseLevel(levelStr)
	if err != nil {
		level = zerolog.InfoLevel
	}
	zerolog.SetGlobalLevel(level)

	// Set logger time format
	const rfc3339Micro = "2006-01-02T15:04:05.000000Z07:00"
	zerolog.TimeFieldFormat = rfc3339Micro

	serviceName := fmt.Sprintf("%s-%s", appName, env)
	rootLogger := zerolog.New(os.Stdout).With().
		Timestamp().
		Str("service", serviceName).
		Logger()

	return rootLogger
}
