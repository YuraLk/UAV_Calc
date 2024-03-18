package hover_service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/dtos"
	copter_dtos "github.com/YuraLk/teca_server/internal/dtos/copter_dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
)

func GetHoverProperties(pr request.CalculateCopter, st copter_dtos.StandartProperties) (response_properties.HoverProperties, *[]dtos.WarningDto) {
	// Навесное оборудование
	// var attachments = props.AttachmentsProperties
	// ESC
	// var esc = props.ControllerProperties
	// Внешняя среда
	// var environment = props.EnvironmentProperties
	// Мотор
	// var motor = props.MotorProperties
	// Рама
	var frame = pr.FrameProperties
	// Пропеллер
	var propeller = pr.PropellerProperties
	// Аккумулятор
	// var battery = props.BatteryProperties

	// Подъемная сила каждого пропеллера, необходимая для поддержания ЛА в воздухе, (Н):
	var PropellerHangingLift float64 = float64(st.GeneralProperties.Weight) / float64(frame.PropellersNumber)

	// Частота вращения винта, (Гц)
	var PropellerSpeed float64 = math.Sqrt(PropellerHangingLift / (float64(propeller.DimensionlessPowerConstant) * st.EnvironmentProperties.AirDensity * math.Pow(float64(propeller.Diameter), 4)))

	warnings := warning_service.AppendWarnings()

	return response_properties.HoverProperties{
		PropellerHangingLift: PropellerHangingLift,
		PropellerSpeed:       PropellerSpeed,
	}, warnings
}
