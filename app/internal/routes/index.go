package routes

import (
	"github.com/YuraLk/drone_calc/backend/internal/configs"
	"github.com/YuraLk/drone_calc/backend/internal/validators"
	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	router := gin.New() // Инициализируем роутер

	validators.Init()

	gin.SetMode(configs.Cfg.Env)

	apiGroup := router.Group("/api") // Добавляем префикс /api

	AuthRouter{Prefix: apiGroup}.Router()
	UserRouter{Prefix: apiGroup}.Router() // Добавляем роуты пользователя
	CopterRouter{Prefix: apiGroup}.Router()
	CompositRouter{Prefix: apiGroup}.Router()

	return router
}
