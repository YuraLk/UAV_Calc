package responses

import (
	responses_properties "github.com/YuraLk/teca_server/internal/responses/properties"
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
}
