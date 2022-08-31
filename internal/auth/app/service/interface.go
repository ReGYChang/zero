package service

import (
	"context"

	"zero/internal/auth/domain"
	"zero/internal/auth/domain/common"
)

//go:generate mockgen -destination automock/user_repository.go -package=automock . UserRepository
type UserRepository interface {
	AuthenticateUser(ctx context.Context, email string, password string) common.Error
	GetUserByEmail(ctx context.Context, email string) (*domain.User, common.Error)
	CreateUser(ctx context.Context, user domain.User) (*domain.User, common.Error)
}
