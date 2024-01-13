package routes

import (
	"github.com/YuraLk/drone_calc/internal/controller"
	"github.com/YuraLk/drone_calc/internal/middleware"
	"github.com/gin-gonic/gin"
)

func CalculationRouter(router *gin.RouterGroup) {

	group := router.Group("/calculation")

	group.POST("/", middleware.AuthMiddleware(), controller.Calculate)
}
