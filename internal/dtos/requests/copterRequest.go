package requests

import requests_properties "github.com/YuraLk/teca_server/internal/dtos/requests/requests_properties"

type CalculateCopter struct {
	AttachmentsProperties requests_properties.AttachmentsProperties `json:"attachmentsProperties" binding:"required"` // Навесное оборудование
	ControllerProperties  requests_properties.ControllerProperties  `json:"controllerProperties" binding:"required"`  // Параметры контроллера
	EnvironmentProperties requests_properties.EnvironmentProperties `json:"environmentProperties" binding:"required"` // Параметры окружения
	MotorProperties       requests_properties.MotorProperties       `json:"motorProperties" binding:"required"`       // Параметры мотора
	FrameProperties       requests_properties.FrameProperties       `json:"frameProperties" binding:"required"`       // Параметры рамы
	PropellerProperties   requests_properties.PropellerProperties   `json:"propellerProperties" binding:"required"`   // Параметры пропеллеров
	BatteryProperties     requests_properties.BatteryProperties     `json:"batteryProperties" binding:"required"`     // Параметры батареи
}
