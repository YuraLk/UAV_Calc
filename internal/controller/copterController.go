package controller

import (
	request "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request"
	"github.com/YuraLk/teca_server/internal/exeptions"
	copter_service "github.com/YuraLk/teca_server/internal/service/calculation_service/copter_service"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

func CalculateCopterProperties(c *gin.Context) {
	var req request.CalculateCopter

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	properties, err := copter_service.CalculateCopterProperties(c, req)
	if err == nil {
		c.JSON(200, properties)
	}
}
