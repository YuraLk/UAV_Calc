package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Composit struct {
	gorm.Model
	Id        uint        `json:"ID" gorm:"primary_key"`
	Name      string      `json:"name" gorm:"size:32;not null"`
	CVC       types.JSONB `json:"cvc" gorm:"type:jsonb;not null"` // ВАХ аккумулятора
	File      string      `gorm:"not null"`                       // Исходный файл
	Batteries []Battery
}
