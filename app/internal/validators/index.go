package validators

import (
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func Init() {
	// Регистрация кастомных функций валидации
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("isSafety", isSafety)
		v.RegisterValidation("isEmailUnique", isEmailUnique)
	}
}
