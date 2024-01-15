package controller

import (
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/utils"
	"github.com/gin-gonic/gin"
)

type CalculateRequest struct {
	// Рама
	EnvTemp   float32 `json:"env_temp" binding:"required"`
	EnvPress  float32 `json:"env_press" binding:"required"`
	RamaMass  uint64  `json:"rama_mass" binding:"required"`
	RamaVents uint8   `json:"rama_vents" binding:"required"`
	// Аккуиулятор
	AccVol     uint32 `json:"acc_vol" binding:"required"`
	AccVoltage struct {
		min float32
		max float32
	} `json:"acc_voltage" binding:"required"`
	AccOut struct {
		inv float32
		max float32
	} `json:"acc_out" binding:"required"`
	AccMass    uint64 `json:"acc_mass" binding:"required"`
	AccBanks   uint8  `json:"acc_banks" binding:"required"`
	AccCount   uint8  `json:"acc_count" binding:"required"`
	CompositID uint   `json:"compositID"` // Тип аккумулятора. Необзателен.
	// Регулятор
	ContCurrent struct {
		inv uint8
		max uint8
	} `json:"cont_current" binding:"required"`
	ContVoltage    float32 `json:"cont_voltage" binding:"required"`
	ContResistance float32 `json:"cont_resistance" binding:"required"`
	ContWeight     uint    `json:"cont_weight" binding:"required"`
	// Навесное оборудование
	EquipCurrent float32 `json:"equip_current" binding:"required"`
	EquipWeight  uint    `json:"equip_weight" binding:"required"`
	// Двигатель
	MotKv          uint8   `json:"mot_kv" binding:"required"`
	MotCurrent     uint8   `json:"mot_current" binding:"required"`
	MotVoltage     float32 `json:"mot_voltage" binding:"required"`
	MotPower       uint32  `json:"mot_power" binding:"required"`
	MotPeakCurrent uint8   `json:"mot_peak_current" binding:"required"`
	MotResistance  float32 `json:"mot_resistance" binding:"required"`
	MotLength      float32 `json:"mot_length" binding:"required"`
	MotDiameter    float32 `json:"mot_diameter" binding:"required"`
	MotMagnets     uint8   `json:"mot_magnets" binding:"required"`
	MotWeight      uint32  `json:"mot_weight" binding:"required"`
	// Пропеллер
	PropDiameter      float32 `json:"prop_diameter" binding:"required"`
	PropStep          float32 `json:"prop_step" binding:"required"`
	PropBlades        uint8   `json:"prop_blades" binding:"required"`
	PropTorsionAngle  float32 `json:"prop_torsion_angle" binding:"required"`
	PropGearRatio     float32 `json:"prop_gear_ratio" binding:"required"`
	PropPowerConst    float32 `json:"prop_power_const" binding:"required"`
	PropTractionConst float32 `json:"prop_traction_const" binding:"required"`
	PropWeight        uint32  `json:"prop_weight" binding:"required"`
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

}
