package service

import (
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type Auth interface {
	SignIn(input *models.SignInInput) (*models.SignInOutput, error)
	SignUp(input *models.SignUpInput) (*models.SignUpOutput, error)
	SignOut()
}

type Account interface {
	GetUserByEmail(email string) (*models.Account, error)
	GetUserByLogin(login string) (*models.Account, error)
	IsRegistered(input string) (bool, error)
}

type Service struct {
	Auth
	Account
}

func NewService(repos *repository.Repository, cfg *models.Config, env *models.Environment) *Service {
	return &Service{
		Auth:    NewAuthService(repos.Auth, NewAccountService(repos.Account)),
		Account: NewAccountService(repos.Account),
	}
}
