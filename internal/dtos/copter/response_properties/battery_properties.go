package response_properties

import (
	dtos "github.com/YuraLk/teca_server/internal/dtos/battery"
)

type BatteryProperties struct {
	CurrentOutput           CurrentOutput     `json:"currentOutput"`
	Capacity                float32           `json:"capacity"`
	UsableCapacity          float32           `json:"usableCapacity"`
	Mass                    float32           `json:"mass"`
	CVCRange                []dtos.BatteryDto `json:"cvcRange"`
	BatteryVoltage          float64           `json:"batteryVoltage"`
	BatteryVoltageUnderLoad float64           `json:"batteryVoltageUnderLoad"`
	BatteryPower            float64           `json:"batteryPower"`
	BatteryUsablePower      float64           `json:"batteryUsablePower"`
}

type CurrentOutput struct {
	Per float32 `json:"per"` // Постоянное значение тока, (А)
	Max float32 `json:"max"` // Максимальное значениие тока, (А)
}
