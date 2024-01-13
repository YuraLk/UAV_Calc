package routes

import (
	"github.com/YuraLk/drone_calc/internal/controller"
	"github.com/YuraLk/drone_calc/internal/middleware"
	"github.com/gin-gonic/gin"
)

func UserRouter(router *gin.RouterGroup) {

	group := router.Group("/user")

	group.POST("/register", controller.Register)

	group.POST("/auth", controller.Auth)

	group.POST("/logout", middleware.AuthMiddleware(), controller.Logout)

	group.GET("/", middleware.RoleMiddleware([]string{"ADMIN"}), controller.GetUsers)

	group.GET("/refresh", controller.Refresh)

	group.PUT("/", middleware.AuthMiddleware(), controller.UpdateUser)
}
