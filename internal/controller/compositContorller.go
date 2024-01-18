package controller

import (
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type CreateCompositRequest struct {
	Name         string        `json:"name" binding:"required"`
	Voltage      types.Voltage `json:"voltage" binding:"required"`
	CRating      types.Current `json:"c_rating" binding:"required"`
	SafeCapacity float32       `json:"safe_capacity" binding:"required"`
}

type UpdateCompositRequest struct {
	Id           uint          `json:"ID" binding:"required"`
	Name         string        `json:"name" binding:"required"`
	Voltage      types.Voltage `json:"voltage" binding:"required"`
	CRating      types.Current `json:"c_rating" binding:"required"`
	SafeCapacity float32       `json:"safe_capacity" binding:"required"`
}

func GetComposits(c *gin.Context) {
	composits := service.GetComposits(c)
	c.JSON(200, &composits)
}

func CreateComposit(c *gin.Context) {
	var req CreateCompositRequest
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

func UpdateComposit(c *gin.Context) {
	var req UpdateCompositRequest
	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	Composit, err := service.UpdateComposit(c, req.Id, req.Name, req.Voltage, req.CRating, req.SafeCapacity)
	if err == nil {
		c.JSON(200, &Composit)
	}
}

func DeleteComposit(c *gin.Context) {
	id := c.Param("id")

	if err := service.DeleteComposit(c, id); err == nil {
		c.JSON(200, gin.H{})
	}
}
