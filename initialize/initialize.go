package initialize

import (
	"encoding/json"
	"fmt"
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"io"
	"os"
	"path/filepath"
	"strings"
	"time"
	"vertexUP/models"
)

const (
	envFilePath = `.env.local`
)

var (
	Config = models.Config{}
	Env    = &models.Environment{}
)

func LoadConfiguration() error {
	if err := loadEnvironment(envFilePath); err != nil {
		return err
	}
	logrus.Info("environment variables: loaded")
	if err := loadConfig(Env.ConfigPath); err != nil {
		return err
	}
	logrus.Info("config file: loaded")
	return nil
}

func loadConfig(path string) error {
	file, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(file, &Config)
	if err != nil {
		return err
	}

	return nil
}

func loadEnvironment(envPath string) error {
	if err := godotenv.Load(envPath); err != nil {
		logrus.Warning("load file not found, environment variables load from environment")
	}
	if err := env.Parse(Env); err != nil {
		return err
	}
	return nil
}

type CustomFormatter struct{}

func (f *CustomFormatter) Format(entry *logrus.Entry) ([]byte, error) {
	pwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	prefixPath := pwd + "/"

	shortFilePath := strings.TrimPrefix(filepath.ToSlash(entry.Caller.File), filepath.ToSlash(prefixPath))

	var fields string
	for key, value := range entry.Data {
		fields += fmt.Sprintf("\"%s\":\"%v\",", key, value)
	}

	if len(fields) > 0 {
		fields = fields[:len(fields)-1]
	}

	if len(fields) > 0 {
		fields = ", " + fields
	}

	log := fmt.Sprintf(
		"{\"level\":\"%s\",\"msg\":\"%s\",\"point\": \" %s:%d \",\"short_point\":\"%s:%d\", \"time\":\"%s\"%s}\n",
		entry.Level.String(),
		entry.Message,
		entry.Caller.File,
		entry.Caller.Line,
		shortFilePath,
		entry.Caller.Line,
		entry.Time.Format(time.RFC3339),
		fields,
	)
	return []byte(log), nil
}

func RunLogger() error {
	logrus.SetReportCaller(true)
	logrus.SetFormatter(&CustomFormatter{})

	currentTime := time.Now()
	yearMonthDir := fmt.Sprintf("logs/%d-%02d", currentTime.Year(), currentTime.Month())

	err := os.MkdirAll(yearMonthDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("ошибка создания папки для логов: %v", err)
	}

	logFile := fmt.Sprintf("%s/%d-%02d-%02d.log", yearMonthDir, currentTime.Year(), currentTime.Month(), currentTime.Day())

	logFileHandle, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		return fmt.Errorf("ошибка открытия файла логов: %v", err)
	}

	logrus.SetOutput(io.MultiWriter(os.Stdout, logFileHandle))

	return nil
}
