package service

import (
	requests "github.com/YuraLk/teca_server/internal/requests/properties"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetControllerProperties(obj requests.ControllerProperties, BatteryVoltage float64) *[]types.Warning {
	controllerVoltageWarning := ControllerVoltageCheck(float64(obj.Voltage), BatteryVoltage)

	warnings := AppendWarnings(controllerVoltageWarning)
	return warnings
}
