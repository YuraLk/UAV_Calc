package properties_service

import (
	requests_properties "github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"

	"github.com/YuraLk/teca_server/internal/types"
)

func GetControllerProperties(controller requests_properties.ControllerProperties, BatteryVoltage float64) *[]types.Warning {
	controllerVoltageWarning := warning_service.ControllerVoltageCheck(float64(controller.Voltage), BatteryVoltage)

	warnings := warning_service.AppendWarnings(controllerVoltageWarning)
	return warnings
}
