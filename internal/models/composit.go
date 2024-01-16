package models

import "gorm.io/gorm"

type ResistanceRange struct {
	min float32
	max float32
}

type VoltageRange struct {
	min float32
	nom float32
	max float32
}

type CRatingRange struct {
	inv uint
	max uint
}

type Composit struct {
	gorm.Model
	Id           uint         `json:"ID" gorm:"primary_key"`
	Name         string       `json:"name" gorm:"size:32;not null"`
	Voltage      VoltageRange `json:"voltage" gorm:"type:jsonb;not null"`
	C_rating     CRatingRange `json:"c_rating" gorm:"type:jsonb;not null"`
	SafeCapacity float32      `json:"safe_capacity" gorm:"not null"`
	Batteries    []Battery
}
