package service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/gin-gonic/gin"
)

type BatteryProperties struct {
	BattMass              uint64
	BattCapacity          float32
	BattCurrPer           float32
	BattCurrMax           float32
	BattAvailableCapacity float32
	BattMinVoltage        float32
	BattNomVoltage        float32
	BattMaxVoltage        float32
	BattEnergy            float32 // Доступная энергия заряженного аккумулятора в Дж, исходя из доступной безопасной энергии
	BattSpecificEnergyVol float32 // Удельная энергоемкость
	BattEnergyReserve     float32 // Запас энергии батареи в Дж
}

type MotorProperties struct {
	MotElectricPower   float32
	MotMechanicalPower float32
}

type FlightProperties struct {
	Minimal Mode
}

type Mode struct {
	PropFreq           float32
	PropPower          float32
	UsefulPowerOfPlant float32
	Time               float32
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

func GetContMass(c *gin.Context, ContMass uint, AxisNumber uint8, LayoutID uint) (uint, error) {
	var layout models.Layout

	if err := postgres.DB.Where("id = ?", LayoutID).First(&layout).Error; err != nil {
		exeptions.NotFound(c, "Компоновка с данным ID не найдена!")
		return 0, err
	}
	// Кол-во витов делим на множитель компоновки, откуда понимаем, сколько ESC данной компоновки нам требуется
	ContCount := math.Ceil(float64(AxisNumber) / float64(layout.Multipler))

	// Высчитываем вес ESC
	ContTolalMass := ContMass * uint(ContCount)

	return ContTolalMass, nil
}

func GetAssemblyMass(masses ...uint64) uint64 {
	var totalMass uint64 = 0
	for _, num := range masses {
		totalMass += num
	}
	return totalMass
}

func GetBattFeatures(c *gin.Context, CellCapacity float32, CellCRating types.Current, CellMass uint64, BattTypeS uint8, BattTypeP uint8, CompositID uint) (BatteryProperties, error) {
	// Ищем химический тип аккумулятора в БД
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", CompositID).First(&composit).Error; err != nil {
		exeptions.NotFound(c, "Характеристики аккумулятора с данным ID не найдены!")
		return BatteryProperties{}, err
	}

	// Масса аккумулятора
	BattMass := CellMass * uint64(BattTypeS) * uint64(BattTypeP)
	// Номинальное напряжение сборочных элементов аккумулятора
	nomVoltage := composit.Voltage["nom"]
	nomVoltageFloat64 := nomVoltage.(float64)
	BattNomVoltage := nomVoltageFloat64 * float64(BattTypeS)
	// Минимальное напряжение сборочных элементов аккумулятора
	minVoltage := composit.Voltage["min"]
	minVoltageFloat64 := minVoltage.(float64)
	BattMinVoltage := minVoltageFloat64 * float64(BattTypeS)
	// Максимальное напряжение сборочных элементов аккумулятора
	maxVoltage := composit.Voltage["max"]
	maxVoltageFloat64 := maxVoltage.(float64)
	BattMaxVoltage := maxVoltageFloat64 * float64(BattTypeS)

	// Общая емкость аккумулятора в Ач
	BattCapacity := CellCapacity * float32(BattTypeP)
	// Максимальная токоотдача аккумулятора
	BattCurrMax := BattCapacity * float32(CellCRating.Max)
	// Постоянная токоотдача аккумулятора
	BattCurrPer := BattCapacity * float32(CellCRating.Per)
	// Доступная безопасная емкость аккумулятора в A/Ч
	BattAvailableCapacity := BattCapacity * composit.SafeCapacity
	// Номинальная энергия аккумулятора в Дж
	BattEnergy := BattCapacity * float32(BattNomVoltage) * 3600
	// Удельная энергоемкость аккумулятора в Дж/кг
	BattSpecificEnergyVol := BattEnergy / (float32(BattMass) / 1000)
	// Запас энергии батареи в Дж
	BattEnergyReserve := BattSpecificEnergyVol * (float32(BattMass) / 1000)

	return BatteryProperties{
		BattMass:              BattMass,
		BattCapacity:          BattCapacity,
		BattCurrPer:           BattCurrPer,
		BattCurrMax:           BattCurrMax,
		BattAvailableCapacity: BattAvailableCapacity,
		BattMinVoltage:        float32(BattMinVoltage),
		BattNomVoltage:        float32(BattNomVoltage),
		BattMaxVoltage:        float32(BattMaxVoltage),
		BattEnergy:            BattEnergy,
		BattSpecificEnergyVol: BattSpecificEnergyVol,
		BattEnergyReserve:     BattEnergyReserve,
	}, nil
}

func GetMotorFeatures(MotPeakCurrent float32, AccMaxVoltage float32) MotorProperties {
	const N float32 = 0.93 // Коэффицент полезного действия двигателя

	// Электрическая мощность двигателя в Вт
	MotElectricPower := MotPeakCurrent * AccMaxVoltage
	// Механическа мощность двигателя в Вт
	MotMechanicalPower := MotElectricPower * N
	return MotorProperties{
		MotElectricPower:   MotElectricPower,
		MotMechanicalPower: MotMechanicalPower,
	}
}

func GetAssemblyWeight(AssemblyMass uint64, BattMass uint64) float32 {
	const G float32 = 9.8
	// Получаем массу сборки без массы батареи в граммах
	UAVMass := AssemblyMass - BattMass
	// Получаем массу сборки в кг
	UAVMassSI := float32(UAVMass) / 1000
	// Получаем вес сборки в Ньютонах
	UAVWeight := UAVMassSI * G
	return UAVWeight
}

// Первый аргумент - вес сборки в Ньютонах
func GetFlightFeatures(AssemblyWeight float32, PropPowerConst float32, PropTractionConst float32, EnvAirPressure float32, PropDiameter float32, BattEnergyReserve float32, MotorMaxPower uint) FlightProperties {
	const n float32 = 0.9 // Коэффицент полезного действия всей силовой установки

	// МИНИМАЛЬНЫЕ ХАРАКТЕРИСТИКИ

	// Вычисление минимальной частоты вращения винта, необходимого для поддержания БПЛА в воздухе
	PropMinFreq := math.Sqrt(float64(AssemblyWeight) / (float64(PropPowerConst) * float64(EnvAirPressure/1000) * math.Pow(float64(PropDiameter/1000), 4)))
	// Количество оборотов в минут

	// Вычисление минимальной необходимой мощности для поддержания БПЛА в воздухе
	PropMinPower := float64(PropTractionConst) * float64(EnvAirPressure/1000) * math.Pow(PropMinFreq, 3) * math.Pow(float64(PropDiameter/1000), 5)
	// Эффективность пропеллера
	PropEffectivity := AssemblyWeight / float32(PropMinPower)
	// Мощность, расходуемая пропеллерами, то есть полезная мощность всей силовой установки, которая может включать несколько ВМГ
	UsefulMinPowerOfPlant := AssemblyWeight / PropEffectivity
	// Время полета в режиме висения
	MinTime := (BattEnergyReserve * n) / UsefulMinPowerOfPlant

	// МАКСИМАЛЬНЫЕ ХАРАКТЕРИСТИКИ

	return FlightProperties{
		Minimal: Mode{
			PropFreq:           float32(PropMinFreq),
			PropPower:          float32(PropMinPower),
			UsefulPowerOfPlant: UsefulMinPowerOfPlant,
			Time:               MinTime,
		},
	}
}
