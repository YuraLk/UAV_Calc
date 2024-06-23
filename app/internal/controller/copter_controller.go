package controller

import (
	"github.com/YuraLk/drone_calc/backend/internal/dtos/copter"
	"github.com/YuraLk/drone_calc/backend/internal/exeptions"
	"github.com/YuraLk/drone_calc/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type CopterController struct{}

func (CopterController) Get(c *gin.Context) {
	var req copter.CalculateCopter

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	// properties, err := service.CopterService{C: c, Props: req}.GetProperties()
	// if err == nil {
	// 	c.JSON(200, properties)
	// }
}
