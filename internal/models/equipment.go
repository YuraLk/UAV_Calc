package models

import "gorm.io/gorm"

type Equipment struct {
	gorm.Model
	Id       uint    `json:"ID" gorm:"primary_key"`
	Current  float32 `json:"current" gorm:"not null"`
	Weight   uint32  `json:"weight" gorm:"not null"`
	Assembly Assembly
}
