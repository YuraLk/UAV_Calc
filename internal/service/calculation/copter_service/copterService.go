package copter_service

import (
	"github.com/YuraLk/teca_server/internal/database/postgres"
	copter_dtos "github.com/YuraLk/teca_server/internal/dtos/copter_dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/service/calculation/copter_service/hover_service"
	"github.com/YuraLk/teca_server/internal/service/calculation/properties_service"
	"github.com/YuraLk/teca_server/internal/service/calculation/warning_service"

	"github.com/gin-gonic/gin"
)

type CopterService struct {
	C     *gin.Context
	Props request.CalculateCopter
}

func (S CopterService) CopterProperties() (response.CopterResponse, error) {

	// Навесное оборудование
	var attachments = S.Props.AttachmentsProperties
	// ESC
	var esc = S.Props.ControllerProperties
	// Внешняя среда
	var environment = S.Props.EnvironmentProperties
	// Мотор
	var motor = S.Props.MotorProperties
	// Рама
	var frame = S.Props.FrameProperties
	// Пропеллер
	var propeller = S.Props.PropellerProperties
	// Аккумулятор
	var battery = S.Props.BatteryProperties

	// Ищем композит аккумулятора с ВАХ
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", battery.CompositId).First(&composit).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return response.CopterResponse{}, err
	}

	// Вычисляем параметры окружающей среды
	envProps, envWarn := properties_service.GetEnvironmentProperties(environment)

	// Вычисляем параметры батареи
	battProps, err := properties_service.GetBatteryProperties(battery, composit)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
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
	hoverProps, hoverWarn := hover_service.GetHoverProperties(S.Props, copter_dtos.StandartProperties{
		EnvironmentProperties: envProps,
		BatteryProperties:     battProps,
		PropellerProperties:   propProps,
		MotorProperties:       motorProps,
		GeneralProperties:     generalProps,
	})

	// Собираем предупреждения
	warnings := warning_service.AppendWarningsArrays(envWarn, escWarn, propWarn, motorWarn, hoverWarn)

	// Возвращаем расчитанные параметры
	var response response.CopterResponse = response.CopterResponse{
		CopterProperties: response.CopterProperties{
			EnvironmentProperties: envProps,
			BatteryProperties:     battProps,
			PropellerProperties:   propProps,
			MotorProperties:       motorProps,
			GeneralProperties:     generalProps,
			HoverProperties:       hoverProps,
		},
		Warings: warnings,
	}

	return response, nil
}
