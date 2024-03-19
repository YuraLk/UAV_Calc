package service

import (
	"github.com/YuraLk/teca_server/internal/dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter/request_properties"
)

type ControllerService struct{}

func (ControllerService) GetProperties(controller request_properties.ControllerProperties, BatteryVoltage float64) *[]dtos.WarningDto {
	controllerVoltageWarning := WarningService{}.ControllerVoltageCheck(float64(controller.Voltage), BatteryVoltage)

	warnings := WarningService{}.Append(controllerVoltageWarning)
	return warnings
}
