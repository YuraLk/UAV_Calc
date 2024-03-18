package exeptions

import (
	"net/http"

	"github.com/YuraLk/teca_server/internal/dtos"
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
}

func UnprocessableEntity(c *gin.Context, errors []dtos.ValidationErrorDto) {
	response := Error{
		Error:   "Ошибка валидации!",
		Details: errors,
	}
	// 422
	c.AbortWithStatusJSON(http.StatusUnprocessableEntity, response)
}

func BadRequest(c *gin.Context, msg string, err error) {
	response := Error{
		Error:   msg,
		Details: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusBadRequest, response)
}

func UnauthorizedError(c *gin.Context, err error) {
	response := Error{
		Error:   "Не авторизован!",
		Details: err.Error(),
	}
	c.AbortWithStatusJSON(http.StatusUnauthorized, response)
}

func RequestTimeout(c *gin.Context) {
	response := Error{
		Error: "Тайм-аут!",
	}
	c.AbortWithStatusJSON(http.StatusRequestTimeout, response)
}

func Forbidden(c *gin.Context) {
	response := Error{
		Error: "Доступ запрещен!",
	}
	c.AbortWithStatusJSON(http.StatusForbidden, response)
}

func NotFound(c *gin.Context, msg string) {
	response := Error{
		Error: msg,
	}
	c.AbortWithStatusJSON(http.StatusNotFound, response)
}
