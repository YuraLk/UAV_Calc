package controller

import (
	"github.com/YuraLk/teca_server/internal/dtos"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

func GetComposits(c *gin.Context) {
	composits, err := service.GetComposits(c)
	if err == nil {
		c.JSON(200, &composits)
	}
}

func CreateComposit(c *gin.Context) {
	// Название композита
	name := c.PostForm("name")
	// Таблица с ВАХ аккумулятора
	file, err := c.FormFile("file")

	// Проверка валидности данных из FormData
	var errors = utils.BindFormData("Composit", []dtos.BindingDto{
		{
			Key:   "Name",
			Value: name,
			Error: nil,
		},
		{
			Key:   "File",
			Value: file,
			Error: err,
		},
	})

	if len(errors) > 0 {
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Передаем проверенные данные в сервис
	Composit, err := service.CreateComposit(c, name, file)
	if err == nil {
		c.JSON(200, &Composit)
	}
}

func UpdateComposit(c *gin.Context) {
	id := c.PostForm("id")
	name := c.PostForm("name")
	// Таблица с ВАХ аккумулятора
	file, err := c.FormFile("file")

	// Проверка валидности данных из FormData
	var errors = utils.BindFormData("Composit", []dtos.BindingDto{
		{
			Key:   "Id",
			Value: id,
			Error: nil,
		},
		{
			Key:   "Name",
			Value: name,
			Error: nil,
		},
		{
			Key:   "File",
			Value: file,
			Error: err,
		},
	})

	if len(errors) > 0 {
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Передаем проверенные данные в сервис

	Composit, err := service.UpdateComposit(c, id, name, file)
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
