package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/consts"
	"github.com/YuraLk/teca_server/internal/dtos"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter/request_properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter/response_properties"
)

type MotorService struct{}

func (MotorService) GetProperties(
	motor request_properties.MotorProperties,
	frame request_properties.FrameProperties,
	battery response_properties.BatteryProperties,
) (response_properties.MotorProperties, *[]dtos.WarningDto) {

	// Масса моторов, (Кг)
	var Mass float32 = motor.Mass * float32(frame.PropellersNumber)

	// Фазовое значение константы ЭДС
	var PhaseValueOfEMFConst float64 = motor.TorqueConst * motor.Currents.NoLoadConst

	// Максимальное значение электрической мощности двигателя на борту при текущем аккумуляторе, (Вт)
	var MaxPowerOfMotorOnBoard float64 = float64(motor.Currents.Max) * battery.VoltageUnderLoad

	// Радиус мотора в виде цилиндра, (М)
	var Radius float64 = float64(motor.Diameter) / 2

	// Площадь поверхности теплообмена, (М^2)
	var HeatTransferSurfaceArea float64 = 2*consts.Pi*Radius*float64(motor.Height) + consts.Pi*math.Pow(Radius, 2)

	// Объем ротора двигателя, (М^3)
	var Volume float64 = consts.Pi * math.Pow(Radius, 2) * float64(motor.Height)

	// Характеристическая длина двигателя в виде цилиндра, (М)
	var CharacteristicLength float64 = HeatTransferSurfaceArea / Volume // Площадь верхней и боковой поверхностей цилиндра

	warnings := WarningService{}.Append()
	return response_properties.MotorProperties{
		Mass:                   Mass,
		PhaseValueOfEMFConst:   PhaseValueOfEMFConst,
		MaxPowerOfMotorOnBoard: MaxPowerOfMotorOnBoard,
		CharacteristicLength:   CharacteristicLength,
	}, warnings
}
