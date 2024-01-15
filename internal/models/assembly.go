package models

import "gorm.io/gorm"

type Assembly struct {
	gorm.Model
	Id           uint   `json:"ID" gorm:"primary_key"`
	Name         string `json:"name" gorm:"size:128;not null"`
	BaseID       uint   `json:"baseID" gorm:"not null"`
	UserID       uint   `json:"userID" gorm:"not null"`
	AtmosphereID uint   `json:"atmosphereID" gorm:"not null"`
	BatteryID    uint   `json:"batteryID" gorm:"not null"`
	EquipmentID  uint   `json:"equipmentID" gorm:"not null"`
	MotorID      uint   `json:"motorID" gorm:"not null"`
	ControllerID uint   `json:"controllerID" gorm:"not null"`
	PropellerID  uint   `json:"propellerID" gorm:"not null"`
}
