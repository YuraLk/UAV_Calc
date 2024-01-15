package models

import "gorm.io/gorm"

type Composit struct {
	gorm.Model
	Id      uint   `json:"ID" gorm:"primary_key"`
	Name    string `json:"name" gorm:"size:32;not null"`
	Voltage struct {
		min float32
		max float32
	} `json:"voltage" gorm:"type:jsonb;not null"`
	C_rating   float32 `json:"c_rating" gorm:"not null"`
	Resistance struct {
		min float32
		max float32
	} `json:"resistance" gorm:"type:jsonb;not null"`
	Batteries []Battery
}
