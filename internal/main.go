package main

import (
	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/routes"
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
