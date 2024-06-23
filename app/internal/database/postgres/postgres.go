package postgres

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"

	"github.com/YuraLk/drone_calc/backend/internal/configs"
)

var DB *pgx.Conn

func Connect() {

	dsn := fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s", configs.Cfg.Database.User, configs.Cfg.Database.Password, configs.Cfg.Database.Host, configs.Cfg.Database.Port, configs.Cfg.Database.Name, configs.Cfg.Database.SSLMode)

	db, err := pgx.Connect(context.Background(), dsn)

	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}

	DB = db

	defer db.Close(context.Background())
}
