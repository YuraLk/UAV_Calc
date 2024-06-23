package response_properties

import (
	dtos "github.com/YuraLk/drone_calc/backend/internal/dtos/battery"
)

type BatteryProperties struct {
	CurrentOutput    CurrentOutput     `json:"currentOutput"`
	Capacity         float32           `json:"capacity"`
	UsableCapacity   float32           `json:"usableCapacity"`
	Mass             float32           `json:"mass"`
	CVCRange         []dtos.BatteryDto `json:"cvcRange"`
	Voltage          float64           `json:"voltage"`
	VoltageUnderLoad float64           `json:"voltageUnderLoad"`
	Power            float64           `json:"power"`
	UsablePower      float64           `json:"usablePower"`
}

type CurrentOutput struct {
	Per float32 `json:"per"` // Постоянное значение тока, (А)
	Max float32 `json:"max"` // Максимальное значениие тока, (А)
}
