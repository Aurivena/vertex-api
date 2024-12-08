package models

type ServerConfig struct {
	Port string `json:"port"`
}

type BusinessDBConfig struct {
	Host     string `json:"host"`
	Port     string `json:"port"`
	Username string `json:"username"`
	DbName   string `json:"db_name"`
	SslMode  string `json:"ssl_mode"`
}

type Jwt struct {
	Secret string `json:"secret"`
}

type Config struct {
	Server   ServerConfig     `json:"server"`
	Database BusinessDBConfig `json:"business-database"`
	Secret   Jwt              `json:"jwt"`
}
