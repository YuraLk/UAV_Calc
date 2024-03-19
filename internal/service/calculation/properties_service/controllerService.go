package properties_service

import (
	"github.com/YuraLk/teca_server/internal/dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation/warning_service"
)

func GetControllerProperties(controller properties.ControllerProperties, BatteryVoltage float64) *[]dtos.WarningDto {
	controllerVoltageWarning := warning_service.ControllerVoltageCheck(float64(controller.Voltage), BatteryVoltage)

	warnings := warning_service.AppendWarnings(controllerVoltageWarning)
	return warnings
}
