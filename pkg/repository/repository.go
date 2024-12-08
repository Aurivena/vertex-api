package repository

import (
	"github.com/jmoiron/sqlx"
	"vertexUP/models"
)

type Sources struct {
	BusinessDB *sqlx.DB
}

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

type Token interface {
	SaveToken(login, string, token models.Token) error
	RevokeToken(token string) error
	IsTokenActive(token string) (bool, error)
}

type Repository struct {
	Sources *Sources
	Auth
	Account
	Token
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Sources: sources,
		Auth:    NewAuthPostgres(sources.BusinessDB),
		Account: NewAccountPostgres(sources.BusinessDB),
		Token:   NewTokenRepository(sources.BusinessDB),
	}
}
