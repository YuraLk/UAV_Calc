package models

import "gorm.io/gorm"

type Battery struct {
	gorm.Model
	Id         uint    `json:"ID" gorm:"primary_key"`
	Banks      uint8   `json:"banks" gorm:"not null"`
	Number     uint8   `json:"number" gorm:"not null"`
	Capacity   uint32  `json:"capacity" gorm:"not null"` // Вместимоть банки
	Resistance float32 `json:"resistance" gorm:"not null"`
	Voltage    struct {
		min float32
		max float32
	} `json:"voltage" gorm:"type:jsonb;not null"`
	C_rating struct {
		inv float32
		max float32
	} `json:"c_rating" gorm:"type:jsonb;not null"`
	Weight     uint64 `json:"weight" gorm:"not null"`
	CompositID uint   `json:"compositID" gorm:"not null"`
	Assembly   Assembly
}
