package service

import (
	"context"
	"testing"

	"github.com/bxcodec/faker/v3"
	"github.com/golang/mock/gomock"

	"zero/internal/auth/app/service/automock"
)

type serviceMock struct {
	UserRepo *automock.MockUserRepository
}

func buildServiceMock(ctrl *gomock.Controller) serviceMock {
	return serviceMock{
		UserRepo: automock.NewMockUserRepository(ctrl),
	}
}
func buildService(mock serviceMock) *AuthService {
	param := &AuthServiceParam{
		UserRepo: mock.UserRepo,
	}
	return NewAuthService(context.Background(), param)
}

// nolint
func TestMain(m *testing.M) {
	// To avoid getting an empty object slice
	_ = faker.SetRandomMapAndSliceMinSize(2)

	// To avoid getting a zero random number
	_ = faker.SetRandomNumberBoundaries(1, 100)

	m.Run()
}
