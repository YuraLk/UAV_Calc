package service

import (
	"github.com/YuraLk/teca_server/internal/consts"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/utils"
)

type GeneralService struct{}

func (GeneralService) GetProperties(frame request_properties.FrameProperties, attachments request_properties.AttachmentsProperties, array ...float32) response_properties.GeneralProperties {

	// Общая масса сборки, (Кг)
	Mass := utils.SumElements(array) + frame.Mass + attachments.Mass

	// Масса ВМГ, (Кг)
	PMGMass := utils.SumElements(array)

	// Общий вес сборки, (Н)
	Weight := Mass * consts.G

	return response_properties.GeneralProperties{
		Mass:    Mass,
		Weight:  Weight,
		PMGMass: PMGMass,
	}
}
