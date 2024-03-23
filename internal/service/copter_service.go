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
	escWarn := ControllerService{}.GetProperties(S.Props.ControllerProperties, battProps.Voltage)

	// Вычисляем параметры пропеллера
	propProps, propWarn := PropellerService{}.GetProperties(S.Props.PropellerProperties, S.Props.FrameProperties)

	// Вычисляем параметры мотора
	motorProps, motorWarn := MotorService{}.GetProperties(S.Props.MotorProperties, S.Props.FrameProperties, battProps)

	// Вычисление общих параметров для обоих режимов полета
	generalProps := GeneralService{}.GetProperties(S.Props.FrameProperties, S.Props.AttachmentsProperties, battProps.Mass, S.Props.ControllerProperties.Mass, propProps.Mass, motorProps.Mass)

	// Вычисляем параметры для режима зависания
	hoverProps, hoverWarn := ModeProperties{Props: S.Props, Calc: copter.StandartProperties{
		EnvironmentProperties: envProps,
		BatteryProperties:     battProps,
		PropellerProperties:   propProps,
		MotorProperties:       motorProps,
		GeneralProperties:     generalProps,
	}}.GetHoverProperties()

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

// РЕЖИМ ЗАВИСАНИЯ
func (S ModeProperties) GetHoverProperties() (response_properties.HoverProperties, *[]dtos.WarningDto) {

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

	// Мощность, необходимая для вращения пропеллера с заданной частотой, (Вт):
	var PowerForPropeller float64 = float64(S.Props.PropellerProperties.DimensionlessPowerConstant) * S.Calc.EnvironmentProperties.AirDensity * math.Pow(PropellerSpeed, 3) * math.Pow(float64(S.Props.PropellerProperties.Diameter), 5)

	// Коэффицент совершенства пропеллера в режиме висения и условный КПД пропеллера в режиме висения
	var AerodynamicCleanness float64 = math.Sqrt(float64(2)/consts.Pi) * S.Calc.PropellerProperties.AerodynamicQuality

	// Скорость подсасывания воздуха, (М/С)
	var AirSuctionSpeed float64 = math.Sqrt((float64(2)*PropellerHangingLift)/(consts.Pi*S.Calc.EnvironmentProperties.AirDensity)) / float64(S.Props.PropellerProperties.Diameter)

	// Скорость отбрасывания воздуха, (М/С)
	var AirEjectionSpeed float64 = 2 * AirSuctionSpeed

	// Максимальная эффективность, (Н/Вт)
	var MaximumEfficiency float64 = 1 / AirSuctionSpeed

	// Реальная эффективность пропеллера, (Н/Вт)
	var RealEfficiency float64 = MaximumEfficiency * AerodynamicCleanness

	// Механическая мощность электродвигателя, необходимая для нависания, (Вт)
	var MotorMechanicalPower float64 = AerodynamicCleanness * PowerForPropeller

	// Электрическая мощность электродвигателя, необходимая для нависания, (Вт)
	var MotorElectricalPower float64 = float64(S.Props.MotorProperties.Efficiency) * MotorMechanicalPower

	// Газ линейный при висении, (%)
	var GasLinear float64 = MotorElectricalPower / S.Calc.MotorProperties.MaxPowerOfMotorOnBoard

	// Электрическая мощность силовой установки на нависании, (Вт)
	var ElectricalPowerOfPowerPlant float64 = MotorElectricalPower * float64(S.Props.FrameProperties.PropellersNumber)

	// Электрическая мощность, потребляемая всеми элементами цепи на нависании, (Вт)
	var TotalElectricalPower float64 = ElectricalPowerOfPowerPlant + float64(S.Props.AttachmentsProperties.PowerConsumption)

	// Средний ток потребления при мощности, необходимой для зависания, (А)
	var AverageCurrentConsumption float64 = BatteryService{}.GetAverageCurrent(TotalElectricalPower, S.Calc.BatteryProperties.CVCRange)

	// Средний ток потребления силовой установки при зависании, (А)
	var AverageCurrentOfPowerPlant float64 = BatteryService{}.GetAverageCurrent(ElectricalPowerOfPowerPlant, S.Calc.BatteryProperties.CVCRange)

	// Средний ток потребления на один двигатель при зависании, (А)
	var MotorAverageCurrent float64 = AverageCurrentOfPowerPlant / float64(S.Props.FrameProperties.PropellersNumber)

	// Напряжение на двигателе под нагрузкой при зависании, (В)
	var MotorVoltageUnderLoad float64 = S.Calc.BatteryProperties.VoltageUnderLoad - (AverageCurrentConsumption * S.Props.MotorProperties.WindingResistance * S.Calc.MotorProperties.PhaseValueOfEMFConst)

	// Время зависания, (Мин)
	var TimeOfFlight float64 = (float64(S.Calc.BatteryProperties.UsableCapacity) / AverageCurrentConsumption) * 60

	// Температура мотора при зависании, (°С)
	var MotorTemperature float64 = TemperatureService(S).GetMotorTemperature(MotorAverageCurrent)

	// Проверка маневренности ЛА
	maneuverabilityCheck := WarningService{}.ManeuverabilityCheck(GasLinear)

	// Проверка доступности энергии для висения
	motorPowerCheck := WarningService{}.MotorPowerCheck(S.Calc.MotorProperties.MaxPowerOfMotorOnBoard, MotorElectricalPower)

	warnings := WarningService{}.Append(maneuverabilityCheck, motorPowerCheck)

	return response_properties.HoverProperties{
		PropellerHangingLift:  PropellerHangingLift,
		PropellerSpeed:        PropellerSpeed,
		RPM:                   RPM,
		PropellerAngularSpeed: PropellerAngularSpeed,
		MotorBackEMF:          MotorBackEMF,
		PowerForPropeller:     PowerForPropeller,
		AerodynamicCleanness:  AerodynamicCleanness,
		AirSuctionSpeed:       AirSuctionSpeed,
		AirEjectionSpeed:      AirEjectionSpeed,
		MaximumEfficiency:     MaximumEfficiency,
		RealEfficiency:        RealEfficiency,
		MotorMechanicalPower:  MotorMechanicalPower,
		MotorElectricalPower:  MotorElectricalPower,
		GasLinear:             GasLinear,
		MotorAverageCurrent:   MotorAverageCurrent,
		MotorVoltageUnderLoad: MotorVoltageUnderLoad,
		TimeOfFlight:          TimeOfFlight,
		MotorTemperature:      MotorTemperature,
	}, warnings
}
