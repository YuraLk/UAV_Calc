package validators

import (
	"context"
	"unicode"

	"github.com/YuraLk/drone_calc/backend/internal/database/postgres"
	"github.com/go-playground/validator/v10"
	"github.com/jackc/pgx/v5"
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

// Проверка уникальности почты
var isEmailUnique validator.Func = func(fl validator.FieldLevel) bool {
	email := fl.Field().String()

	// Проверка уникальности почты
	var email_is_exists bool
	if err := postgres.DB.QueryRow(context.Background(), "SELECT EXISTS(SELECT 1 FROM users WHERE email = $1) as is_exists", email).Scan(&email_is_exists); err != nil {
		if err != pgx.ErrNoRows {
			return false
		}
	}

	return !email_is_exists
}
