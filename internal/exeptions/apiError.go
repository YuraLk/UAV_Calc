package exeptions

import (
	"net/http"

	"github.com/YuraLk/drone_calc/internal/utils"
	"github.com/gin-gonic/gin"
)

type Error struct {
	Error   string
	Details interface{}
}

func InternalServerError(c *gin.Context, err error) {
	response := Error{
		Error:   "Внутренняя ошибка сервера!",
		Details: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusInternalServerError, response)
	return
}

func UnprocessableEntity(c *gin.Context, errors []utils.ValidationError) {
	response := Error{
		Error:   "Ошибка валидации!",
		Details: errors,
	}
	// 422
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
	return
}

func BadRequest(c *gin.Context, msg string, err error) {
	response := Error{
		Error:   msg,
		Details: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
	return
}

func UnauthorizedError(c *gin.Context, err error) {
	response := Error{
		Error:   "Не авторизован!",
		Details: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
	return
}

func RequestTimeout(c *gin.Context) {
	response := Error{
		Error: "Тайм-аут!",
	}
	c.AbortWithStatusJSON(http.StatusRequestTimeout, response)
	return
}

func Forbidden(c *gin.Context) {
	response := Error{
		Error: "Доступ запрещен!",
	}
	c.AbortWithStatusJSON(http.StatusForbidden, response)
	return
}

func NotFound(c *gin.Context, msg string) {
	response := Error{
		Error: msg,
	}
	c.AbortWithStatusJSON(http.StatusNotFound, response)
	return
}
