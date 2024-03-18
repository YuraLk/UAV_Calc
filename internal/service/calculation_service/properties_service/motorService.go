package properties_service

import (
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"

	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetMotorProperties(motor request_properties.MotorProperties, frame request_properties.FrameProperties) (response_properties.MotorProperties, *[]types.Warning) {

	// Масса моторов, (Кг)
	var Mass float32 = motor.Mass * float32(frame.PropellersNumber)

	// Фазовое значение константы ЭДС
	var PhaseValueOfEMFConst float64 = motor.TorqueConst * motor.Currents.NoLoadConst

	warnings := warning_service.AppendWarnings()
	return response_properties.MotorProperties{
		Mass:                 Mass,
		PhaseValueOfEMFConst: PhaseValueOfEMFConst,
	}, warnings
}
