package responses

import (
	responses_properties "github.com/YuraLk/teca_server/internal/dtos/responses/responses_properties"
	"github.com/YuraLk/teca_server/internal/types"
)

type CopterResponse struct {
	CopterProperties CopterProperties `json:"properties"`
	// Массив предупреждений
	Warings []types.Warning `json:"warnings"`
}

type CopterProperties struct {
	EnvironmentProperties responses_properties.EnvironmentProperties `json:"environmentProperties"`
	BatteryProperties     responses_properties.BatteryProperties     `json:"batteryProperties"`
	PropellerProperties   responses_properties.PropellerProperties   `json:"propellerProperties"`
	MotorProperties       responses_properties.MotorProperties       `json:"motorProperties"`
	GeneralProperties     responses_properties.GeneralProperties     `json:"generalProperties"`
}
