package controller

import (
	"errors"
	"time"

	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/dtos"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type RegisterRequest struct {
	Name     string `json:"name" binding:"required,fullName"`
	Email    string `json:"email" binding:"required,email"`
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password" binding:"required,min=10,max=40,safety"`
	Device   string `json:"device" binding:"required"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
	Device   string `json:"device" binding:"required"`
}

type UpdateUserRequest struct {
	Name   string `json:"name" binding:"required,fullName"`
	Email  string `json:"email" binding:"required,email"`
	Phone  string `json:"phone" binding:"required"`
	Device string `json:"device" binding:"required"`
}

func GetUsers(c *gin.Context) {
	users := service.GetUsers(c)
	c.JSON(200, &users)
}

func Auth(c *gin.Context) {
	var req LoginRequest

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.Auth(c, req.Email, req.Password, req.Device)
	if err == nil {
		// Помещаем refresh в http cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

// Функция регистрации
func Register(c *gin.Context) {
	var req RegisterRequest

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.Register(c, req.Name, req.Email, req.Phone, req.Password, req.Device)
	if err == nil {
		// Помещаем refresh в http cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

// Обновлние токена при сгорании старого
func Refresh(c *gin.Context) {
	// Получаем Refresh - токен
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return
	}

	dto, tokens, err := service.Refresh(c, refreshToken)
	if err == nil {
		// Помещаем новый Refresh - токен в Cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}

}

func UpdateUser(c *gin.Context) {
	var req UpdateUserRequest

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Из контекста достаем декодированныые данные
	decodedData, exists := c.Get("user")
	if !exists {
		exeptions.InternalServerError(c, errors.New("decoded data not found"))
		return
	}

	// Присваиваем декодированным данным из контекста тип DTO
	data, ok := decodedData.(dtos.UserDTO)
	if !ok {
		exeptions.InternalServerError(c, errors.New("error type assignment"))
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.UpdateUser(c, data.Id, req.Name, req.Email, req.Phone, req.Device)
	if err == nil {
		// Помещаем новый Refresh - токен в Cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

func Logout(c *gin.Context) {
	// Получаем Refresh - токен
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return
	}

	if err := service.Logout(c, refreshToken); err == nil {
		// Устанавливаем cookie с нулевым сроком действия (удаляем)
		c.SetCookie("refreshToken", "", 0, "/", config.Cfg.HTTPServer.Host, false, true)
		c.JSON(200, gin.H{})
	}
}
