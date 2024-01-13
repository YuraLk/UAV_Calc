package utils

import (
	"regexp"
	"strings"
)

type ValidationError struct {
	Key   string
	Field string
	Error string
	Tag   string
}

func FormatErrors(input string) []ValidationError {
	// Регулярное выражение для извлечения значений
	re := regexp.MustCompile(`Key: '(?P<Key>[^']+)' Error:Field validation for '(?P<Field>[^']+)' failed on the '(?P<Tag>[^']+)' tag`)

	// Разбиваем входную строку на строки, если ошибок несколько
	lines := strings.Split(input, "\n")

	// Создаем массив для объектов ValidationError
	var errors []ValidationError

	// Обрабатываем каждую строку
	for _, line := range lines {
		match := re.FindStringSubmatch(line)
		if match != nil {
			errorObj := ValidationError{
				Key:   match[1],
				Field: match[2],
				Tag:   match[3],
				Error: setErrorMessage(match[2], match[3]),
			}
			errors = append(errors, errorObj)
		}
	}

	return errors
}

// По тегу выставляем ошибку валидации
func setErrorMessage(field string, tag string) string {
	var message string
	switch field {
	case "Email":
		switch tag {
		case "email":
			message = "Неверная почта!"
			break
		}
		break
	}

	return message
}
