package service

import (
	"JettyPayGo/internal/biz"
)

type Service struct {
	biz *biz.Biz
}

func ProvideService(biz *biz.Biz) *Service {
	return &Service{biz: biz}
}
