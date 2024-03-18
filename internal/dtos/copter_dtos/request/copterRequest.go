package request

import properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"

type CalculateCopter struct {
	AttachmentsProperties properties.AttachmentsProperties `json:"attachmentsProperties" binding:"required"` // Навесное оборудование
	ControllerProperties  properties.ControllerProperties  `json:"controllerProperties" binding:"required"`  // Параметры контроллера
	EnvironmentProperties properties.EnvironmentProperties `json:"environmentProperties" binding:"required"` // Параметры окружения
	MotorProperties       properties.MotorProperties       `json:"motorProperties" binding:"required"`       // Параметры мотора
	FrameProperties       properties.FrameProperties       `json:"frameProperties" binding:"required"`       // Параметры рамы
	PropellerProperties   properties.PropellerProperties   `json:"propellerProperties" binding:"required"`   // Параметры пропеллеров
	BatteryProperties     properties.BatteryProperties     `json:"batteryProperties" binding:"required"`     // Параметры батареи
}
