package models

import "gorm.io/gorm"

type Assembly struct {
	gorm.Model
	Id            uint   `json:"ID" gorm:"primary_key"`
	Name          string `json:"name" gorm:"size:128;not null"`
	FrameID       uint   `json:"frameID" gorm:"not null"`
	UserID        uint   `json:"userID" gorm:"not null"`
	EnvironmentID uint   `json:"environmentID" gorm:"not null"`
	BatteryID     uint   `json:"batteryID" gorm:"not null"`
	AttachmentsID uint   `json:"attachmentsID" gorm:"not null"`
	MotorID       uint   `json:"motorID" gorm:"not null"`
	ControllerID  uint   `json:"controllerID" gorm:"not null"`
	PropellerID   uint   `json:"propellerID" gorm:"not null"`
}
