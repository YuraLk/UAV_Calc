package models

import "gorm.io/gorm"

type Motor struct {
	gorm.Model
	Id         uint    `json:"ID" gorm:"primary_key"`
	Kv         uint8   `json:"kv" gorm:"not null"`
	Current    uint8   `json:"current" gorm:"not null"`
	Voltage    float32 `json:"voltage" gorm:"not null"`
	Power      uint32  `json:"power" gorm:"not null"`
	Resistance float32 `json:"resistance" gorm:"not null"`
	Length     float32 `json:"length" gorm:"not null"`
	Diameter   float32 `json:"diameter" gorm:"not null"`
	Magnets    uint8   `json:"magnets" gorm:"not null"`
	Weight     uint32  `json:"weight" gorm:"not null"`
	ModelID    uint    `json:"modelID" gorm:"not null"`
	Assembly   Assembly
}
