package service

import (
	// "fmt"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	requests_properties "github.com/YuraLk/teca_server/internal/requests"
	responses_properties "github.com/YuraLk/teca_server/internal/responses"
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
	envProps, envWarn := GetEnvironmentProperties(environment)

	// Вычисляем параметры батареи
	battProps, err := GetBatteryProperties(battery, composit)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return responses_properties.CopterResponse{}, err
	}
	// Вычисляем параметры ESC
	escWarn := GetControllerProperties(esc, battProps.BatteryVoltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := GetPropellerProperties(propeller)

	// Собираем предупреждения
	warnings := AppendWarningsArrays(envWarn, escWarn, propWarn)

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
