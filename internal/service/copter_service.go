package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/consts"
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

	// Ищем композит аккумулятора с ВАХ
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", S.Props.BatteryProperties.CompositId).First(&composit).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return copter.CopterResponse{}, err
	}

	// Вычисляем параметры окружающей среды
	envProps, envWarn := EnvironmentService{}.GetProperties(S.Props.EnvironmentProperties)

	// Вычисляем параметры батареи
	battProps, err := BatteryService{}.GetProperties(S.Props.BatteryProperties, composit)
	if err != nil {
		exeptions.InternalServerError(S.C, err)
		return copter.CopterResponse{}, err
	}
	// Вычисляем параметры ESC
	escWarn := ControllerService{}.GetProperties(S.Props.ControllerProperties, battProps.BatteryVoltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := PropellerService{}.GetProperties(S.Props.PropellerProperties, S.Props.FrameProperties)

	// Вычисляем параметры мотора
	motorProps, motorWarn := MotorService{}.GetProperties(S.Props.MotorProperties, S.Props.FrameProperties)

	// Вычисление общих параметров для обоих режимов полета
	generalProps := GeneralService{}.GetProperties(S.Props.FrameProperties, S.Props.AttachmentsProperties, battProps.Mass, S.Props.ControllerProperties.Mass, propProps.Mass, motorProps.Mass)

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

	// Обороты при зависании, (Об/Мин)
	var RPM float64 = PropellerSpeed * 60

	// Угловая скорость пропеллера, (Рад/С)
	var PropellerAngularSpeed float64 = float64(2) * consts.Pi * PropellerSpeed

	// Обратная ЭДС электродвигателя при нависании, (В)
	var MotorBackEMF float64 = PropellerAngularSpeed * S.Calc.MotorProperties.PhaseValueOfEMFConst

	// Скорость воздушного потока, проходящего через пропеллер при нависании, (М/С)
	var AirFlowSpeed float64 = math.Sqrt(PropellerHangingLift / (float64(2) * S.Calc.EnvironmentProperties.AirDensity * S.Calc.PropellerProperties.SweptArea))

	// Мощность, необходимая для вращения пропеллера с заданной частотой, (Вт):
	var PowerForPropeller float64 = float64(S.Props.PropellerProperties.DimensionlessPowerConstant) * S.Calc.EnvironmentProperties.AirDensity * math.Pow(PropellerSpeed, 3) * math.Pow(float64(S.Props.PropellerProperties.Diameter), 5)

	// КПД в режиме висения
	var PropellerEfficiency float64 = S.Calc.PropellerProperties.AerodynamicQuality * float64(S.Props.PropellerProperties.Diameter) * math.Sqrt(S.Calc.EnvironmentProperties.AirDensity/PropellerHangingLift)

	warnings := WarningService{}.Append()

	return response_properties.HoverProperties{
		PropellerHangingLift:  PropellerHangingLift,
		PropellerSpeed:        PropellerSpeed,
		RPM:                   RPM,
		PropellerAngularSpeed: PropellerAngularSpeed,
		MotorBackEMF:          MotorBackEMF,
		AirFlowSpeed:          AirFlowSpeed,
		PowerForPropeller:     PowerForPropeller,
		PropellerEfficiency:   PropellerEfficiency,
	}, warnings
}
