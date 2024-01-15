package models

import "gorm.io/gorm"

type Model struct {
	gorm.Model
	Id             uint   `json:"ID" gorm:"primary_key"`
	Name           string `json:"name" gorm:"size:128;not null"`
	ManufacturerID uint   `json:"manufacturerID" gorm:"not null"`
	Motors         []Motor
}
