package repository

import (
	"github.com/jmoiron/sqlx"
)

type MiddlewareRepository struct {
	db *sqlx.DB
}

func NewMiddlewareRepository(db *sqlx.DB) *MiddlewareRepository {
	return &MiddlewareRepository{db: db}
}
