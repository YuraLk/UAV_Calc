package properties_service

import (
	"github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"
	"github.com/YuraLk/teca_server/internal/dtos/responses/responses_properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetMotorProperties(motor requests_properties.MotorProperties, frame requests_properties.FrameProperties) (responses_properties.MotorProperties, *[]types.Warning) {

	// Масса моторов, (Кг)
	var Mass float32 = motor.Mass * float32(frame.PropellersNumber)

	// Фазовое значение константы ЭДС
	var PhaseValueOfEMFConst float64 = motor.TorqueConst * motor.Currents.NoLoadConst

	warnings := warning_service.AppendWarnings()
	return responses_properties.MotorProperties{
		Mass:                 Mass,
		PhaseValueOfEMFConst: PhaseValueOfEMFConst,
	}, warnings
}
