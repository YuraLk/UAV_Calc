package service

import (
	"github.com/YuraLk/teca_server/internal/models"
	requests "github.com/YuraLk/teca_server/internal/requests/properties"
	responses "github.com/YuraLk/teca_server/internal/responses/properties"
	"github.com/YuraLk/teca_server/internal/types"
)

func GetBatteryProperties(obj requests.BatteryProperties, composit models.Composit) (responses.BatteryProperties, *[]types.Warning) {

	return responses.BatteryProperties{}, nil
}
