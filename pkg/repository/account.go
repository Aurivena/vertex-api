package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"strconv"
	"strings"
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
		logrus.Errorf(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) GetUserByLogin(login string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT * FROM "User" WHERE login = $1 `, login)
	if err != nil {
		logrus.Errorf(err.Error())
		return nil, err
	}

	return &output, nil
}

func (r AccountPostgres) GetUserByAccessToken(accessToken string) (*models.Account, error) {
	output := models.Account{}

	err := r.db.Get(&output, `SELECT name,email,status,"User".login FROM "User"
									INNER JOIN "Token" on "Token".user = "User".uuid
									WHERE "Token".access_token = $1`, accessToken)
	if err != nil {
		logrus.Errorf(err.Error())
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
		logrus.Errorf(err.Error())
		return false, nil
	}
	return exists, nil
}

func (r AccountPostgres) UpdateInfoUser(info *models.UpdateInfoAccountInput, token string) error {
	tx, err := r.db.Beginx()
	defer tx.Rollback()

	query := `UPDATE "User" SET`
	args := []interface{}{}
	conditions := []string{}
	count := 1

	if info.Name != "" {
		conditions = append(conditions, `"name" = $`+strconv.Itoa(count))
		args = append(args, info.Name)
		count++
	}

	if info.Login != "" {
		conditions = append(conditions, `"login" = $`+strconv.Itoa(count))
		args = append(args, info.Login)
		count++
	}

	if info.Email != "" {
		conditions = append(conditions, `"email" = $`+strconv.Itoa(count))
		args = append(args, info.Email)
		count++
	}

	if info.Password != "" {
		conditions = append(conditions, `"password" = $`+strconv.Itoa(count))
		args = append(args, info.Password)
		count++
	}

	if len(conditions) == 0 {
		return fmt.Errorf("no data to update")
	}

	query += " " + strings.Join(conditions, ",") + `
							FROM "Token"
							WHERE "Token"."user" = "User"."uuid" AND "Token"."access_token" = $` + strconv.Itoa(count)

	args = append(args, token)

	_, err = tx.Exec(query, args...)
	if err != nil {
		logrus.Errorf("ошибка при обновлении данных %w: ", err)
		return err
	}

	if err = tx.Commit(); err != nil {
		logrus.Errorf("failed to commit transaction: %v", err)
		return err
	}

	return nil
}
