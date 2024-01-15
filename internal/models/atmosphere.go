package models

import "gorm.io/gorm"

type Atmosphere struct {
	gorm.Model
	Id          uint   `json:"ID" gorm:"primary_key"`
	Height      int64  `json:"height" gorm:"not null"`
	Temperature int8   `json:"temperature" gorm:"not null"`
	Pressure    uint32 `json:"pressure" gorm:"not null"`
	Assembly    Assembly
}
