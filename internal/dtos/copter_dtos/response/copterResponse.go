package response

import (
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/types"
)

type CopterResponse struct {
	CopterProperties CopterProperties `json:"properties"`
	// Массив предупреждений
	Warings []types.Warning `json:"warnings"`
}

type CopterProperties struct {
	EnvironmentProperties properties.EnvironmentProperties `json:"environmentProperties"`
	BatteryProperties     properties.BatteryProperties     `json:"batteryProperties"`
	PropellerProperties   properties.PropellerProperties   `json:"propellerProperties"`
	MotorProperties       properties.MotorProperties       `json:"motorProperties"`
	GeneralProperties     properties.GeneralProperties     `json:"generalProperties"`
	HoverProperties       properties.HoverProperties       `json:"hoverProperties"`
}
