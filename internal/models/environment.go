package models

import (
	"encoding/json"

	"gorm.io/gorm"
)

type Environment struct {
	gorm.Model
	Id             uint            `json:"ID" gorm:"primary_key"`
	Altitude       json.RawMessage `json:"altitude" gorm:"type:jsonb;not null"`       // Высоты
	AirTemperature float32         `json:"airTemperature" gorm:"type:float;not null"` // Температура воздуха на высоте запуска ЛА, (K)
	AirHumidity    float32         `json:"airHumidity" gorm:"type:float;not null"`    // Влажность воздуха
	Assembly       Assembly
}
