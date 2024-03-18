package copter_dtos

import "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"

// Свойства, независящие от режима полета
type StandartProperties struct {
	EnvironmentProperties properties.EnvironmentProperties
	BatteryProperties     properties.BatteryProperties
	PropellerProperties   properties.PropellerProperties
	MotorProperties       properties.MotorProperties
	GeneralProperties     properties.GeneralProperties
}
