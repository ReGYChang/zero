package service

import (
	"context"

	"zero/internal/auth/domain"
	"zero/internal/auth/domain/common"
)

type AuthServer interface {
	RegisterAccount(ctx context.Context, email string, password string) (string, common.Error)
	AuthenticateAccount(ctx context.Context, email string, password string) common.Error
}

type UserRepository interface {
	GetUserByEmail(ctx context.Context, email string) (*domain.User, common.Error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, common.Error)
}
