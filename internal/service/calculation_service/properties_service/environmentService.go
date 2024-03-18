package properties_service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/consts"
	requests_properties "github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"
	responses_properties "github.com/YuraLk/teca_server/internal/dtos/responses/responses_properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
	"github.com/YuraLk/teca_server/internal/types"
)

func getVerticalTemperatureGradient(AirHumidity float64, AirTemperature float64) float64 {
	var DryAdiabaticGradient float64 = consts.G / consts.SpecificHeatCapacityAtConstantPressure
	switch AirHumidity {
	case 0:
		return DryAdiabaticGradient
	default:
		var MoistAdiabaticGradient float64 = DryAdiabaticGradient * ((1 + ((consts.SpecificHeatWaterEvaporation * AirHumidity) / (consts.SpecificGasConstForDryAir * AirTemperature))) / (1 + ((math.Pow(consts.SpecificHeatWaterEvaporation, 2) * AirHumidity) / (consts.SpecificHeatCapacityAtConstantPressure * consts.SpecificGasConstForWaterVapor * math.Pow(AirTemperature, 2)))))
		return MoistAdiabaticGradient
	}
}

func getAirDensity(AirHumidity float64, AirTemperature float64, Pressure float64, PartialPressureOfWaterVapor float64) float64 {
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

func GetEnvironmentProperties(obj requests_properties.EnvironmentProperties) (responses_properties.EnvironmentProperties, *[]types.Warning) {

	// Проверка на допустимую влажность воздуха
	airHumidityWarning := warning_service.EnvironmentAirHumidityCheck(obj.AirHumidity)

	// Высота взлета летательного аппарата относительно оператора, (М)
	var LocalAltitude float64 = obj.AltitudeRange.Flight - obj.AltitudeRange.Start

	// Давление на высоте полета летательного аппарата, (Па)
	var Pressure float64 = consts.SeaLevelPressure * math.Pow((float64(1-((consts.TemperatureLaps*obj.AltitudeRange.Flight)/consts.SeaLevelStandardTemperature))), ((consts.G*consts.M)/consts.UniversalGasConstant*consts.AtmosphericTemperatureGradient))

	// Вертикальный температурный градиент, (K/км)
	var VerticalTemperatureGradient float64 = getVerticalTemperatureGradient(obj.AirHumidity, obj.AirTemperature)

	// Температура воздуха на высоте полета ЛА, (K)
	var FlightAirTemperature float64 = obj.AirTemperature - VerticalTemperatureGradient*float64(LocalAltitude/1000)

	// Давление насыщения при определенной температуре, (Па)
	var SaturationPressureAtCertainTemperature float64 = consts.ApproximateSaturationPressureOfWaterVaporAtSurfaceOfWaterAt0 * math.Exp(consts.EmpiricalTetensCoefficientA*FlightAirTemperature/(consts.EmpiricalTetensCoefficientB+FlightAirTemperature))

	// Парциальное давление водяного пара, (Па)
	var PartialPressureOfWaterVapor float64 = obj.AirHumidity * SaturationPressureAtCertainTemperature

	// Плотность воздуха, (Кг/М^3)
	var AirDensity float64 = getAirDensity(obj.AirHumidity, obj.AirTemperature, Pressure, PartialPressureOfWaterVapor)

	// Возвращаем расчитанные параметры
	properties := responses_properties.EnvironmentProperties{
		AirDensity: AirDensity,
	}

	warnings := warning_service.AppendWarnings(airHumidityWarning)

	return properties, warnings
}
