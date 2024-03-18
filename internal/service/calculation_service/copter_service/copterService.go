package copter_service

import (
	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/properties_service"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"

	"github.com/gin-gonic/gin"
)

func CalculateCopterProperties(c *gin.Context, props request.CalculateCopter) (response.CopterResponse, error) {

	// Навесное оборудование
	var attachments = props.AttachmentsProperties
	// ESC
	var esc = props.ControllerProperties
	// Внешняя среда
	var environment = props.EnvironmentProperties
	// Мотор
	var motor = props.MotorProperties
	// Рама
	var frame = props.FrameProperties
	// Пропеллер
	var propeller = props.PropellerProperties
	// Аккумулятор
	var battery = props.BatteryProperties

	// Ищем композит аккумулятора с ВАХ
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", battery.CompositId).First(&composit).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return response.CopterResponse{}, err
	}

	// Вычисляем параметры окружающей среды
	envProps, envWarn := properties_service.GetEnvironmentProperties(environment)

	// Вычисляем параметры батареи
	battProps, err := properties_service.GetBatteryProperties(battery, composit)
	if err != nil {
		exeptions.InternalServerError(c, err)
		return response.CopterResponse{}, err
	}
	// Вычисляем параметры ESC
	escWarn := properties_service.GetControllerProperties(esc, battProps.BatteryVoltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := properties_service.GetPropellerProperties(propeller, frame)

	// Вычисляем параметры мотора
	motorProps, motorWarn := properties_service.GetMotorProperties(motor, frame)

	// Вычисление общих параметров для обоих режимов полета
	generalProps := properties_service.GetGeneralProperties(frame, attachments, battProps.Mass, esc.Mass, propProps.Mass, motorProps.Mass)

	// Вычисляем параметры для режима зависания

	// Собираем предупреждения
	warnings := warning_service.AppendWarningsArrays(envWarn, escWarn, propWarn, motorWarn)

	// Возвращаем расчитанные параметры
	var response response.CopterResponse = response.CopterResponse{
		CopterProperties: response.CopterProperties{
			EnvironmentProperties: envProps,
			BatteryProperties:     battProps,
			PropellerProperties:   propProps,
			MotorProperties:       motorProps,
			GeneralProperties:     generalProps,
		},
		Warings: warnings,
	}

	return response, nil
}
