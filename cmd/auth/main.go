package main

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"zero/config"
	"zero/internal/auth/app"
	"zero/internal/auth/router"
	"zero/pkg/http"
)

const (
	appName                        = "zero"
	appVersion                     = "0.0.0"
	defaultEnv                     = "staging"
	defaultLogLevel                = "info"
	defaultPort                    = 8787
	defaultTokenSigningKey         = "cb-signing-key" // nolint
	defaultTokenExpiryDurationHour = "8"
	defaultTokenTokenIssuer        = "zero"
)

type AppConfig struct {
	// General configuration
	Env      string
	LogLevel string

	// Database configuration
	DatabaseDSN string

	// HTTP configuration
	Port int

	// Token configuration
	TokenSigningKey         string
	TokenExpiryDurationHour int
	TokenIssuer             string
}

var appConfig AppConfig

func main() {
	a := &cli.App{
		Name:     "auth",
		Usage:    "Start the auth service",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Usage:       "Log filtering level",
				Value:       defaultLogLevel,
				EnvVars:     []string{"ZR_LOG_LEVEL"},
				Destination: &appConfig.Env,
			},

			&cli.IntFlag{
				Name:        "server-port",
				Usage:       "server port",
				Value:       defaultPort,
				EnvVars:     []string{"ZR_PORT"},
				Destination: &config.Entrypoint.Port,
			},
		},
		Before: func(ctx *cli.Context) error {
			return nil
		},
		Action: func(c *cli.Context) error {
			r := router.NewRouter(app.ApplicationParams{
				Env:                 appConfig.Env,
				DatabaseDSN:         appConfig.DatabaseDSN,
				TokenSigningKey:     []byte(appConfig.TokenSigningKey),
				TokenExpiryDuration: time.Duration(appConfig.TokenExpiryDurationHour),
				TokenIssuer:         appConfig.TokenIssuer,
			})
			srv := http.NewServer(r.Load())

			return srv.Start()
		},
	}

	if err := a.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
