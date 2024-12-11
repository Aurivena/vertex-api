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
				("login","access_token","refresh_token","access_token_expiration","refresh_token_expiration") 
				VALUES ($1,$2,$3,$4,$5) RETURNING "id"`

	_, err := r.db.Exec(query, login, token.AccessToken, token.RefreshToken, token.AccessTokenExpires, token.RefreshTokenExpires)
	if err != nil {
		logrus.Errorf("ошибка сохранения токенов: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) UpdateAccessToken(login string, newAccessToken string, time time.Time) error {
	query := `UPDATE "Token" SET "access_token"=$1, access_token_expiration = $2  WHERE "login"=$3`

	_, err := r.db.Exec(query, newAccessToken, time, login)
	if err != nil {
		logrus.Errorf("ошибка в обновлении токена: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) UpdateRefreshToken(login string, newRefreshToken string, time time.Time) error {
	query := `UPDATE "Token" SET "refresh_token"=$1, refresh_token_expiration = $2  WHERE "login"=$3`

	_, err := r.db.Exec(query, newRefreshToken, time, login)
	if err != nil {
		logrus.Errorf("ошибка в обновлении токена: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) DeleteToken(token string) error {
	query := `DELETE FROM "Token" WHERE "access_token" = $1`
	_, err := r.db.Exec(query, token)
	if err != nil {
		logrus.Errorf("ошибка удаления токена: %w", err)
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

func (r TokenRepository) GetAllInfoToken(login string) (*models.Token, error) {
	output := models.Token{}

	query := `SELECT "access_token", "refresh_token", "access_token_expiration","refresh_token_expiration" FROM "Token" WHERE login = $1`

	err := r.db.Get(&output, query, login)
	if err != nil {
		logrus.Errorf("ошибка при получение инфомарции токенов")
		return nil, err
	}

	return &output, nil
}
