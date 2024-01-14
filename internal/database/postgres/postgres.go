package postgres

import (
	"fmt"

	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Connect() {

	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s", config.Cfg.Database.Host, config.Cfg.Database.User, config.Cfg.Database.Password, config.Cfg.Database.Name, config.Cfg.Database.Port, config.Cfg.Database.SSLMode)

	db, err := gorm.Open(postgres.Open(dsn))

	if err != nil {
		panic(err)
	}

	db.AutoMigrate(&models.User{}, &models.Session{}, &models.Access{})
	DB = db

}
