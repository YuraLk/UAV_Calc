package utils

import "github.com/YuraLk/teca_server/internal/types"

var BindFormData = func(area string, args []types.BindingScruct) []types.ValidationError {
	// Создаем массив для объектов ValidationError
	var errors []types.ValidationError

	for _, arg := range args {

		// Проверка для файлов
		if arg.Error != nil {
			errorObj := types.ValidationError{
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
				errorObj := types.ValidationError{
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
