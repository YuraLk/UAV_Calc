package routes

import (
	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/validators"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New() // Инициализируем роутер

	validators.Init()

	gin.SetMode(config.Cfg.Env)

	apiGroup := router.Group("/api") // Добавляем префикс /api

	UserRouter{Prefix: apiGroup}.Router() // Добавляем роуты пользователя
	CopterRouter{Prefix: apiGroup}.Router()
	CompositRouter{Prefix: apiGroup}.Router()

	return router
}
