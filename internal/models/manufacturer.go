package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	Id     uint   `json:"ID" gorm:"primary_key"`
	Name   string `json:"name" gorm:"size:128;not null"`
	Models []Model
}
