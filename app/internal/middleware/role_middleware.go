package middleware

import (
	"errors"
	"strings"

	"github.com/YuraLk/drone_calc/backend/internal/exeptions"
	"github.com/YuraLk/drone_calc/backend/internal/service"

	"github.com/gin-gonic/gin"
)

func RoleMiddleware(roles []string) gin.HandlerFunc {
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

		if data.UserDTO.Role != nil {
			for _, s := range roles {
				if s == "ADMIN" {
					c.Set("user", data.UserDTO)
					c.Next()
					break
				} else {
					exeptions.Forbidden(c)
				}
			}
		} else {
			exeptions.Forbidden(c)
		}
	}
}
