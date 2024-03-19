package routes

import (
	"github.com/YuraLk/teca_server/internal/controller"
	"github.com/YuraLk/teca_server/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Prefix *gin.RouterGroup
}

func (S UserRouter) Router() {

	group := S.Prefix.Group("/user")

	group.POST("/register", controller.UserController{}.Register)

	group.POST("/auth", controller.UserController{}.Auth)

	group.POST("/logout", middleware.AuthMiddleware(), controller.UserController{}.Logout)

	group.GET("/", middleware.RoleMiddleware([]string{"ADMIN"}), controller.UserController{}.Get)

	group.GET("/refresh", controller.UserController{}.Refresh)

	group.PUT("/", middleware.AuthMiddleware(), controller.UserController{}.UpdateUser)
}
