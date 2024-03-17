package requests

import requestsProperties "github.com/YuraLk/teca_server/internal/requests/properties"

type CalculateCopter struct {
	AttachmentsProperties requestsProperties.AttachmentsProperties `json:"attachmentsProperties" binding:"required"` // Навесное оборудование
	ControllerProperties  requestsProperties.ControllerProperties  `json:"controllerProperties" binding:"required"`  // Параметры контроллера
	EnvironmentProperties requestsProperties.EnvironmentProperties `json:"environmentProperties" binding:"required"` // Параметры окружения
}
