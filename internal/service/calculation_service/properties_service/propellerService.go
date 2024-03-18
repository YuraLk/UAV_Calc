package properties_service

import (
	requests_properties "github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"
	responses_properties "github.com/YuraLk/teca_server/internal/dtos/responses/responses_properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"

	"github.com/YuraLk/teca_server/internal/types"
)

func GetPropellerProperties(propeller requests_properties.PropellerProperties, frame requests_properties.FrameProperties) (responses_properties.PropellerProperties, *[]types.Warning) {

	// Радиус пропеллера, (М)
	// var PropellerRadius float32 = propeller.Diameter / 2

	// Относительный шаг винта, (М)
	// var PropellerRelativePitch float32 = propeller.Pitch / propeller.Diameter

	// Масса лопасти,(Кг)

	// Собираем все предупреждения
	warnings := warning_service.AppendWarnings()

	// Возвращаем расчитанные параметры
	properties := responses_properties.PropellerProperties{}

	return properties, warnings
}
