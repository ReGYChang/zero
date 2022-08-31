package service

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"

	"zero/internal/auth/domain"
	"zero/internal/auth/domain/common"
)

func TestAuthService_RegisterUser(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		User     domain.User
		Password string
	}
	var args Args
	_ = faker.FakeData(&args)

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *AuthService
		ExpectError  bool
	}{
		{
			Name: "user does not exist",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(nil, common.DomainError{})
				mock.UserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(&args.User, nil)

				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "failed to register user",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(nil, common.DomainError{})
				mock.UserRepo.EXPECT().CreateUser(gomock.Any(), gomock.Any()).Return(nil, common.DomainError{})

				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
		{
			Name: "user exist",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(&args.User, nil)

				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			param := RegisterUserParam{
				Email:    args.User.Email,
				Name:     args.User.Name,
				Password: args.Password,
			}

			_, err := service.RegisterUser(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}

func TestAuthService_LoginUser(t *testing.T) {
	t.Parallel()
	// Args
	type Args struct {
		User     domain.User
		Password string
	}
	var args Args
	_ = faker.FakeData(&args)

	// Init
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	// Test cases
	testCases := []struct {
		Name         string
		SetupService func(t *testing.T) *AuthService
		ExpectError  bool
	}{
		{
			Name: "success",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(&args.User, nil)

				service := buildService(mock)
				return service
			},
			ExpectError: false,
		},
		{
			Name: "invalid args.Password",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(&args.User, nil)

				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
		{
			Name: "email does not exist",
			SetupService: func(t *testing.T) *AuthService {
				mock := buildServiceMock(ctrl)

				mock.UserRepo.EXPECT().GetUserByEmail(gomock.Any(), args.User.Email).Return(nil, common.DomainError{})

				service := buildService(mock)
				return service
			},
			ExpectError: true,
		},
	}

	for i := range testCases {
		c := testCases[i]
		t.Run(c.Name, func(t *testing.T) {
			service := c.SetupService(t)
			param := LoginUserParam{
				Email:    args.User.Email,
				Password: args.Password,
			}

			_, err := service.LoginUser(context.Background(), param)

			if c.ExpectError {
				require.Error(t, err)
			} else {
				require.NoError(t, err)
			}
		})
	}
}
