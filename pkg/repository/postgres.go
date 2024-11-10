package repository

import (
	"errors"
	"fmt"
	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"github.com/sirupsen/logrus"
	"os"
	"vertexUP/server/ServerMode"
)

const (
	dbDriverName                = "pgx"
	sqlScriptsDataBaseDirectory = "migrations"
	createTablesFilePrefix      = "create.tables."
)

type PostgresDBConfig struct {
	Host     string
	Port     string
	Username string
	Password string
	DBName   string
	SSLMode  string
}

func NewPostgresDB(config *PostgresDBConfig, serverMode string) (*sqlx.DB, error) {
	var pgErr pgx.PgError

	db, err := getDBConnection(config)

	if err == nil {
		return db, nil
	} else {
		fmt.Println(err.Error())
	}

	if serverMode == ServerMode.DEVELOPMENT && errors.As(err, &pgErr) {
		if pgerrcode.IsInvalidCatalogName(pgErr.Code) {
			logrus.Warning(fmt.Sprintf(`database "%s" not found`, config.DBName))
			if err = createDatabase(config); err != nil {
				logrus.Errorf(`failed create "%s" database: %s`, nil, config.DBName, err.Error())
			}
			if err = createTables(config); err != nil {
				logrus.Errorf(`create tables in "%s" database: %s`, nil, config.DBName, err.Error())
			}
			return getDBConnection(config)
		}
	}

	return nil, err
}

func getDBConnection(config *PostgresDBConfig) (*sqlx.DB, error) {
	return sqlx.Connect(dbDriverName, getConnectionString(config))
}

func createDatabase(config *PostgresDBConfig) error {
	db, err := sqlx.Connect(
		dbDriverName, fmt.Sprintf("host=%s port=%s user=%s password=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.SSLMode))
	if err != nil {
		logrus.Error(err.Error())
		return err
	}

	query := fmt.Sprintf("%s %s", "create database", config.DBName)
	_, err = db.Exec(query)
	if err != nil {
		logrus.Fatal(err.Error())
		return err
	}

	if err = db.Close(); err != nil {
		logrus.Fatal(err.Error())
		return err
	}

	return nil
}

func createTables(config *PostgresDBConfig) error {
	db, err := sqlx.Connect(
		dbDriverName, fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
			config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode))
	if err != nil {
		logrus.Fatal(err.Error())
		return err
	}

	query, err := os.ReadFile(fmt.Sprintf("%s/%s%s.sql",
		sqlScriptsDataBaseDirectory, createTablesFilePrefix, config.DBName))
	if err != nil {
		logrus.Fatal(err.Error())
		return err
	}
	stringQuery := string(query)

	_, err = db.Exec(stringQuery)
	if err != nil {
		logrus.Fatal(err.Error())
		return err
	}

	if err = db.Close(); err != nil {
		logrus.Fatal(err.Error())
		return err
	}

	return nil
}

func getConnectionString(config *PostgresDBConfig) string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=%s",
		config.Host, config.Port, config.Username, config.Password, config.DBName, config.SSLMode)
}
