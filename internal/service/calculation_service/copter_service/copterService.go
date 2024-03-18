package copter_service

import (
	// "fmt"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	requests_properties "github.com/YuraLk/teca_server/internal/requests"
	responses_properties "github.com/YuraLk/teca_server/internal/responses"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/properties_service"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"

	"github.com/gin-gonic/gin"
)

func CalculateCopterProperties(c *gin.Context, req requests_properties.CalculateCopter) (responses_properties.CopterResponse, error) {

	// Навесное оборудование
	// var attachments = req.AttachmentsProperties
	// ESC
	var esc = req.ControllerProperties
	// Внешняя среда
	var environment = req.EnvironmentProperties
	// Мотор
	// var motor = req.MotorProperties
	// Рама
	// var frame = req.FrameProperties
	// Пропеллер
	var propeller = req.PropellerProperties
	// Аккумулятор
	var battery = req.BatteryProperties

	// Ищем композит аккумулятора с ВАХ
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", battery.CompositId).First(&composit).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return responses_properties.CopterResponse{}, err
	}

	// Вычисляем параметры окружающей среды
	envProps, envWarn := properties_service.GetEnvironmentProperties(environment)

	// Вычисляем параметры батареи
	battProps, err := properties_service.GetBatteryProperties(battery, composit)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return responses_properties.CopterResponse{}, err
	}
	// Вычисляем параметры ESC
	escWarn := properties_service.GetControllerProperties(esc, battProps.BatteryVoltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := properties_service.GetPropellerProperties(propeller)

	// Собираем предупреждения
	warnings := warning_service.AppendWarningsArrays(envWarn, escWarn, propWarn)

	// Возвращаем расчитанные параметры
	var response responses_properties.CopterResponse = responses_properties.CopterResponse{
		CopterProperties: responses_properties.CopterProperties{
			EnvironmentProperties: envProps,
			BatteryProperties:     battProps,
			PropellerProperties:   propProps,
		},
		Warings: warnings,
	}

	return response, nil
}
