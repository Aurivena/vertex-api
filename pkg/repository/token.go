package repository

import (
	"github.com/google/uuid"
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

func (r TokenRepository) SaveToken(uuid uuid.UUID, token models.Token) error {
	tx, err := r.db.Beginx()
	defer tx.Rollback()
	query := `INSERT INTO "Token"
				("user","access_token","refresh_token","access_token_expiration","refresh_token_expiration") 
				VALUES ($1,$2,$3,$4,$5) RETURNING "id"`

	_, err = tx.Exec(query, uuid, token.AccessToken, token.RefreshToken, token.AccessTokenExpires, token.RefreshTokenExpires)
	if err != nil {
		logrus.Errorf("ошибка сохранения токенов: %w", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("ошибка при коммите транзакции: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) UpdateAccessToken(oldAccessToken string, newAccessToken string, time time.Time) error {
	tx, err := r.db.Beginx()
	defer tx.Rollback()
	query := `UPDATE "Token" SET "access_token"=$1, access_token_expiration = $2  WHERE "access_token"=$3`

	_, err = tx.Exec(query, newAccessToken, time, oldAccessToken)
	if err != nil {
		logrus.Errorf("ошибка в обновлении токена: %w", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("ошибка при коммите транзакции: %w", err)
		return err
	}

	return nil
}

func (r TokenRepository) UpdateRefreshToken(oldRefreshToken string, newRefreshToken string, time time.Time) error {
	tx, err := r.db.Beginx()
	defer tx.Rollback()
	query := `UPDATE "Token" SET "refresh_token"=$1, refresh_token_expiration = $2  WHERE "refresh_token"=$3`

	_, err = tx.Exec(query, newRefreshToken, time, oldRefreshToken)
	if err != nil {
		logrus.Errorf("ошибка в обновлении токена: %w", err)
		return err
	}

	err = tx.Commit()
	if err != nil {
		logrus.Errorf("ошибка при коммите транзакции: %w", err)
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

func (r TokenRepository) CheckCount(uuid uuid.UUID) error {
	count := 0
	query := `SELECT COUNT(*) FROM "Token" WHERE "user" = $1`

	err := r.db.QueryRow(query, uuid).Scan(&count)
	if err != nil {
		return err
	}

	if count == 5 {
		query = `DELETE FROM "Token"
				 WHERE "id" = (
				 	SELECT "id"
				 	FROM "Token"
				 	WHERE "user" = $1
				 	ORDER BY "access_token_expiration"
				 	LIMIT 1
				 )
				`
		_, err = r.db.Exec(query, uuid)
		if err != nil {
			logrus.Errorf("ошибка удаления токена: %w", err)
			return err
		}
	}

	return err
}

func (r TokenRepository) GetAllInfoToken(uuid uuid.UUID) (*models.Token, error) {
	output := models.Token{}

	query := `SELECT "access_token", "refresh_token", "access_token_expiration","refresh_token_expiration" FROM "Token" WHERE "user" = $1`

	err := r.db.Get(&output, query, uuid)
	if err != nil {
		logrus.Errorf("ошибка при получение инфомарции токенов")
		return nil, err
	}

	return &output, nil
}
