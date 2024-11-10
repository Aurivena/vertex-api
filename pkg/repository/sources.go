package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"vertexUP/models"
)

type Source struct {
	BusinessDB *sqlx.DB
}

func NewBusinessDatabase(env *models.Environment, config *models.BusinessDBConfig) *sqlx.DB {
	fmt.Println("start database connected")
	database, err := NewPostgresDB(&PostgresDBConfig{
		Host:     config.Host,
		Port:     config.Port,
		Username: config.Username,
		Password: env.BusinessDBPassword,
		DBName:   config.DbName,
		SSLMode:  config.SslMode,
	}, env.ServerMode)
	if err != nil {
		logrus.Fatalf("failed to initialize business db: %s", err.Error())
	}
	fmt.Println("database connected")
	return database
}
