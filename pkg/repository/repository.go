package repository

import (
	"github.com/google/uuid"
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
	UpdateInfoUser(info *models.UpdateInfoAccountInput, token string) error
	GetUserByEmail(email string) (*models.Account, error)
	GetUserByLogin(login string) (*models.Account, error)
	IsRegistered(input string) (bool, error)
	GetUserByAccessToken(accessToken string) (*models.Account, error)
}

type Token interface {
	SaveToken(uuid uuid.UUID, token models.Token) error
	DeleteToken(token string) error
	CheckCount(uuid uuid.UUID) error
	UpdateAccessToken(oldAccessToken string, newAccessToken string, time time.Time) error
	UpdateRefreshToken(oldRefreshToken string, newRefreshToken string, time time.Time) error
	GetAllInfoToken(uuid uuid.UUID) (*models.Token, error)
}

type Middleware interface {
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
