package hover_service

import (
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetHoverProperties() (response_properties.HoverProperties, *[]types.Warning) {

	warnings := warning_service.AppendWarnings()

	return response_properties.HoverProperties{}, warnings
}
