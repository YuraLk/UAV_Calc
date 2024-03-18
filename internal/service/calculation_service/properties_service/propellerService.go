package properties_service

import (
	"math"

	"github.com/YuraLk/teca_server/internal/consts"
	"github.com/YuraLk/teca_server/internal/dtos"
	request_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/request/properties"
	response_properties "github.com/YuraLk/teca_server/internal/dtos/copter_dtos/response/properties"
	"github.com/YuraLk/teca_server/internal/service/calculation_service/warning_service"
)

func GetPropellerProperties(propeller request_properties.PropellerProperties, frame request_properties.FrameProperties) (response_properties.PropellerProperties, *[]dtos.WarningDto) {

	// Радиус пропеллера, (М)
	var PropellerRadius float32 = propeller.Diameter / 2

	// Относительный шаг винта, (М)
	var RelativePitch float32 = propeller.Pitch / propeller.Diameter

	// Масса одной лопасти пропеллера,(Кг)
	var BladeMass float32 = propeller.Mass / float32(propeller.BladesNumber)

	// Масса всех пропеллеров, (Кг)
	var Mass float32 = propeller.Mass * float32(frame.PropellersNumber)

	// Площадь, ометаемая одним пропеллером, (М^2)
	var SweptArea float64 = consts.Pi * math.Pow(float64(PropellerRadius), 2)

	// Момент инерции пропеллера, (Кг·М^2)
	var MomentOfInertia float64 = float64(1) / float64(3) * float64(BladeMass) * math.Pow(float64(PropellerRadius), 2)

	// Аэродинамическое качество пропеллера
	var AerodynamicQuality float64 = math.Pow(float64(propeller.DimensionlessThrustConstant), (float64(3)/float64(2))) / float64(propeller.DimensionlessPowerConstant)

	// Собираем все предупреждения
	warnings := warning_service.AppendWarnings()

	// Возвращаем расчитанные параметры
	properties := response_properties.PropellerProperties{
		SweptArea:          SweptArea,
		Mass:               Mass,
		RelativePitch:      RelativePitch,
		MomentOfInertia:    MomentOfInertia,
		AerodynamicQuality: AerodynamicQuality,
	}

	return properties, warnings
}
