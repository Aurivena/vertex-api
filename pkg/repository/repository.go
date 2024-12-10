package repository

import (
	"github.com/jmoiron/sqlx"
	"time"
	"vertexUP/models"
)

type Sources struct {
	BusinessDB *sqlx.DB
}

type Auth interface {
	SignIn(input *models.SignInInput) (*models.SignInOutput, error)
	SignUp(input *models.SignUpInput, time time.Time) (*models.SignUpOutput, error)
}

type Account interface {
	GetUserByEmail(email string) (*models.Account, error)
	GetUserByLogin(login string) (*models.Account, error)
	IsRegistered(input string) (bool, error)
	GetUserByRefreshToken(refreshToken string) (*models.Account, error)
}

type Token interface {
	SaveToken(login string, token models.Token) error
	RevokeToken(token string) error
	CheckCount(login string) int
	UpdateAccessToken(refreshToken string, newAccessToken string, time time.Time) (string, error)
	RefreshAllTokens(login, newAccessToken, newRefreshToken string, newAccessTokenExpiry, newRefreshTokenExpiry time.Time) error
	GetTimeToken(login string) (*models.Token, error)
}

type Middleware interface {
	IsTokenActive(token string) (bool, error)
}

type Repository struct {
	Sources *Sources
	Auth
	Account
	Token
	Middleware
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Sources:    sources,
		Auth:       NewAuthPostgres(sources.BusinessDB),
		Account:    NewAccountPostgres(sources.BusinessDB),
		Token:      NewTokenRepository(sources.BusinessDB),
		Middleware: NewMiddlewareRepository(sources.BusinessDB),
	}
}
