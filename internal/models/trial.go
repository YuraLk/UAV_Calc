package models

import (
	"time"

	"gorm.io/gorm"
)

type Trial struct {
	gorm.Model
	Id     uint      `json:"ID" gorm:"primary_key"`
	Before time.Time `json:"before" gorm:"not null"`
	UserID uint      `json:"userID" gorm:"not null"`
}
