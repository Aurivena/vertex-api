package models

type Environment struct {
	BusinessDBPassword   string `env:"BUSINESS_DB_PASSWORD,notEmpty"`
	ServerMode           string `env:"SERVER_MODE,notEmpty"`
	Domain               string `env:"DOMAIN,notEmpty"`
	LogDirectory         string `env:"LOG_DIRECTORY,notEmpty"`
	ConfigPath           string `env:"CONFIG_PATH,notEmpty"`
	IsVerifyDependencies bool   `env:"IS_VERIFY_DEPENDENCIES,notEmpty"`
}
