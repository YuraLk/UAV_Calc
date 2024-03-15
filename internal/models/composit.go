package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Composit struct {
	gorm.Model
	Id        uint            `json:"ID" gorm:"primary_key"`
	Name      string          `json:"name" gorm:"size:32;not null"`
	CVC       json.RawMessage `json:"cvc" gorm:"type:jsonb;not null"` // ВАХ аккумулятора
	Batteries []Battery
}
