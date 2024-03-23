package service

import (
	"github.com/YuraLk/teca_server/internal/dtos/copter"
)

type TemperatureService struct {
	Props copter.CalculateCopter
	Calc  copter.StandartProperties
}

// Расчет температуры мотора
func (S TemperatureService) GetMotorTemperature(Current float64) float64 {
	// Потери мощности при сопротивлении обмоток
	// var PowerLoss float64 = math.Pow(Current, 2) * S.Props.MotorProperties.WindingResistance

	// Теплопроводность воздуха

	return 0
}
