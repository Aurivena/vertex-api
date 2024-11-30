package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"time"
	"vertexUP/models"
	"vertexUP/pkg/utils"
)

type AuthPostgres struct {
	db *sqlx.DB
}

func NewAuthPostgres(db *sqlx.DB) *AuthPostgres {
	return &AuthPostgres{db: db}
}

func (r *AuthPostgres) SignIn(input *models.SignInInput) (*models.SignInOutput, error) {
	return nil, nil
}

func (r *AuthPostgres) SignUp(input *models.SignUpInput) (*models.SignUpOutput, error) {
	output := &models.SignUpOutput{}

	password, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}

	err = r.db.Get(output, `
    INSERT INTO "User" ("login", "name", "email", "password", "status","date_registration")
    VALUES ($1, $2, $3, $4, $5, $6)
    RETURNING "login", "name", "email"`,
		input.Login, input.Name, input.Email, password, utils.User, time.Now().UTC(),
	)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	err = r.db.Get(&output.Status, `SELECT name FROM "Status" WHERE id = $1`, utils.User)
	if err != nil {
		logrus.Error(err.Error())
		return nil, err
	}

	return output, nil
}

func (r *AuthPostgres) SignOut() {

}
