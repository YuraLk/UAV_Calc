package service

import (
	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/YuraLk/teca_server/internal/types"
	"github.com/gin-gonic/gin"
)

func GetComposits(c *gin.Context) []models.Composit {
	composits := []models.Composit{}
	if err := postgres.DB.Find(&composits).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return []models.Composit{}
	}
	// fmt.Println(composits)
	// Преобразовывем в JSON - формат

	return composits
}

func CreateComposit(c *gin.Context, Name string, Voltage types.Voltage, CRating types.Current, SafeCapacity uint8) (models.Composit, error) {
	composit := models.Composit{
		Name: Name,
		Voltage: types.JSONB{
			"min": Voltage.Min,
			"nom": Voltage.Nom,
			"max": Voltage.Max,
		},
		CRating: types.JSONB{
			"per": CRating.Per,
			"max": CRating.Max,
		},
		SafeCapacity: SafeCapacity,
	}

	// Регистрируем пользователя
	if err := postgres.DB.Create(&composit).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return models.Composit{}, err
	}

	return composit, nil
}
