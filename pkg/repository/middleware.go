package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
)

type MiddlewareRepository struct {
	db *sqlx.DB
}

func NewMiddlewareRepository(db *sqlx.DB) *MiddlewareRepository {
	return &MiddlewareRepository{db: db}
}

func (r MiddlewareRepository) IsTokenActive(token string) (bool, error) {
	var isRevoked bool
	var expertion time.Time

	query := `SELECT is_revoked, token_expiration 
		FROM "Token" 
		WHERE access_token = $1`

	err := r.db.QueryRow(query, token).Scan(&isRevoked, &expertion)
	if err != nil {
		logrus.Errorf("ошибка проверки токена: %v", err)
		return false, err
	}

	return !isRevoked && expertion.After(time.Now()), nil
}
