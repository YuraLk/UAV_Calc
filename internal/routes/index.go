package routes

import (
	"github.com/YuraLk/drone_calc/internal/config"
	"github.com/YuraLk/drone_calc/internal/validators"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New() // Инициализируем роутер

	validators.Init()

	gin.SetMode(config.Cfg.Env)

	apiGroup := router.Group("/api") // Добавляем префикс /api

	UserRouter(apiGroup) // Добавляем роуты пользователя
	CalculationRouter(apiGroup)

	return router
}
