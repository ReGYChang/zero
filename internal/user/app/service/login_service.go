package service

import "context"

type LoginService struct {
}

type LoginServiceParam struct {
}

func NewLoginService(ctx context.Context, param *RegisterServiceParam) *LoginService {
	return &LoginService{}
}
