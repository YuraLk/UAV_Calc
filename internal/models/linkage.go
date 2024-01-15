package models

import "gorm.io/gorm"

type Linkage struct {
	gorm.Model
	Id    uint   `json:"ID" gorm:"primary_key"`
	Value string `json:"value" gorm:"size:64;not null"`
	Bases []Base
}
