package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Controller struct {
	gorm.Model
	Id         uint        `json:"ID" gorm:"primary_key"`
	Current    types.JSONB `json:"current" gorm:"type:jsonb;not null"`
	Voltage    float32     `json:"voltage" gorm:"not null"`
	Resistance float32     `json:"resistance" gorm:"not null"`
	Weight     uint        `json:"weight" gorm:"not null"`
	LayoutID   uint        `json:"layoutID" gorm:"not null"`
	Assembly   Assembly
}
