package models

import "gorm.io/gorm"

type Propeller struct {
	gorm.Model
	Id            uint    `json:"ID" gorm:"primary_key"`
	Diameter      float32 `json:"diameter" gorm:"not null"`
	Step          float32 `json:"step" gorm:"not null"`
	Blades        uint8   `json:"blades" gorm:"not null"`
	TorsionAngle  float32 `json:"torsionAngle" gorm:"not null"`
	GearRatio     float32 `json:"gearRatio" gorm:"not null"`
	PowerConst    float32 `json:"powerConst" gorm:"not null"`
	TractionConst float32 `json:"tractionConst" gorm:"not null"`
	Weight        uint    `json:"weight" gorm:"not null"`
	Assembly      Assembly
}
