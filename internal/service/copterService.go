package service

import (
	// "fmt"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
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
	// Мотор
	// var motor = req.MotorProperties
	// Рама
	// var frame = req.FrameProperties
	// Пропеллер
	// var propeller = req.PropellerProperties
	// Аккумулятор
	var battery = req.BatteryProperties

	// Ищем композит аккумулятора с ВАХ
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", battery.CompositId).First(&composit).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return responses.CopterResponse{}, err
	}

	// DEBUG
	// fmt.Println(attachments)
	// fmt.Println(esc)
	// fmt.Println(environment)
	// fmt.Println(motor)
	// fmt.Println(frame)
	// fmt.Println(propeller)
	// fmt.Println(battery)
	// fmt.Println(composit)

	// Помещаем получаемые предупреждения в один массив
	var warnings []types.Warning

	// Вычисляем параметры окружающей среды
	envProps, envWarn := GetEnvironmentProperties(environment)
	if envWarn != nil {
		warnings = append(warnings, *envWarn...)
	}

	battProps, battWarn := GetBatteryProperties(battery, composit)
	if battWarn != nil {
		warnings = append(warnings, *battWarn...)
	}

	// Возвращаем расчитанные параметры
	var response responses.CopterResponse = responses.CopterResponse{
		CopterProperties: responses.CopterProperties{
			EnvironmentProperties: envProps,
			BatteryProperties:     battProps,
		},
		Warings: warnings,
	}

	return response, nil
}
