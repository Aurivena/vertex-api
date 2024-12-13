package service

import (
	"github.com/google/uuid"
	"vertexUP/models"
	"vertexUP/pkg/repository"
)

type Auth interface {
	SignIn(input *models.SignInInput) (*models.SignInOutput, error)
	SignUp(input *models.SignUpInput) (*models.SignUpOutput, error)
}

type Account interface {
	UpdateInfoAccount(info *models.UpdateInfoAccountInput, token string) error
	GetUserByEmail(email string) (*models.Account, error)
	GetUserByLogin(login string) (*models.Account, error)
	IsRegistered(input string) (bool, error)
	GetUserByAccessToken(refreshToken string) (*models.Account, error)
}

type Token interface {
	GenerateTokenAndSave(uuid uuid.UUID, login string) (*models.Token, error)
	Logout(token string) error
	RefreshAllToken(uuid uuid.UUID, login string) (*models.Token, error)
	CheckValidUser(uuid uuid.UUID) (bool, error)
}

type Middleware interface {
}

type Service struct {
	Auth
	Account
	Token
	Middleware
}

func NewService(repos *repository.Repository, cfg *models.Config, env *models.Environment) *Service {
	return &Service{
		Auth:       NewAuthService(repos.Auth, NewAccountService(repos.Account)),
		Account:    NewAccountService(repos.Account),
		Token:      NewTokenService(repos.Token, cfg.Secret.Secret),
		Middleware: NewMiddlewareService(repos.Middleware),
	}
}
