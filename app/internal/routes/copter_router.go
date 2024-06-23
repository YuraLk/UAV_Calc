package routes

import (
	"github.com/YuraLk/drone_calc/backend/internal/controller"
	"github.com/YuraLk/drone_calc/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type CopterRouter struct {
	Prefix *gin.RouterGroup
}

func (S CopterRouter) Router() {

	group := S.Prefix.Group("/copter")

	group.POST("/", middleware.AuthMiddleware(), controller.CopterController{}.Get)
}
