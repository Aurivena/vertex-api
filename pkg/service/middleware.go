package service

import (
	"vertexUP/pkg/repository"
)

type MiddlewareService struct {
	repo repository.Middleware
}

func NewMiddlewareService(repo repository.Middleware) *MiddlewareService {
	return &MiddlewareService{repo: repo}
}
