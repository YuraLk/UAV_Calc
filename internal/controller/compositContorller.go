package controller

import (
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type CompositRequest struct {
	Name         string        `json:"name" binding:"required"`
	Voltage      types.Voltage `json:"voltage" binding:"required"`
	CRating      types.Current `json:"c_rating" binding:"required"`
	SafeCapacity uint8         `json:"safe_capacity" binding:"required"`
}

func GetComposits(c *gin.Context) {
	composits := service.GetComposits(c)
	c.JSON(200, &composits)
}

func CreateComposit(c *gin.Context) {
	var req CompositRequest
	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	Composit, err := service.CreateComposit(c, req.Name, req.Voltage, req.CRating, req.SafeCapacity)
	if err == nil {
		c.JSON(200, &Composit)
	}
}
