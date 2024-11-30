package service

import (
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type Auth interface {
	SignIn()
	SignUp()
	SignOut()
}

type Service struct {
	Auth
}

func NewService(repos *repository.Repository, cfg *models.Config, env *models.Environment) *Service {
	return &Service{
		Auth: NewAuthService(repos.Auth),
	}
}
