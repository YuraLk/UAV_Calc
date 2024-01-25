package controller

import (
	"fmt"

	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	// Окружающая среда
	EnvTemp  float32 `json:"env_temp" binding:"required"`
	EnvPress float32 `json:"env_press" binding:"required"`
	// Рама
	FrameMass  uint64 `json:"frame_mass" binding:"required"`
	AxisNumber uint8  `json:"axis_number" binding:"required"` // Кол-во винтов
	// Аккумулятор
	CellCapacity float32       `json:"cell_capacity" binding:"required"` // Емкость банки в Ампер - Часах
	CellMass     uint64        `json:"cell_mass" binding:"required"`     // Масса банки аккумулятора
	BattTypeS    uint8         `json:"batt_type_s" binding:"required"`   // Кол-во банок
	BattTypeP    uint8         `json:"batt_type_p" binding:"required"`   // Кол-во аккумуляторов с таким же кол-вом банок
	CellCRating  types.Current `json:"cell_c_rating" binding:"required"`
	CompositID   uint          `json:"compositID"` // Химический состав и сопутствующие свойства.
	// Регулятор
	RegCurrent    types.Current `json:"reg_current" binding:"required"`
	RegVoltage    float32       `json:"reg_voltage" binding:"required"`
	RegResistance float32       `json:"reg_resistance" binding:"required"`
	RegMass       uint          `json:"reg_mass" binding:"required"`
	LayoutID      uint          `json:"layoutID" binding:"required"` // Компоновка ESC
	// Навесное оборудование
	EquipCurrent float32 `json:"equip_current" binding:"required"` // Потребление навесного оборудования в Ампер - Часах
	EquipMass    uint    `json:"equip_mass" binding:"required"`
	// Двигатель
	MotorKv             uint    `json:"motor_kv" binding:"required"`
	MotorNoLoadCurrent  float32 `json:"motor_no_load_current" binding:"required"` // Постоянная тока холостого хода
	MotorMaxVoltage     float32 `json:"motor_max_voltage" binding:"required"`     // Напряжение без нагрузки
	MotorMaxPower       uint    `json:"motor_max_power" binding:"required"`
	MotorPeakCurrent    float32 `json:"motor_peak_current" binding:"required"`
	MotorResistance     float32 `json:"motor_resistance" binding:"required"`      // Oм
	MotorStatorLength   float32 `json:"motor_stator_length" binding:"required"`   // Мм
	MotorStatorDiameter float32 `json:"motor_stator_diameter" binding:"required"` // Мм
	MotorMagnets        uint8   `json:"motor_magnets" binding:"required"`
	MotorMass           uint    `json:"motor_mass" binding:"required"` // Грамм
	// Пропеллер
	PropDiameter      float32 `json:"prop_diameter" binding:"required"` // Мм
	PropStep          float32 `json:"prop_step" binding:"required"`     // Мм
	PropBlades        uint8   `json:"prop_blades" binding:"required"`
	PropTorsionAngle  float32 `json:"prop_torsion_angle"`                     // Угол кручения. Град. По умолчанию равен 0
	PropGearRatio     float32 `json:"prop_gear_ratio" binding:"required"`     // Передаточное число
	PropPowerConst    float32 `json:"prop_power_const" binding:"required"`    // Безразмерный коэффициент мощности
	PropTractionConst float32 `json:"prop_traction_const" binding:"required"` // Безразмерный коэффициент тяги
	PropMass          float32 `json:"prop_mass" binding:"required"`           // Грамм
}

type CalculateResponse struct {
	// Характеристики общей сборки
	AssemblyMass   uint64  `json:"assembly_mass"`
	AssemblyWeight float32 `json:"assembly_weight"`
	EnvAirPressure float32 `json:"env_air_pressure"`
	// Характеристики аккумулятора
	BattMass              uint64  `json:"batt_mass"`
	BattCapacity          float32 `json:"batt_capacity"`
	BattCurrPer           float32 `json:"batt_curr_per"`
	BattCurrMax           float32 `json:"batt_curr_max"`
	BattAvailableCapacity float32 `json:"batt_available_capacity"`
	BattMinVoltage        float32 `json:"batt_min_voltage"`
	BattNomVoltage        float32 `json:"batt_nom_voltage"`
	BattMaxVoltage        float32 `json:"batt_max_voltage"`
	BattEnergy            float32 `json:"batt_energy"`
	BattSpecificEnergyVol float32 `json:"batt_specific_energy_vol"`
	BattEnergyReserve     float32 `json:"batt_energy_reserve"`
	// Характеристики двигателя
	MotElectricPower   float32 `json:"mot_electric_power"`
	MotMechanicalPower float32 `json:"mot_mechanical_power"`
}

type Mode struct {
	PropFreq           float32 `json:"prop_freq"`
	PropPower          float32 `json:"prop_power"`
	UsefulPowerOfPlant float32 `json:"useful_power_of_plant"`
	Time               float32 `json:"time"`
}

// Расчет характеристик
func Calculate(c *gin.Context) {
	var req CalculateRequest

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		fmt.Println(err)
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	// Вычисление плотности воздуха
	EnvAirPressure := service.GetAirDensity(req.EnvTemp, req.EnvPress)
	// Вычисление характеристик аккумулятора
	AP, err := service.GetBattFeatures(c, req.CellCapacity, req.CellCRating, req.CellMass, req.BattTypeS, req.BattTypeP, req.CompositID)
	if err != nil {
		return
	}
	// Вычисление массы ESC
	ContTolalMass, err := service.GetContMass(c, req.RegMass, req.AxisNumber, req.LayoutID)
	if err != nil {
		return
	}
	// Вычисление общей массы БПЛА
	AssemblyMass := service.GetAssemblyMass(req.FrameMass, AP.BattMass, uint64(ContTolalMass), uint64(req.EquipMass), uint64(req.MotorMass*uint(req.AxisNumber)), uint64(req.PropMass*float32(req.AxisNumber)))

	// Общий вес БПЛА в Ньютонах, за исключением массы аккумулятора
	AssemblyWeight := service.GetAssemblyWeight(AssemblyMass, AP.BattMass)

	// Вычисление характеристик мотора
	MP := service.GetMotorFeatures(req.MotorPeakCurrent, AP.BattMaxVoltage)

	// Вычисление характеристик пропеллера

	// Возвращаем вычисленные значения
	CalculateResponse := CalculateResponse{
		AssemblyMass:          AssemblyMass,
		AssemblyWeight:        AssemblyWeight,
		EnvAirPressure:        EnvAirPressure,
		BattMass:              AP.BattMass,
		BattCapacity:          AP.BattCapacity,
		BattCurrPer:           AP.BattCurrPer,
		BattCurrMax:           AP.BattCurrMax,
		BattAvailableCapacity: AP.BattAvailableCapacity,
		BattMinVoltage:        AP.BattMinVoltage,
		BattNomVoltage:        AP.BattNomVoltage,
		BattMaxVoltage:        AP.BattMaxVoltage,
		BattEnergy:            AP.BattEnergy,
		BattSpecificEnergyVol: AP.BattSpecificEnergyVol,
		BattEnergyReserve:     AP.BattEnergyReserve,
		MotElectricPower:      MP.MotElectricPower,
		MotMechanicalPower:    MP.MotMechanicalPower,
	}

	c.JSON(200, &CalculateResponse)
}
