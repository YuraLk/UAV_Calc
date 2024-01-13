package config

import (
	"flag"
	"log"
	"os"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	Env        string `yaml:"env" env-default:"release"`
	HTTPServer HTTPServer
	Database   Database
	JWT        JWT
}

type HTTPServer struct {
	Host string `yaml:"host" env-default:"localhost"`
	Port string `yaml:"port" env-default:":8080"`
}

type Database struct {
	Host     string `yaml:"host" env-default:"localhost"`
	Port     string `yaml:"port" env-default:":5432"`
	User     string `yaml:"user" env-default:"postgres"`
	Password string `yaml:"password" env-default:"1427"`
	Name     string `yaml:"name" env-default:"drone_calc"`
	SSLMode  string `yaml:"sslmode" env-default:"disable"`
}

type JWT struct {
	JWTAccessKey  string `yaml:"jwt_access_key"`
	JWTRefreshKey string `yaml:"jwt_refresh_key"`
}

var Cfg *Config

func Load() {
	// Парсим флаг пути к конфигурационному файлу при запуске
	p := flag.String("file", "configs/config.yaml", "Specify the configuration file!")

	flag.Parse()

	configPath := *p

	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Fatalf("Configuration file does not exist: %s", configPath)
	}

	var cfg Config

	if err := cleanenv.ReadConfig(configPath, &cfg); err != nil {
		log.Fatalf("Cannot read configuration file: %s", err)
	}

	Cfg = &cfg
}
