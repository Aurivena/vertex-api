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

type Repository struct {
	Sources *Sources
	Auth
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{
		Sources: sources,
		Auth:    NewAuthPostgres(sources.BusinessDB),
	}
}
