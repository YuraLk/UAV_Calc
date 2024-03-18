package service

import "github.com/YuraLk/teca_server/internal/types"

func AppendWarningsArrays(arrays ...*[]types.Warning) []types.Warning {
	var warnings []types.Warning
	for _, warning := range arrays {
		if warning != nil {
			warnings = append(warnings, *warning...)
		}
	}
	return warnings
}

func AppendWarnings(array ...*types.Warning) *[]types.Warning {
	var warnings []types.Warning
	for _, warning := range array {
		if warning != nil {
			warnings = append(warnings, *warning)
		}
	}
	return &warnings
}

// Проверка приемлимой влажности воздуха для полета
func EnvironmentAirHumidityCheck(AirHumidity float64) *types.Warning {
	if AirHumidity >= 0.8 {
		return &types.Warning{
			Level: 3,
			Field: "Environment.AirHumidity",
			Text:  "Влажность воздуха превышает рекомендуемый предел. Из-за конденсации влаги на электронике могут возникнуть короткие замыкания или другие поломки.",
		}
	}

	return nil
}

func ControllerVoltageCheck(ControllerVoltage float64, BatteryVoltage float64) *types.Warning {
	if ControllerVoltage < BatteryVoltage {
		return &types.Warning{
			Level: 2,
			Field: "Controller.Voltage",
			Text:  "Номинальное напряжение аккумулятора превышает номинальное напряжение регулятора скорости (ESC). Рекомендуем снизить количество последовательно соединенных ячеек аккумулятора (S) или подобрать ESC с большим номинальным напряжением.",
		}
	}
	return nil
}
