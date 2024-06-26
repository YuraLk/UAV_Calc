package validators

import (
	"unicode"

	"github.com/go-playground/validator/v10"
)

var isSafety validator.Func = func(fl validator.FieldLevel) bool {
	password := fl.Field().String()

	// Проверка наличия хотя бы одной буквы в верхнем регистре
	hasUppercase := false
	// Проверка наличия хотя бы одной буквы в нижнем регистре
	hasLowercase := false
	// Проверка наличия хотя бы одной специальной буквы
	hasSpecialChar := false

	for _, char := range password {
		switch {
		case unicode.IsUpper(char):
			hasUppercase = true
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsPunct(char) || unicode.IsSymbol(char):
			hasSpecialChar = true
		}
	}

	if !hasUppercase || !hasLowercase || !hasSpecialChar {
		return false
	}

	return true
}
