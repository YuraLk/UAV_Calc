package properties_service

import (
	"github.com/YuraLk/teca_server/internal/consts"
	"github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"
	"github.com/YuraLk/teca_server/internal/dtos/responses/responses_properties"
)

func sum(array []float32) float32 {
	var sum float32
	for _, v := range array {
		sum += v
	}
	return sum
}

func GetGeneralProperties(frame requests_properties.FrameProperties, attachments requests_properties.AttachmentsProperties, array ...float32) responses_properties.GeneralProperties {

	// Общая масса сборки, (Кг)
	Mass := sum(array) + frame.Mass + attachments.Mass

	// Масса ВМГ, (Кг)
	PMGMass := sum(array)

	// Общий вес сборки, (Н)
	Weight := Mass * consts.G

	return responses_properties.GeneralProperties{
		Mass:    Mass,
		Weight:  Weight,
		PMGMass: PMGMass,
	}
}
