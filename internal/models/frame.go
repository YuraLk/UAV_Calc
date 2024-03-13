package models

import "gorm.io/gorm"

type Frame struct {
	gorm.Model
	Id                  uint    `json:"ID" gorm:"primary_key"`
	Mass                float32 `json:"mass" gorm:"type:float;not null"`
	PropellersNumber    uint8   `json:"propellersNumber" gorm:"not null"`
	DiagonalSize        float32 `json:"diagonalSize" gorm:"type:float;not null"`
	RollAngleLimitation uint8   `json:"rollAngleLimitation" gorm:"not null"`
	LinkageID           uint    `json:"linkageID" gorm:"not null"`
	Assembly            Assembly
}
