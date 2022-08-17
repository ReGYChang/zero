package service

import (
	"context"
	"errors"

	"github.com/google/uuid"

	"zero/internal/auth/domain/auth"
	"zero/internal/auth/domain/common"
)

type RegisterUserParam struct {
	Email    string
	Name     string
	Password string
}

func (s *AuthService) RegisterUser(ctx context.Context, param RegisterUserParam) (*auth.User, common.Error) {
	// Check the given user email exist or not
	_, err := s.userRepo.GetUserByEmail(ctx, param.Email)
	if err == nil {
		msg := "email exists"
		s.logger(ctx).Error().Msg(msg)
		return nil, common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg))
	}

	// If not existed:
	// 1. Create a user in the application.
	uid := uuid.NewString()
	user, err := s.userRepo.CreateUser(ctx, auth.NewUser(uid, param.Email, param.Name))
	if err != nil {
		s.logger(ctx).Error().Err(err).Msg("failed to register user")
		return nil, err
	}

	return user, nil
}

type LoginUserParam struct {
	Email    string
	Password string
}

func (s *AuthService) LoginUser(ctx context.Context, param LoginUserParam) (*auth.User, common.Error) {
	// Authenticate the account
	user, err := s.userRepo.GetUserByEmail(ctx, param.Email)
	if err != nil {
		s.logger(ctx).Error().Err(err).Msg("failed to get user")
		msg := "email does not exist"
		return nil, common.NewError(common.ErrorCodeParameterInvalid, errors.New(msg), common.WithMsg(msg))
	}

	return user, nil
}
