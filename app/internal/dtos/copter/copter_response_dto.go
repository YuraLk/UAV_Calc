package copter

import (
	"github.com/YuraLk/drone_calc/backend/internal/dtos"
	properties "github.com/YuraLk/drone_calc/backend/internal/dtos/copter/response_properties"
)

type CopterResponse struct {
	CopterProperties CopterProperties `json:"properties"`
	// Массив предупреждений
	Warings []dtos.WarningDto `json:"warnings"`
}

type CopterProperties struct {
	EnvironmentProperties properties.EnvironmentProperties `json:"environmentProperties"`
	BatteryProperties     properties.BatteryProperties     `json:"batteryProperties"`
	PropellerProperties   properties.PropellerProperties   `json:"propellerProperties"`
	MotorProperties       properties.MotorProperties       `json:"motorProperties"`
	GeneralProperties     properties.GeneralProperties     `json:"generalProperties"`
	HoverProperties       properties.HoverProperties       `json:"hoverProperties"`
	MaxProperties         properties.MaxProperties         `json:"maxProperties"`
}
