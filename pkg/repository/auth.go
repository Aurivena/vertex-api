package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
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
	var output models.SignInOutput

	err := r.db.Get(&output,
		`SELECT "User"."name" as "name" , "login", "email", st."name" as "status", "date_registration"
			   FROM "User"
			   INNER JOIN "Status" st ON st.id = "User".status
			   WHERE "login" = $1 or "email" = $1`, input.Input)
	if err != nil {
		return nil, err
	}

	return &output, nil
}

func (r *AuthPostgres) SignUp(input *models.SignUpInput, time time.Time) (*models.SignUpOutput, error) {
	var output models.SignUpOutput

	uuid, err := utils.GenerateUUID()
	if err != nil {
		return nil, err
	}

	err = r.db.Get(&output, `
    INSERT INTO "User" ("uuid","login", "name", "email", "password", "status","date_registration")
    VALUES ($1, $2, $3, $4, $5, $6,$7)
    RETURNING "name","login","email","date_registration"`,
		uuid, input.Login, input.Name, input.Email, input.Password, utils.User, time,
	)
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}

	err = r.db.Get(&output.Status, `SELECT name FROM "Status" WHERE id = $1`, utils.User)
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}

	return &output, nil
}
