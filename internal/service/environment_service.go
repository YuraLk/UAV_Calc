package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/consts"
	"github.com/YuraLk/teca_server/internal/dtos"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter/request_properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter/response_properties"
)

type EnvironmentService struct{}

func (EnvironmentService) GetVerticalTemperatureGradient(AirHumidity float64, AirTemperature float64) float64 {
	var DryAdiabaticGradient float64 = consts.G / consts.SpecificHeatCapacityAtConstantPressure
	switch AirHumidity {
	case 0:
		return DryAdiabaticGradient
	default:
		var MoistAdiabaticGradient float64 = DryAdiabaticGradient * ((1 + ((consts.SpecificHeatWaterEvaporation * AirHumidity) / (consts.SpecificGasConstForDryAir * AirTemperature))) / (1 + ((math.Pow(consts.SpecificHeatWaterEvaporation, 2) * AirHumidity) / (consts.SpecificHeatCapacityAtConstantPressure * consts.SpecificGasConstForWaterVapor * math.Pow(AirTemperature, 2)))))
		return MoistAdiabaticGradient
	}
}

func (EnvironmentService) GetAirDensity(AirHumidity float64, AirTemperature float64, Pressure float64, PartialPressureOfWaterVapor float64) float64 {
	switch AirHumidity {
	case 0: // Если воздух сухой
		DryAirDensity := Pressure / (consts.SpecificGasConstForDryAir * AirTemperature)
		return DryAirDensity
	default: // Если воздух влажный
		DryAirPressure := Pressure - PartialPressureOfWaterVapor // Давление сухого воздуха
		DensityOfHumidAir := (DryAirPressure/(consts.SpecificGasConstForDryAir*AirTemperature) + (PartialPressureOfWaterVapor / (consts.SpecificGasConstForWaterVapor * AirTemperature)))
		return DensityOfHumidAir
	}
}

func (EnvironmentService) GetProperties(environment request_properties.EnvironmentProperties) (response_properties.EnvironmentProperties, *[]dtos.WarningDto) {

	// Проверка на допустимую влажность воздуха
	airHumidityWarning := WarningService{}.EnvironmentAirHumidityCheck(environment.AirHumidity)

	// Высота взлета летательного аппарата относительно оператора, (М)
	var LocalAltitude float64 = environment.AltitudeRange.Flight - environment.AltitudeRange.Start

	// Давление на высоте полета летательного аппарата, (Па)
	var Pressure float64 = consts.SeaLevelPressure * math.Pow((float64(1-((consts.TemperatureLaps*environment.AltitudeRange.Flight)/consts.SeaLevelStandardTemperature))), ((consts.G*consts.M)/consts.UniversalGasConstant*consts.AtmosphericTemperatureGradient))

	// Вертикальный температурный градиент, (K/км)
	var VerticalTemperatureGradient float64 = EnvironmentService{}.GetVerticalTemperatureGradient(environment.AirHumidity, environment.AirTemperature)

	// Температура воздуха на высоте полета ЛА, (K)
	var FlightAirTemperature float64 = environment.AirTemperature - VerticalTemperatureGradient*float64(LocalAltitude/1000)

	// Давление насыщения при определенной температуре, (Па)
	var SaturationPressureAtCertainTemperature float64 = consts.ApproximateSaturationPressureOfWaterVaporAtSurfaceOfWaterAt0 * math.Exp(consts.EmpiricalTetensCoefficientA*FlightAirTemperature/(consts.EmpiricalTetensCoefficientB+FlightAirTemperature))

	// Парциальное давление водяного пара, (Па)
	var PartialPressureOfWaterVapor float64 = environment.AirHumidity * SaturationPressureAtCertainTemperature

	// Плотность воздуха, (Кг/М^3)
	var AirDensity float64 = EnvironmentService{}.GetAirDensity(environment.AirHumidity, environment.AirTemperature, Pressure, PartialPressureOfWaterVapor)

	// Возвращаем расчитанные параметры
	properties := response_properties.EnvironmentProperties{
		AirDensity: AirDensity,
	}

	warnings := WarningService{}.Append(airHumidityWarning)

	return properties, warnings
}
