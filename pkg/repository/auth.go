package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"vertexUP/clerr"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignIn() {

}

func (r *AuthPostgres) SignUp() {

}

func (r *AuthPostgres) SignOut() {

}

func validatePassword(password string) error {
	if len(password) < 8 {
		logrus.Error("password must be at least 8 characters")
		return errors.New(clerr.ErrorPasswordTooShort)
	}

	return nil
}
