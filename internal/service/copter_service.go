package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter"
	"github.com/YuraLk/teca_server/internal/dtos/copter/response_properties"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"

	"github.com/gin-gonic/gin"
)

type CopterService struct {
	C     *gin.Context
	Props copter.CalculateCopter
}

type ModeProperties struct {
	Props copter.CalculateCopter
	Calc  copter.StandartProperties
}

func (S CopterService) GetProperties() (copter.CopterResponse, error) {

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
		return copter.CopterResponse{}, err
	}

	// Вычисляем параметры окружающей среды
	envProps, envWarn := EnvironmentService{}.GetProperties(environment)

	// Вычисляем параметры батареи
	battProps, err := BatteryService{}.GetProperties(battery, composit)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return copter.CopterResponse{}, err
	}
	// Вычисляем параметры ESC
	escWarn := ControllerService{}.GetProperties(esc, battProps.BatteryVoltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := PropellerService{}.GetProperties(propeller, frame)

	// Вычисляем параметры мотора
	motorProps, motorWarn := MotorService{}.GetProperties(motor, frame)

	// Вычисление общих параметров для обоих режимов полета
	generalProps := GeneralService{}.GetProperties(frame, attachments, battProps.Mass, esc.Mass, propProps.Mass, motorProps.Mass)

	// Вычисляем параметры для режима зависания
	hoverProps, hoverWarn := ModeProperties{Props: S.Props, Calc: copter.StandartProperties{
		EnvironmentProperties: envProps,
		BatteryProperties:     battProps,
		PropellerProperties:   propProps,
		MotorProperties:       motorProps,
		GeneralProperties:     generalProps,
	}}.getHoverProperties()

	// Собираем предупреждения
	warnings := WarningService{}.AppendArrays(envWarn, escWarn, propWarn, motorWarn, hoverWarn)

	// Возвращаем расчитанные параметры
	var response copter.CopterResponse = copter.CopterResponse{
		CopterProperties: copter.CopterProperties{
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

func (S ModeProperties) getHoverProperties() (response_properties.HoverProperties, *[]dtos.WarningDto) {

	// Подъемная сила каждого пропеллера, необходимая для поддержания ЛА в воздухе, (Н):
	var PropellerHangingLift float64 = float64(S.Calc.GeneralProperties.Weight) / float64(S.Props.FrameProperties.PropellersNumber)

	// Частота вращения винта, (Гц)
	var PropellerSpeed float64 = math.Sqrt(PropellerHangingLift / (float64(S.Props.PropellerProperties.DimensionlessPowerConstant) * S.Calc.EnvironmentProperties.AirDensity * math.Pow(float64(S.Props.PropellerProperties.Diameter), 4)))

	warnings := WarningService{}.Append()

	return response_properties.HoverProperties{
		PropellerHangingLift: PropellerHangingLift,
		PropellerSpeed:       PropellerSpeed,
	}, warnings
}
