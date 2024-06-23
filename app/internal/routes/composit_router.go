package routes

import (
	"github.com/YuraLk/drone_calc/backend/internal/controller"
	"github.com/YuraLk/drone_calc/backend/internal/middleware"
	"github.com/gin-gonic/gin"
)

type CompositRouter struct {
	Prefix *gin.RouterGroup
}

func (S CompositRouter) Router() {

	group := S.Prefix.Group("/composit")

	group.GET("/", middleware.AuthMiddleware(), controller.CompositController{}.Get)
	group.POST("/", middleware.RoleMiddleware([]string{"ADMIN"}), controller.CompositController{}.Create)
	group.PUT("/", middleware.RoleMiddleware([]string{"ADMIN"}), controller.CompositController{}.Update)
	group.DELETE("/:id", middleware.RoleMiddleware([]string{"ADMIN"}), controller.CompositController{}.Delete)
}
