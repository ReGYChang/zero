package main

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"github.com/urfave/cli/v2"

	"zero/pkg/config"
	"zero/pkg/logger"
)

const (
	appName                        = "zero"
	appVersion                     = "0.0.0"
	defaultEnv                     = "staging"
	defaultLogLevel                = "info"
	defaultPort                    = "8787"
	defaultTokenSigningKey         = "cb-signing-key" // nolint
	defaultTokenExpiryDurationHour = "8"
	defaultTokenTokenIssuer        = "zero"
)

func main() {
	app := &cli.App{
		Name:     "user",
		Usage:    "Start the user service",
		Compiled: time.Now(),
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "log-level",
				Usage:       "Log filtering level",
				EnvVars:     []string{"ZR_LOG_LEVEL"},
				Destination: config.AppConfig.Env,
			},

			&cli.IntFlag{
				Name:        "server-port",
				Usage:       "server port",
				EnvVars:     []string{"ZR_PORT"},
				Destination: config.AppConfig.Port,
			},
		},
		Before: func(ctx *cli.Context) error {
			logger.SetupLogger()
			return nil
		},
		Action: func(c *cli.Context) error {
			//r := router.NewRouter()
			//s := grpc.NewServer(r)
			//
			//return s.Start()
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal().Err(err).Msg("")
	}
}
