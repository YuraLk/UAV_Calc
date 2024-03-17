package service

import "github.com/YuraLk/teca_server/internal/types"

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
