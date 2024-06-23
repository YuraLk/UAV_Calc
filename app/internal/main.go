package main

import (
	"github.com/YuraLk/drone_calc/backend/internal/configs"
	"github.com/YuraLk/drone_calc/backend/internal/database/postgres"
	"github.com/YuraLk/drone_calc/backend/internal/routes"
)

func main() {

	// Загружаем параметры из файла конфигурации
	configs.Load()

	// Инициялизируем маршруты
	router := routes.Init()

	// Подключаемся к базе данных
	postgres.Connect()

	router.Run(configs.Cfg.HTTPServer.Port)
}
