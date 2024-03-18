package properties_service

import (
	"github.com/YuraLk/teca_server/internal/consts"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
)

func sum(array []float32) float32 {
	var sum float32
	for _, v := range array {
		sum += v
	}
	return sum
}

func GetGeneralProperties(frame request_properties.FrameProperties, attachments request_properties.AttachmentsProperties, array ...float32) response_properties.GeneralProperties {

	// Общая масса сборки, (Кг)
	Mass := sum(array) + frame.Mass + attachments.Mass

	// Масса ВМГ, (Кг)
	PMGMass := sum(array)

	// Общий вес сборки, (Н)
	Weight := Mass * consts.G

	return response_properties.GeneralProperties{
		Mass:    Mass,
		Weight:  Weight,
		PMGMass: PMGMass,
	}
}
