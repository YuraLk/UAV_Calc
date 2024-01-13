package models

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Id       uint   `json:"ID" gorm:"primary_key"`
	Name     string `json:"name" gorm:"size:128;not null"`
	Email    string `json:"email" gorm:"size:128;unique;not null"`
	Phone    string `json:"phone" gorm:"size:128;unique;not null"`
	Password string `json:"password" gorm:"size:256;not null"`
	Access   Access
	Sessions []Session
}
