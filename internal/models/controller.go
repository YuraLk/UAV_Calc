package models

import "gorm.io/gorm"

type Controller struct {
	gorm.Model
	Id      uint `json:"ID" gorm:"primary_key"`
	Current struct {
		inv uint8
		max uint8
	} `json:"current" gorm:"type:jsonb;not null"`
	Voltage  float32 `json:"voltage" gorm:"not null"`
	Weight   uint    `json:"weight" gorm:"not null"`
	Assembly Assembly
}
