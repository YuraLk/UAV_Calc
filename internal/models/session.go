package models

import "gorm.io/gorm"

type Session struct {
	gorm.Model
	Id           uint   `json:"ID" gorm:"primary_key"`
	RefreshToken string `json:"refreshToken" gorm:"size:256;not null"`
	Device       string `json:"device" gorm:"size:128;not null"`
	UserID       uint   `json:"userID" gorm:"not null"`
}
