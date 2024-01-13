package models

import "gorm.io/gorm"

type Access struct {
	gorm.Model
	Id     uint   `json:"ID" gorm:"primary_key"`
	Role   string `json:"role" gorm:"size:16;not null"`
	UserID uint   `json:"userID" gorm:"not null"`
}
