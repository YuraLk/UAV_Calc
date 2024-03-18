package requests

import requestsProperties "github.com/YuraLk/teca_server/internal/requests/properties"

type CalculateCopter struct {
	AttachmentsProperties requestsProperties.AttachmentsProperties `json:"attachmentsProperties" binding:"required"` // Навесное оборудование
	ControllerProperties  requestsProperties.ControllerProperties  `json:"controllerProperties" binding:"required"`  // Параметры контроллера
	EnvironmentProperties requestsProperties.EnvironmentProperties `json:"environmentProperties" binding:"required"` // Параметры окружения
	MotorProperties       requestsProperties.MotorProperties       `json:"motorProperties" binding:"required"`       // Параметры мотора
	FrameProperties       requestsProperties.FrameProperties       `json:"frameProperties" binding:"required"`       // Параметры рамы
	PropellerProperties   requestsProperties.PropellerProperties   `json:"propellerProperties" binding:"required"`   // Параметры пропеллеров
	BatteryProperties     requestsProperties.BatteryProperties     `json:"batteryProperties" binding:"required"`     // Параметры батареи
}
