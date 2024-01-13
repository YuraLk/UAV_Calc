package main

import (
	"github.com/YuraLk/drone_calc/internal/config"
	"github.com/YuraLk/drone_calc/internal/database/postgres"
	"github.com/YuraLk/drone_calc/internal/routes"
)

func main() {

	// Загружаем параметры из файла конфигурации
	config.Load()

	// Инициялизируем маршруты
	router := routes.Init()

	// Подключаемся к базе данных
	postgres.Connect()

	router.Run(config.Cfg.HTTPServer.Port)
}
