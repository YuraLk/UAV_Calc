package routes

import (
	"github.com/YuraLk/teca_server/internal/controller"
	"github.com/YuraLk/teca_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func CopterRouter(router *gin.RouterGroup) {

	group := router.Group("/copter")

	group.POST("/", middleware.AuthMiddleware(), controller.CalculateCopterProperties)
}
