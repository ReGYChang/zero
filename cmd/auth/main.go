package main

import (
	"fmt"
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"zero/config"
	"zero/internal/auth/app"
	"zero/internal/auth/router"
	"zero/pkg/http"
)

const (
	appName                        = "auth"
	appVersion                     = "0.0.0"
	defaultEnv                     = "staging"
	defaultLogLevel                = "info"
	defaultPort                    = 8787
	defaultTokenSigningKey         = "cb-signing-key" // nolint
	defaultTokenExpiryDurationHour = 8
	defaultTokenIssuer             = "zero"
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
		Version: fmt.Sprintf(
			"version %s, built on %s",
			appVersion,
			runtime.Version(),
		),
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

			&cli.StringFlag{
				Name:        "token-signing-key",
				Usage:       "token signing key",
				Value:       defaultTokenSigningKey,
				EnvVars:     []string{"ZR_TOKEN_SIGNING_KEY"},
				Destination: &appConfig.TokenSigningKey,
			},

			&cli.IntFlag{
				Name:        "token-expiry-duration-hour",
				Usage:       "token expiry duration hour",
				Value:       defaultTokenExpiryDurationHour,
				EnvVars:     []string{"ZR_TOKEN_EXPIRE_DURATION_HOUR"},
				Destination: &appConfig.TokenExpiryDurationHour,
			},

			&cli.StringFlag{
				Name:        "token-issuer",
				Usage:       "token issuer",
				Value:       defaultTokenIssuer,
				EnvVars:     []string{"ZR_PORT"},
				Destination: &appConfig.TokenIssuer,
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
