package controller

import (
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/service"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	// Рама
	EnvTemp   float32 `json:"env_temp" binding:"required"`
	EnvPress  float32 `json:"env_press" binding:"required"`
	RamaMass  uint64  `json:"rama_mass" binding:"required"`
	RamaVents uint8   `json:"rama_vents" binding:"required"` // Кол-во винтов
	// Аккумулятор
	AccVol     uint    `json:"acc_vol" binding:"required"`
	AccVoltage float32 `json:"acc_voltage" binding:"required"`
	AccOut     float32 `json:"acc_out" binding:"required"`
	AccMass    uint64  `json:"acc_mass" binding:"required"`  // Масса банки аккумулятора
	AccBanks   uint8   `json:"acc_banks" binding:"required"` // Кол-во банок
	AccCount   uint8   `json:"acc_count" binding:"required"` // Кол-во аккумуляторов с таким же кол-вом банок
	CompositID uint    `json:"compositID"`                   // Химический состав и сопутствующие свойства. Необзателены.
	// Регулятор
	ContCurrent    models.CurrentRange `json:"cont_current" binding:"required"`
	ContVoltage    float32             `json:"cont_voltage" binding:"required"`
	ContResistance float32             `json:"cont_resistance" binding:"required"`
	ContWeight     uint                `json:"cont_weight" binding:"required"`
	LayoutID       uint                `json:"layoutID" binding:"required"` // Компоновка ESC
	// Навесное оборудование
	EquipCurrent float32 `json:"equip_current" binding:"required"`
	EquipWeight  uint    `json:"equip_weight" binding:"required"`
	// Двигатель
	MotKv          uint8   `json:"mot_kv" binding:"required"`
	MotCurrent     uint8   `json:"mot_current" binding:"required"`
	MotVoltage     float32 `json:"mot_voltage" binding:"required"`
	MotPower       uint    `json:"mot_power" binding:"required"`
	MotPeakCurrent uint8   `json:"mot_peak_current" binding:"required"`
	MotResistance  float32 `json:"mot_resistance" binding:"required"`
	MotLength      float32 `json:"mot_length" binding:"required"`
	MotDiameter    float32 `json:"mot_diameter" binding:"required"`
	MotMagnets     uint8   `json:"mot_magnets" binding:"required"`
	MotWeight      uint    `json:"mot_weight" binding:"required"`
	// Пропеллер
	PropDiameter      float32 `json:"prop_diameter" binding:"required"`
	PropStep          float32 `json:"prop_step" binding:"required"`
	PropBlades        uint8   `json:"prop_blades" binding:"required"`
	PropTorsionAngle  float32 `json:"prop_torsion_angle" binding:"required"`
	PropGearRatio     float32 `json:"prop_gear_ratio" binding:"required"`
	PropPowerConst    float32 `json:"prop_power_const" binding:"required"`
	PropTractionConst float32 `json:"prop_traction_const" binding:"required"`
	PropWeight        uint    `json:"prop_weight" binding:"required"`
}

type CalculateResponse struct {
	TotalMass       uint64  `json:"total_mass"`
	EnvAirPressure  float32 `json:"env_air_pressure"`
	AccTotalVoltage float32 `json:"acc_total_voltage"`
	AccTotalVol     uint    `json:"acc_total_vol"`
	AccMaxOut       float32 `json:"acc_max_out"`
}

// Расчет характеристик
func Calculate(c *gin.Context) {
	var req CalculateRequest

	// Проверка валидации
	if err := c.ShouldBind(&req); err != nil {
		errors := utils.FormatErrors(err.Error())
		exeptions.UnprocessableEntity(c, errors)
		return
	}

	// Извлекаем данные из тела application/json
	c.ShouldBindJSON(&req)

	// Вычисление плотности воздуха
	EnvAirPressure := service.GetAirDensity(req.EnvTemp, req.EnvPress)
	// Вычисление характеристик аккумулятора
	AccTotalMass, AccTotalVoltage, AccTotalVol, AccMaxOut := service.GetAccFeatures(req.AccVol, req.AccVoltage, req.AccOut, req.AccMass, req.AccBanks, req.AccCount)
	// Вычисление массы ESC
	ContTolalMass, err := service.GetContMass(req.ContWeight, req.RamaVents, req.LayoutID)
	if err != nil {
		exeptions.NotFound(c, "Layout с данным ID не найден!")
		return
	}
	// Вычисление общей массы БПЛА
	TotalMass := service.GetTotalMass(req.RamaMass, AccTotalMass, uint64(ContTolalMass), uint64(req.EquipWeight), uint64(req.MotWeight*uint(req.RamaVents)), uint64(req.PropWeight*uint(req.RamaVents)))

	// Возвращаем вычисленные значения
	CalculateResponse := CalculateResponse{
		TotalMass:       TotalMass,
		EnvAirPressure:  EnvAirPressure,
		AccTotalVoltage: AccTotalVoltage,
		AccTotalVol:     AccTotalVol,
		AccMaxOut:       AccMaxOut,
	}

	c.JSON(200, &CalculateResponse)
}
