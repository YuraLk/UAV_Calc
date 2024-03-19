package service

import (
	"github.com/YuraLk/teca_server/internal/dtos"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
)

type MotorService struct{}

func (MotorService) GetProperties(motor request_properties.MotorProperties, frame request_properties.FrameProperties) (response_properties.MotorProperties, *[]dtos.WarningDto) {

	// Масса моторов, (Кг)
	var Mass float32 = motor.Mass * float32(frame.PropellersNumber)

	// Фазовое значение константы ЭДС
	var PhaseValueOfEMFConst float64 = motor.TorqueConst * motor.Currents.NoLoadConst

	warnings := WarningService{}.Append()
	return response_properties.MotorProperties{
		Mass:                 Mass,
		PhaseValueOfEMFConst: PhaseValueOfEMFConst,
	}, warnings
}
