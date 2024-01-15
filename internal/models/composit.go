package models

import "gorm.io/gorm"

type ResistanceRange struct {
	min float32
	max float32
}

type VoltageRange struct {
	min float32
	max float32
}

type Composit struct {
	gorm.Model
	Id         uint            `json:"ID" gorm:"primary_key"`
	Name       string          `json:"name" gorm:"size:32;not null"`
	Voltage    VoltageRange    `json:"voltage" gorm:"type:jsonb;not null"`
	C_rating   float32         `json:"c_rating" gorm:"not null"`
	Resistance ResistanceRange `json:"resistance" gorm:"type:jsonb;not null"`
	Batteries  []Battery
}
