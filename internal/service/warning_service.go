package service

import (
	"github.com/YuraLk/teca_server/internal/dtos"
)

type WarningService struct{}

// Функция для объединения массивов предупреждений из сервисов
func (WarningService) AppendArrays(arrays ...*[]dtos.WarningDto) []dtos.WarningDto {
	var warnings []dtos.WarningDto
	for _, warning := range arrays {
		if warning != nil {
			warnings = append(warnings, *warning...)
		}
	}
	return warnings
}

// Функция для объединения предупреждений
func (WarningService) Append(array ...*dtos.WarningDto) *[]dtos.WarningDto {
	var warnings []dtos.WarningDto
	for _, warning := range array {
		if warning != nil {
			warnings = append(warnings, *warning)
		}
	}
	return &warnings
}

// Проверка приемлимой влажности воздуха для полета
func (WarningService) EnvironmentAirHumidityCheck(AirHumidity float64) *dtos.WarningDto {
	if AirHumidity >= 0.8 {
		return &dtos.WarningDto{
			Level: 3,
			Field: "Environment.AirHumidity",
			Text:  "Влажность воздуха превышает рекомендуемый предел. Из-за конденсации влаги на электронике могут возникнуть короткие замыкания или другие поломки.",
		}
	}

	return nil
}

func (WarningService) ControllerVoltageCheck(ControllerVoltage float64, BatteryVoltage float64) *dtos.WarningDto {
	if ControllerVoltage < BatteryVoltage {
		return &dtos.WarningDto{
			Level: 2,
			Field: "Controller.Voltage",
			Text:  "Номинальное напряжение аккумулятора превышает номинальное напряжение регулятора скорости (ESC). Рекомендуем снизить количество последовательно соединенных ячеек аккумулятора (S) или подобрать ESC с большим номинальным напряжением.",
		}
	}
	return nil
}

func (WarningService) ManeuverabilityCheck(GasLinear float64) *dtos.WarningDto {
	if GasLinear >= 0.8 {
		return &dtos.WarningDto{
			Level: 2,
			Field: "Motor.Power",
			Text:  "Для минимальной маневренности значение газа должно быть меньше 80%.",
		}
	}
	return nil
}

func (WarningService) MotorPowerCheck(AvailablePower float64, HoverPower float64) *dtos.WarningDto {
	if AvailablePower <= HoverPower {
		return &dtos.WarningDto{
			Level: 1,
			Field: "Motor.Power",
			Text:  "Доступной энергии недостаточно для зависания.",
		}
	}
	return nil
}
