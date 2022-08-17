package logger

//import (
//	"os"
//	"time"
//
//	"github.com/rs/zerolog"
//	"github.com/rs/zerolog/log"
//
//	"zero/config"
//)
//
//func SetupLogger() {
//	switch *config.AppConfig.LogLevel {
//	case "debug":
//		zerolog.SetGlobalLevel(zerolog.DebugLevel)
//	case "info":
//		zerolog.SetGlobalLevel(zerolog.InfoLevel)
//	case "warn":
//		zerolog.SetGlobalLevel(zerolog.WarnLevel)
//	case "error":
//		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
//	case "disabled":
//		zerolog.SetGlobalLevel(zerolog.Disabled)
//	}
//
//	log.Logger = log.Output(
//		zerolog.ConsoleWriter{
//			Out:        os.Stderr,
//			TimeFormat: time.RFC3339,
//		},
//	).With().Caller().Logger()
//}
