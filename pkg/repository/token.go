package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"time"
	"vertexUP/models"
)

type TokenRepository struct {
	db *sqlx.DB
}

func NewTokenRepository(db *sqlx.DB) *TokenRepository {
	return &TokenRepository{db: db}
}

func (r TokenRepository) SaveToken(login string, token models.Token) error {

	query := `INSERT INTO "Token"
				("login","token","token_expiration","refresh_token_expiration","is_revoked") 
				VALUES ($1,$2,$3,$4,$5) RETURNING "id"`

	_, err := r.db.Exec(query, login, token.AccessToken, token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt, false)
	if err != nil {
		logrus.Error("ошибка сохранения токенов: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) RevokeToken(token string) error {
	query := `UPDATE "Token" "is_revoked" = true WHERE "token" = $1`

	_, err := r.db.Exec(query, token)
	if err != nil {
		logrus.Error("ошибка отзыва токена: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) IsTokenActive(token string) (bool, error) {
	var isRevoked bool
	var expertion time.Time

	query := `SELECT is_revoked, token_expiration 
		FROM "Token" 
		WHERE token = $1`

	err := r.db.QueryRow(query, token).Scan(&isRevoked, &expertion)
	if err != nil {
		logrus.Errorf("ошибка проверки токена: %v", err)
		return false, err
	}

	return !isRevoked && expertion.After(time.Now()), nil
}
