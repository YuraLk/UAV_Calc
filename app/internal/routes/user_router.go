package routes

import (
	"github.com/YuraLk/drone_calc/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type UserRouter struct {
	Prefix *gin.RouterGroup
}

func (S UserRouter) Router() {

	group := S.Prefix.Group("/user")

	group.PUT("/", middleware.AuthMiddleware(), func(ctx *gin.Context) {})
}
