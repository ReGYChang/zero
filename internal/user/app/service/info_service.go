package service

import "context"

type InfoService struct {
}

type InfoServiceParam struct {
}

func NewInfoService(ctx context.Context, param *RegisterServiceParam) *InfoService {
	return &InfoService{}
}
