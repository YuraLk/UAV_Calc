package models

import "gorm.io/gorm"

type CurrentRange struct {
	inv uint8
	max uint8
}

type Controller struct {
	gorm.Model
	Id         uint         `json:"ID" gorm:"primary_key"`
	Current    CurrentRange `json:"current" gorm:"type:jsonb;not null"`
	Voltage    float32      `json:"voltage" gorm:"not null"`
	Resistance float32      `json:"resistance" gorm:"not null"`
	Weight     uint         `json:"weight" gorm:"not null"`
	LayoutID   uint         `json:"layoutID" gorm:"not null"`
	Assembly   Assembly
}
