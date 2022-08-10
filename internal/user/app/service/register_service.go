package service

import "context"

type RegisterService struct {
}

type RegisterServiceParam struct {
}

func NewRegisterService(ctx context.Context, param *RegisterServiceParam) *RegisterService {
	return &RegisterService{}
}
