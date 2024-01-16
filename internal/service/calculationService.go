package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/models"
)

type AccumulatorProperties struct {
	AccTotalMass uint64
	AccTotalVol  uint
	AccMaxOut    float32
}

func GetAirDensity(EnvTemp float32, EnvPress float32) float32 {
	// Переводим температуру в Кельвины
	EnvTemp += 273.15
	const (
		M = 28.97            // Г/Моль - молярная масса для сухого воздуха
		R = 8.31446261815324 // Дж/(Моль*К) - универсальная газовая постоянная
	)

	EnvAirPressure := EnvPress * M / (R * EnvTemp)

	return EnvAirPressure

}

func GetContMass(ContWeight uint, RamaVents uint8, LayoutID uint) (uint, error) {
	var layout models.Layout

	if err := postgres.DB.Where("id = ?", LayoutID).First(&layout).Error; err != nil {
		return 0, err
	}
	// Кол-во витов делим на множитель компоновки, откуда понимаем, сколько ESC данной компоновки нам требуется
	ContCount := math.Ceil(float64(RamaVents) / float64(layout.Multipler))

	// Высчитываем вес ESC
	ContTolalMass := ContWeight * uint(ContCount)

	return ContTolalMass, nil
}

func GetTotalMass(masses ...uint64) uint64 {
	var totalMass uint64 = 0
	for _, num := range masses {
		totalMass += num
	}
	return totalMass
}

func GetAccFeatures(AccVol uint, AccOut float32, AccMass uint64, AccBanks uint8, AccCount uint8) AccumulatorProperties {
	// Масса аккумулятора
	AccTotalMass := AccMass * uint64(AccBanks) * uint64(AccCount)

	// Номинальное напряжение
	// AccTotalVoltage := AccVoltage * float32(AccBanks)
	// Минимальное напряжение

	// Максимальное напряжение

	// Общая емкость аккумулятора в Ач
	AccTotalVol := AccVol * uint(AccCount)
	// Максимальная токоотдача
	AccMaxOut := float32(AccTotalVol) * AccOut

	// Доступная емкость аккумулятора

	return AccumulatorProperties{
		AccTotalMass: AccTotalMass,
		AccTotalVol:  AccTotalVol,
		AccMaxOut:    AccMaxOut,
	}
}
