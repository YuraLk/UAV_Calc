package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/gin-gonic/gin"
)

type AccumulatorProperties struct {
	AccTotalMass uint64
	AccTotalVol  uint
	AccMaxOut    uint
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

func GetContMass(c *gin.Context, ContWeight uint, RamaVents uint8, LayoutID uint) (uint, error) {
	var layout models.Layout

	if err := postgres.DB.Where("id = ?", LayoutID).First(&layout).Error; err != nil {
		exeptions.NotFound(c, "Компоновка с данным ID не найдена!")
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

func GetAccFeatures(c *gin.Context, AccVol uint, AccCRating types.Current, AccMass uint64, AccBanks uint8, AccCount uint8, CompositID uint) (AccumulatorProperties, error) {
	// Ищем химический тип аккумулятора в БД
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", CompositID).First(&composit).Error; err != nil {
		exeptions.NotFound(c, "Характеристики аккумулятора с данным ID не найдены!")
		return AccumulatorProperties{}, err
	}

	// Масса аккумулятора
	AccTotalMass := AccMass * uint64(AccBanks) * uint64(AccCount)
	// Номинальное напряжение
	// AccNomVoltage := composit.Voltage.nom * float32(AccBanks)
	// Минимальное напряжение

	// Максимальное напряжение

	// Общая емкость аккумулятора в Ач
	AccTotalVol := AccVol * uint(AccCount)
	// Максимальная токоотдача
	AccMaxOut := AccTotalVol * uint(AccCRating.Max)
	// Доступная емкость аккумулятора

	return AccumulatorProperties{
		AccTotalMass: AccTotalMass,
		AccTotalVol:  AccTotalVol,
		AccMaxOut:    AccMaxOut,
	}, nil
}
