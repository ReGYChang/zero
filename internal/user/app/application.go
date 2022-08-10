package app

import (
	"time"

	"zero/internal/user/app/service"
)

type Application struct {
	Params          ApplicationParams
	LoginService    *service.LoginService
	RegisterService *service.RegisterService
	InfoService     *service.InfoService
}

type ApplicationParams struct {
	// General configuration
	Env string

	// Database parameters
	DatabaseDSN string

	// Token parameter
	TokenSigningKey     []byte
	TokenExpiryDuration time.Duration
	TokenIssuer         string
}
