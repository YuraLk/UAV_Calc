package hover_service

import (
	dtos "github.com/YuraLk/teca_server/internal/dtos/copter_dtos"
	"github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetHoverProperties(props request.CalculateCopter, st dtos.StandartProperties) (response_properties.HoverProperties, *[]types.Warning) {
	// Навесное оборудование
	// var attachments = props.AttachmentsProperties
	// ESC
	// var esc = props.ControllerProperties
	// Внешняя среда
	// var environment = props.EnvironmentProperties
	// Мотор
	// var motor = props.MotorProperties
	// Рама
	var frame = props.FrameProperties
	// Пропеллер
	// var propeller = props.PropellerProperties
	// Аккумулятор
	// var battery = props.BatteryProperties

	// Подъемная сила каждого пропеллера, необходимая для поддержания ЛА в воздухе, (Н):
	var ProperllerHangingLift float32 = st.GeneralProperties.Weight / float32(frame.PropellersNumber)

	warnings := warning_service.AppendWarnings()

	return response_properties.HoverProperties{
		ProperllerHangingLift: ProperllerHangingLift,
	}, warnings
}
