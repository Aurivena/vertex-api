package usecase

import "vertexUP/pkg/service"

type Usecase struct {
	services *service.Service
}

func NewUsecase(servicee *service.Service) *Usecase {
	return &Usecase{services: servicee}
}
