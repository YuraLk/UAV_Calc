package controller

import (
	"time"

	"github.com/YuraLk/drone_calc/backend/internal/configs"
	"github.com/YuraLk/drone_calc/backend/internal/dtos/auth/request_properties"
	"github.com/YuraLk/drone_calc/backend/internal/exeptions"
	"github.com/YuraLk/drone_calc/backend/internal/service"
	"github.com/YuraLk/drone_calc/backend/internal/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct{}

func (AuthController) Auth(c *gin.Context) {
	var req request_properties.LoginDTO

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	// dto, tokens, err := service.UserService{C: c}.Auth(req.Email, req.Password, req.Device)
	// if err == nil {
	// 	// Помещаем refresh в http cookie
	// 	c.SetCookie("refresh_token", tokens.Refresh, int(time.Hour.Seconds()*24*14), "/", configs.Cfg.HTTPServer.Host, false, true)

	// 	// Возвращаем DTO и токены
	// 	c.JSON(200, gin.H{"user": &dto, "accessToken": tokens.Access})
	// }
}

// Функция регистрации
func (AuthController) Register(ctx *gin.Context) {
	var req request_properties.RegisterDTO

	// Проверка валидации
	if err := ctx.ShouldBind(&req); err != nil {
		errors := utils.FormatJSONErrors(err.Error())
		exeptions.UnprocessableEntity(ctx, errors)
		return
	}

	// Извлекаем данные из тела application/json
	ctx.ShouldBindJSON(&req)

	dto, tokens, err := service.AuthService{C: ctx}.Register(ctx, req.Name, req.Surname, req.Patronymic, req.Email, req.Phone, req.Password, req.Device)
	if err == nil {
		// Помещаем refresh в http cookie
		ctx.SetCookie("refresh_token", tokens.Refresh, int(time.Hour.Seconds()*24*14), "/", configs.Cfg.HTTPServer.Host, false, true)

		ctx.JSON(200, gin.H{"user": &dto, "access_token": tokens.Access})
	}
}

// Обновлние токена при сгорании старого
func (AuthController) Refresh(c *gin.Context) {
	// Получаем Refresh - токен
	// refreshToken, err := c.Cookie("refreshToken")
	// if err != nil {
	// 	exeptions.UnauthorizedError(c, err)
	// 	return
	// }

	// dto, tokens, err := service.UserService{C: c}.Refresh(refreshToken)
	// if err == nil {
	// 	// Помещаем новый Refresh - токен в Cookie
	// 	c.SetCookie("refresh_token", tokens.Refresh, int(time.Hour.Seconds()*24*30), "/", configs.Cfg.HTTPServer.Host, false, true)

	// 	// Возвращаем DTO и токены
	// 	c.JSON(200, gin.H{"user": &dto, "access_token": tokens.Access})
	// }

}

func (AuthController) Logout(ctx *gin.Context) {
	// Получаем Refresh - токен
	refreshToken, err := ctx.Cookie("refresh_token")
	if err != nil {
		exeptions.UnauthorizedError(ctx, err)
		return
	}

	err = service.AuthService{}.Logout(ctx, refreshToken)
	if err == nil {
		// Устанавливаем cookie с нулевым сроком действия (удаляем)
		ctx.SetCookie("refresh_token", "", 0, "/", configs.Cfg.HTTPServer.Host, false, true)
		ctx.JSON(200, gin.H{})
	}
}
