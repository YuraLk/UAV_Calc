package models

import "gorm.io/gorm"

type Propeller struct {
	gorm.Model
	Id                          uint    `json:"ID" gorm:"primary_key"`
	Diameter                    float32 `json:"diameter" gorm:"type:float;not null"`                    // Диаметр пропеллера, (М)
	TorsionAngle                uint8   `json:"torsionAngle" gorm:"type:uint;not null"`                 // Угол кручения, (Град)
	Pitch                       float32 `json:"pitch" gorm:"type:float;not null"`                       // Шаг винта, (М)
	BladesNumber                uint8   `json:"bladesNumber" gorm:"type:uint;not null"`                 // Количество лопастей пропеллера
	DimensionlessPowerConstant  float32 `json:"dimensionlessPowerConstant" gorm:"type:float;not null"`  // Безразмерная константа мощности
	DimensionlessThrustConstant float32 `json:"dimensionlessThrustConstant" gorm:"type:float;not null"` // Безразмерная константа тяги
	GearRatio                   float32 `json:"gearRatio" gorm:"type:float;not null"`                   // Передаточное число
	Mass                        float32 `json:"mass" gorm:"type:float;not null"`                        // Масса пропеллера, (Кг)
	Assembly                    Assembly
}
