package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"vertexUP/models"
)

type AccountPostgres struct {
	db *sqlx.DB
}

func NewAccountPostgres(db *sqlx.DB) *AccountPostgres {
	return &AccountPostgres{db: db}
}

func (r AccountPostgres) GetUserByEmail(email string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT * FROM "User" WHERE email = $1`, email)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) GetUserByLogin(login string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT * FROM "User" WHERE login = $1 `, login)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) GetUserByRefreshToken(refreshToken string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT * FROM "User"
									INNER JOIN "Token" on "Token".login = "User".login
									WHERE "Token".refresh_token = $1`, refreshToken)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) IsRegistered(input string) (bool, error) {
	var exists bool
	err := r.db.Get(&exists,
		`SELECT EXISTS(
                 SELECT 1 
                 FROM "User" 
            	WHERE "login" = $1 OR "email" = $1) `, input)
	if err != nil {
		logrus.Error(err.Error())
		return false, nil
	}
	return exists, nil
}
