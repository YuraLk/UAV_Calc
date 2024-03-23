package service

import (
	"encoding/json"

	dtos "github.com/YuraLk/teca_server/internal/dtos/battery"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter/request_properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter/response_properties"
	"github.com/YuraLk/teca_server/internal/models"
)

type BatteryService struct{}

func (BatteryService) GetVoltageCharacteristics(CVC []dtos.BatteryDto, CriticalChargeProportion float32, InitialStateOfCharge float32) ([]dtos.BatteryDto, float64, float64) {
	// Используемый диапазон ВАХ исходя из заданного диапазона зарядов
	var CVCRange []dtos.BatteryDto

	// Наибольшее сглаженное напряжение исходя из заданного диапазона зарядов
	var SmoothedVoltage float64

	// Наибольшее напряжение под нагрузкой исходя из заданного диапазона зарядов
	var VoltageUnderLoad float64

	for _, value := range CVC {
		var ChargePercentageProportion float32 = float32(value.ChargePercentage) / 100
		// Отбираем подходящий массив значений ВАХ
		if ChargePercentageProportion > CriticalChargeProportion && ChargePercentageProportion <= InitialStateOfCharge {
			CVCRange = append(CVCRange, value)
		}

		// Устанавливаем наибольшее сглаженное напряжение и наибольшее напряжение под нагрузкой
		if InitialStateOfCharge == ChargePercentageProportion {
			SmoothedVoltage = value.SmoothedVoltage
			VoltageUnderLoad = value.LoadVoltage
		}
	}

	return CVCRange, SmoothedVoltage, VoltageUnderLoad
}

// Расчет средней силы тока из участка ВАХ аккумулятора
func (BatteryService) GetAverageCurrent(Power float64, CVC []dtos.BatteryDto) float64 {
	var Current float64

	// Расчитываем силу тока для каждого процента заряда и суммируем
	for _, value := range CVC {
		Current += (Power / value.LoadVoltage)
	}

	return Current / float64(len(CVC))

}

func (BatteryService) GetProperties(battery request_properties.BatteryProperties, composit models.Composit) (response_properties.BatteryProperties, error) {
	// Общая емкость, (А/Ч)
	var Capacity float32 = battery.CellCapacity * float32(battery.P) * float32(battery.S)

	// Критическая доля разряда
	var CriticalChargeProportion float32 = 1 - battery.MaxDischargePercent

	// Используемая емкость АКБ, (А/Ч)
	var UsableCapacity float32 = Capacity * (battery.InitialStateOfCharge - CriticalChargeProportion)

	// Токоотдача аккумулятора, (А)
	CurrentOutput := response_properties.CurrentOutput{
		Per: float32(battery.CRating.Per) * float32(battery.P) * battery.CellCapacity,
		Max: float32(battery.CRating.Max) * float32(battery.P) * battery.CellCapacity,
	}

	// Масса АКБ, (Кг)
	var Mass float32 = float32(battery.S) * float32(battery.P) * battery.CellMass

	// Декодируем ВАХ аккумулятора из jsonb в []dtos.BatteryData
	var CVC []dtos.BatteryDto
	err := json.Unmarshal([]byte(composit.CVC), &CVC)
	// В случае ошибки выбрасываем 500 статус
	if err != nil {
		return response_properties.BatteryProperties{}, err
	}

	// По ВАХ аккумулятора ищем напряжения
	CVCRange, SmoothedCellVoltage, CellVoltageUnderLoad := BatteryService{}.GetVoltageCharacteristics(CVC, CriticalChargeProportion, battery.InitialStateOfCharge)

	// Общее напряжение аккумулятора (В)
	var Voltage float64 = SmoothedCellVoltage * float64(battery.S)

	// Напряжение аккумулятора под нагрузкой, (В)
	var VoltageUnderLoad = CellVoltageUnderLoad * float64(battery.S)

	// Мощность аккумулятора исходя из ВАХ аккумулятора, (Вт * Час)
	var Power float64 = float64(Capacity) * Voltage

	// Используемая мощность аккумулятора, (Вт * Час)
	var UsablePower float64 = float64(UsableCapacity) * Voltage

	return response_properties.BatteryProperties{
		CurrentOutput:    CurrentOutput,
		Capacity:         Capacity,
		UsableCapacity:   UsableCapacity,
		Mass:             Mass,
		CVCRange:         CVCRange,
		Voltage:          Voltage,
		VoltageUnderLoad: VoltageUnderLoad,
		Power:            Power,
		UsablePower:      UsablePower,
	}, nil
}
