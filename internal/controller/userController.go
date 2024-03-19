package controller

import (
	"errors"
	"time"

	"github.com/YuraLk/teca_server/internal/config"
	"github.com/YuraLk/teca_server/internal/dtos/user_dtos"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/utils"

	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (UserController) Get(c *gin.Context) {
	users := service.UserService{C: c}.Get()
	c.JSON(200, &users)
}

func (UserController) Auth(c *gin.Context) {
	var req user_dtos.Login

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.UserService{C: c}.Auth(req.Email, req.Password, req.Device)
	if err == nil {
		// Помещаем refresh в http cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

// Функция регистрации
func (UserController) Register(c *gin.Context) {
	var req user_dtos.Register

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.UserService{C: c}.Register(req.Name, req.Email, req.Phone, req.Password, req.Device)
	if err == nil {
		// Помещаем refresh в http cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

// Обновлние токена при сгорании старого
func (UserController) Refresh(c *gin.Context) {
	// Получаем Refresh - токен
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return
	}

	dto, tokens, err := service.UserService{C: c}.Refresh(refreshToken)
	if err == nil {
		// Помещаем новый Refresh - токен в Cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}

}

func (UserController) UpdateUser(c *gin.Context) {
	var req user_dtos.UpdateUser

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
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
	data, ok := decodedData.(user_dtos.UserDTO)
	if !ok {
		exeptions.InternalServerError(c, errors.New("error type assignment"))
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	dto, tokens, err := service.UserService{}.UpdateUser(data.Id, req.Name, req.Email, req.Phone, req.Device)
	if err == nil {
		// Помещаем новый Refresh - токен в Cookie
		c.SetCookie("refreshToken", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", config.Cfg.HTTPServer.Host, false, true)

		// Возвращаем DTO и токены
		c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	}
}

func (UserController) Logout(c *gin.Context) {
	// Получаем Refresh - токен
	refreshToken, err := c.Cookie("refreshToken")
	if err != nil {
		exeptions.UnauthorizedError(c, err)
		return
	}

	err = service.UserService{}.Logout(refreshToken)
	if err == nil {
		// Устанавливаем cookie с нулевым сроком действия (удаляем)
		c.SetCookie("refreshToken", "", 0, "/", config.Cfg.HTTPServer.Host, false, true)
		c.JSON(200, gin.H{})
	}
}
