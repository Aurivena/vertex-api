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
				("login","access_token","refresh_token","token_expiration","refresh_token_expiration") 
				VALUES ($1,$2,$3,$4,$5) RETURNING "id"`

	_, err := r.db.Exec(query, login, token.AccessToken, token.RefreshToken, token.AccessTokenExpiresAt, token.RefreshTokenExpiresAt)
	if err != nil {
		logrus.Error("ошибка сохранения токенов: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) UpdateAccessToken(refreshToken string, newAccessToken string, time time.Time) (string, error) {
	query := `UPDATE "Token" SET "access_token"=$1, token_expiration = $2  WHERE "refresh_token"=$3`

	_, err := r.db.Exec(query, newAccessToken, time, refreshToken)
	if err != nil {
		logrus.Error("ошибка в обновлении токена: %w", err)
		return "", err
	}

	return newAccessToken, nil
}

func (r TokenRepository) DeleteToken(token string) error {
	query := `DELETE FROM "Token" WHERE "access_token" = $1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		logrus.Error("ошибка удаления токена: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) RefreshAllTokens(login, newAccessToken, newRefreshToken string, newAccessTokenExpiry, newRefreshTokenExpiry time.Time) error {
	query := `
		UPDATE "Token"
		SET access_token = $1, refresh_token = $2, token_expiration = $3, refresh_token_expiration = $4
		WHERE login = $5
	`

	_, err := r.db.Exec(query, newAccessToken, newRefreshToken, newAccessTokenExpiry, newRefreshTokenExpiry, login)
	if err != nil {
		logrus.Errorf("Ошибка обновления токенов для пользователя %s: %v", login, err)
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

func (r TokenRepository) GetTimeToken(login string) (*models.Token, error) {
	output := models.Token{}

	query := `SELECT "token_expiration","refresh_token_expiration" FROM "Token" WHERE login = $1`

	err := r.db.Select(&output, query, login)
	if err != nil {
		logrus.Error("ошибка при получение времени жизни токенов")
		return nil, err
	}

	return &output, nil
}
