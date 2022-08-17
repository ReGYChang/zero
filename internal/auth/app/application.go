package app

import (
	"context"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"

	"zero/internal/auth/adapter"
	"zero/internal/auth/app/service"
)

type Application struct {
	Params      ApplicationParams
	AuthService *service.AuthService
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

func NewApplication(ctx context.Context, wg *sync.WaitGroup, params ApplicationParams) (*Application, error) {
	// Create repositories
	db := sqlx.MustOpen("postgres", params.DatabaseDSN)
	if err := db.Ping(); err != nil {
		return nil, err
	}
	pgRepo := adapter.NewPostgresRepository(ctx, db)

	// Create application
	app := &Application{
		Params: params,
		AuthService: service.NewAuthService(ctx, &service.AuthServiceParam{
			UserRepo:       pgRepo,
			SigningKey:     params.TokenSigningKey,
			ExpiryDuration: params.TokenExpiryDuration,
			Issuer:         params.TokenIssuer,
		}),
	}

	return app, nil
}
