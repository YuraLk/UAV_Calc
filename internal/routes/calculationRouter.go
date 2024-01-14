package routes

import (
	"github.com/YuraLk/teca_server/internal/controller"
	"github.com/YuraLk/teca_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func CalculationRouter(router *gin.RouterGroup) {

	group := router.Group("/calculation")

	group.POST("/", middleware.AuthMiddleware(), controller.Calculate)
}
