package models

import "gorm.io/gorm"

type Layout struct {
	gorm.Model
	Id          uint   `json:"ID" gorm:"primary_key"`
	Name        string `json:"name" gorm:"size:64;not null"`
	Multipler   uint8  `json:"multipler" gorm:"not null"`
	Controllers []Controller
}
