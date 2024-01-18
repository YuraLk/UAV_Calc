package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Composit struct {
	gorm.Model
	Id           uint        `json:"ID" gorm:"primary_key"`
	Name         string      `json:"name" gorm:"size:32;not null"`
	Voltage      types.JSONB `json:"voltage" gorm:"type:jsonb;not null"`
	CRating      types.JSONB `json:"c_rating" gorm:"type:jsonb;not null"`
	SafeCapacity float32     `json:"safe_capacity" gorm:"not null"`
	Batteries    []Battery
}
