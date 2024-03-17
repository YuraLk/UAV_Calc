package service

import (
	// "fmt"

	requests "github.com/YuraLk/teca_server/internal/requests"
	responses "github.com/YuraLk/teca_server/internal/responses"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/gin-gonic/gin"
)

func CalculateCopterProperties(c *gin.Context, req requests.CalculateCopter) (responses.CopterResponse, error) {

	// Навесное оборудование
	// var attachments = req.AttachmentsProperties
	// ESC
	// var esc = req.ControllerProperties
	// Внешняя среда
	var environment = req.EnvironmentProperties

	// Тестовый вывод
	// fmt.Println(attachments)
	// fmt.Println(esc)
	// fmt.Println(environment)

	// Помещаем получаемые предупреждения в один массив
	var warnings []types.Warning

	// Вычисляем параметры окружающей среды
	envProps, envWarn := GetEnvironmentProperties(environment)
	if envWarn != nil {
		warnings = append(warnings, *envWarn...)
	}

	// Возвращаем расчитанные параметры
	var response responses.CopterResponse = responses.CopterResponse{
		CopterProperties: responses.CopterProperties{
			EnvironmentProperties: envProps,
		},
		Warings: warnings,
	}

	return response, nil
}
