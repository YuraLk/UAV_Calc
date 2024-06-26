package routes

import (
	"github.com/YuraLk/drone_calc/backend/internal/controller"
	"github.com/YuraLk/drone_calc/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type AuthRouter struct {
	Prefix *gin.RouterGroup
}

func (S AuthRouter) Router() {

	group := S.Prefix.Group("/auth")

	group.POST("/register", controller.AuthController{}.Register)

	group.POST("/login", controller.AuthController{}.Auth)

	group.POST("/logout", middleware.AuthMiddleware(), controller.AuthController{}.Logout)

	group.GET("/refresh", controller.AuthController{}.Refresh)
}
