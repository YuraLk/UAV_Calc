package models

import "gorm.io/gorm"

type Attachments struct {
	gorm.Model
	Id               uint    `json:"ID" gorm:"primary_key"`
	Mass             float32 `json:"mass" gorm:"type:float;not null"`             // Массса навесного оборудования, (Кг)
	PowerConsumption float32 `json:"powerConsumption" gorm:"type:float;not null"` // Энергопотребление навесного оборудования, (Вт)
	Assembly         Assembly
}
