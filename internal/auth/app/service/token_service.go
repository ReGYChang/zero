package service

import (
	"context"

	"zero/internal/auth/domain"
	"zero/internal/auth/domain/common"
)

func (s *AuthService) GenerateUserToken(_ context.Context, user domain.User) (string, common.Error) {
	signedToken, err := domain.GenerateUserToken(user, s.signingKey, s.expiryDuration, s.issuer)
	if err != nil {
		return "", common.NewError(common.ErrorCodeParameterInvalid, err, common.WithMsg(err.ClientMsg()))
	}
	return signedToken, nil
}

func (s *AuthService) ValidateUserToken(_ context.Context, signedToken string) (*domain.User, common.Error) {
	user, err := domain.ParseUserFromToken(signedToken, s.signingKey)
	if err != nil {
		return nil, common.NewError(common.ErrorCodeAuthNotAuthenticated, err, common.WithMsg(err.ClientMsg()))
	}
	return user, nil
}
