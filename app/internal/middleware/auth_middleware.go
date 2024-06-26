package middleware

import (
	"errors"
	"strings"

	"github.com/YuraLk/drone_calc/backend/internal/exeptions"
	"github.com/YuraLk/drone_calc/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authorizationHeader := c.GetHeader("Authorization")

		if authorizationHeader == "" {
			exeptions.UnauthorizedError(c, errors.New("header value not found"))
			return
		}

		token := strings.Split(authorizationHeader, " ")[1]

		data, err := service.TokenService{}.ValidateAccess(token)
		if err != nil {
			exeptions.UnauthorizedError(c, err)
			return
		}

		c.Set("user", data.AuthDTO)

		c.Next()
	}
}
