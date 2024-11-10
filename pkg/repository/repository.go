package repository

import "github.com/jmoiron/sqlx"

type Sources struct {
	BusinessDB *sqlx.DB
}

type Repository struct {
	Sources *Sources
}

func NewRepository(sources *Sources) *Repository {
	return &Repository{Sources: sources}
}
