package service

import (
	requests_properties "github.com/YuraLk/teca_server/internal/requests/properties"
	responses_properties "github.com/YuraLk/teca_server/internal/responses/properties"

	"github.com/YuraLk/teca_server/internal/types"
)

func GetPropellerProperties(obj requests_properties.PropellerProperties) (responses_properties.PropellerProperties, *[]types.Warning) {

	warnings := AppendWarnings()

	// Возвращаем расчитанные параметры
	properties := responses_properties.PropellerProperties{}

	return properties, warnings
}
