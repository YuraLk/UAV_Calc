package configs

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Config struct {
	Env        string
	HTTPServer HTTPServer
	Database   Database
	JWT        JWT
}

type HTTPServer struct {
	Host string
	Port string
}

type Database struct {
	Host     string
	Port     string
	User     string
	Password string
	Name     string
	SSLMode  string
}

type JWT struct {
	JWTAccessKey  string
	JWTRefreshKey string
}

var Cfg *Config

func Load() {
	var cfg = Config{
		Env: os.Getenv("APP_ENV"),
		HTTPServer: HTTPServer{
			Host: os.Getenv("HOST"),
			Port: os.Getenv("PORT"),
		},
		Database: Database{
			Host:     os.Getenv("POSTGRES_HOST"),
			Port:     os.Getenv("POSTGRES_PORT"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
			Name:     os.Getenv("POSTGRES_DB"),
			SSLMode:  os.Getenv("POSTGRES_SSL_MODE"),
		},
		JWT: JWT{
			JWTAccessKey:  os.Getenv("JWT_ACCESS_SECRET"),
			JWTRefreshKey: os.Getenv("JWT_REFRESH_SECRET"),
		},
	}

	if cfg.Env != "dev" {
		gin.SetMode(gin.ReleaseMode)
	} else {
		gin.SetMode(gin.DebugMode)
	}

	Cfg = &cfg
}
