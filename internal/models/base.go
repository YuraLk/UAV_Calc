package models

import "gorm.io/gorm"

type Base struct {
	gorm.Model
	Id         uint   `json:"ID" gorm:"primary_key"`
	Weight     uint64 `json:"weight" gorm:"not null"`
	Propellers uint8  `json:"propellers" gorm:"not null"`
	Size       uint64 `json:"size" gorm:"not null"`
	Roll       uint8  `json:"roll" gorm:"not null"`
	LinkageID  uint   `json:"linkageID" gorm:"not null"`
	Assembly   Assembly
}
