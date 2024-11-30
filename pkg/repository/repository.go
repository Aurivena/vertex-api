package repository

import "github.com/jmoiron/sqlx"

type Sources struct {
	BusinessDB *sqlx.DB
}

type Auth interface {
	SignIn()
	SignUp()
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
