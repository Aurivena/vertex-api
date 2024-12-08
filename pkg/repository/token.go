package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
				("login","access_token","refresh_token","token_expiration","refresh_token_expiration","is_revoked") 
				VALUES ($1,$2,$3,$4,$5,$6) RETURNING "id"`

	_, err := r.db.Exec(query, login, token.AccessToken, token.RefreshToken, token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt, false)
	if err != nil {
		logrus.Error("ошибка сохранения токенов: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) RevokeToken(token string) error {
	query := `UPDATE "Token" SET "is_revoked" = true WHERE "access_token" = $1`

	_, err := r.db.Exec(query, token)
	if err != nil {
		logrus.Error("ошибка отзыва токена: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) CheckCount(login string) int {
	count := 0
	query := `SELECT COUNT(*) FROM "Token" WHERE login = $1`

	r.db.QueryRow(query, login).Scan(&count)

	return count
}
