package service

import (
	"encoding/json"

	"github.com/YuraLk/teca_server/internal/models"
	requests_properties "github.com/YuraLk/teca_server/internal/requests/properties"
	responses_properties "github.com/YuraLk/teca_server/internal/responses/properties"
	"github.com/YuraLk/teca_server/internal/types"
)

func getVoltageCharacteristics(CVC []types.BatteryData, CriticalChargeProportion float32, InitialStateOfCharge float32) ([]types.BatteryData, float64, float64) {
	// Используемый диапазон ВАХ исходя из заданного диапазона зарядов
	var CVCRange []types.BatteryData

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

func GetBatteryProperties(obj requests_properties.BatteryProperties, composit models.Composit) (responses_properties.BatteryProperties, error) {
	// Общая емкость, (А/Ч)
	var Capacity float32 = obj.CellCapacity * float32(obj.P) * float32(obj.S)

	// Критическая доля разряда
	var CriticalChargeProportion float32 = 1 - obj.MaxDischargePercent

	// Используемая емкость АКБ, (А/Ч)
	var UsableCapacity float32 = Capacity * (obj.InitialStateOfCharge - CriticalChargeProportion)

	// Токоотдача аккумулятора, (А)
	CurrentOutput := responses_properties.CurrentOutput{
		Per: float32(obj.CRating.Per) * float32(obj.P) * obj.CellCapacity,
		Max: float32(obj.CRating.Max) * float32(obj.P) * obj.CellCapacity,
	}

	// Масса АКБ, (Кг)
	var Mass float32 = float32(obj.S) * float32(obj.P) * obj.CellCapacity

	// Декодируем ВАХ аккумулятора из jsonb в []types.BatteryData
	var CVC []types.BatteryData
	err := json.Unmarshal([]byte(composit.CVC), &CVC)
	// В случае ошибки выбрасываем 500 статус
	if err != nil {
		return responses_properties.BatteryProperties{}, err
	}

	// По ВАХ аккумулятора ищем напряжения
	CVCRange, SmoothedCellVoltage, CellVoltageUnderLoad := getVoltageCharacteristics(CVC, CriticalChargeProportion, obj.InitialStateOfCharge)

	// Общее напряжение аккумулятора (В)
	var BatteryVoltage float64 = SmoothedCellVoltage * float64(obj.S)

	// Напряжение аккумулятора под нагрузкой, (В)
	var BatteryVoltageUnderLoad = CellVoltageUnderLoad * float64(obj.S)

	// Мощность аккумулятора исходя из ВАХ аккумулятора, (Вт * Час)
	var BatteryPower float64 = float64(Capacity) * BatteryVoltage

	// Используемая мощность аккумулятора, (Вт * Час)
	var BatteryUsablePower float64 = float64(UsableCapacity) * BatteryVoltage

	return responses_properties.BatteryProperties{
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
