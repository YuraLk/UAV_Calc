package properties_service

import (
	requests_properties "github.com/YuraLk/teca_server/internal/dtos/requests/properties"
	responses_properties "github.com/YuraLk/teca_server/internal/dtos/responses/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"

	"github.com/YuraLk/teca_server/internal/types"
)

func GetPropellerProperties(obj requests_properties.PropellerProperties) (responses_properties.PropellerProperties, *[]types.Warning) {

	warnings := warning_service.AppendWarnings()

	// Возвращаем расчитанные параметры
	properties := responses_properties.PropellerProperties{}

	return properties, warnings
}
