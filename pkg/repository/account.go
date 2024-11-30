package repository

import (
	"github.com/jmoiron/sqlx"
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
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) GetUserByLogin(login string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT * FROM "User" WHERE login = $1 `, login)
	if err != nil {
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
		return false, nil
	}
	return exists, nil
}
