package responses

import (
	responseProperties "github.com/YuraLk/teca_server/internal/responses/properties"
	"github.com/YuraLk/teca_server/internal/types"
)

type CopterResponse struct {
	CopterProperties CopterProperties `json:"properties"`
	// Массив предупреждений
	Warings []types.Warning `json:"warnings"`
}

type CopterProperties struct {
	EnvironmentProperties responseProperties.EnvironmentProperties `json:"environmentProperties"`
}
