package properties_service

import (
	"encoding/json"

	dtos "github.com/YuraLk/teca_server/internal/dtos/battery_dtos"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/models"
)

func getVoltageCharacteristics(CVC []dtos.BatteryDto, CriticalChargeProportion float32, InitialStateOfCharge float32) ([]dtos.BatteryDto, float64, float64) {
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

func GetBatteryProperties(battery request_properties.BatteryProperties, composit models.Composit) (response_properties.BatteryProperties, error) {
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
	CVCRange, SmoothedCellVoltage, CellVoltageUnderLoad := getVoltageCharacteristics(CVC, CriticalChargeProportion, battery.InitialStateOfCharge)

	// Общее напряжение аккумулятора (В)
	var BatteryVoltage float64 = SmoothedCellVoltage * float64(battery.S)

	// Напряжение аккумулятора под нагрузкой, (В)
	var BatteryVoltageUnderLoad = CellVoltageUnderLoad * float64(battery.S)

	// Мощность аккумулятора исходя из ВАХ аккумулятора, (Вт * Час)
	var BatteryPower float64 = float64(Capacity) * BatteryVoltage

	// Используемая мощность аккумулятора, (Вт * Час)
	var BatteryUsablePower float64 = float64(UsableCapacity) * BatteryVoltage

	return response_properties.BatteryProperties{
		CurrentOutput:           CurrentOutput,
		Capacity:                Capacity,
		UsableCapacity:          UsableCapacity,
		Mass:                    Mass,
		CVCRange:                CVCRange,
		BatteryVoltage:          BatteryVoltage,
		BatteryVoltageUnderLoad: BatteryVoltageUnderLoad,
		BatteryPower:            BatteryPower,
		BatteryUsablePower:      BatteryUsablePower,
	}, nil
}
