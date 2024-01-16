package routes

import (
	"github.com/YuraLk/teca_server/internal/controller"
	"github.com/YuraLk/teca_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

func CompositRouter(router *gin.RouterGroup) {

	group := router.Group("/composit")

	group.GET("/", middleware.AuthMiddleware(), controller.GetComposits)
	group.POST("/", middleware.RoleMiddleware([]string{"ADMIN"}), controller.CreateComposit)
}
