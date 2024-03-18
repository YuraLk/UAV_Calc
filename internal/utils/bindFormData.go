package utils

import (
	"github.com/YuraLk/teca_server/internal/dtos"
)

var BindFormData = func(area string, args []dtos.BindingDto) []dtos.ValidationErrorDto {
	// Создаем массив для объектов ValidationError
	var errors []dtos.ValidationErrorDto

	for _, arg := range args {

		// Проверка для файлов
		if arg.Error != nil {
			errorObj := dtos.ValidationErrorDto{
				Key:   area + "." + arg.Key,
				Field: arg.Key,
				Tag:   "required",
				Error: arg.Error.Error(),
			}
			errors = append(errors, errorObj)
			continue
		}

		switch v := arg.Value.(type) {
		case string:
			if v == "" {
				errorObj := dtos.ValidationErrorDto{
					Key:   area + "." + arg.Key,
					Field: arg.Key,
					Tag:   "required",
					Error: "Строка не может быть пустой",
				}
				errors = append(errors, errorObj)
			}
		}
	}

	return errors
}
