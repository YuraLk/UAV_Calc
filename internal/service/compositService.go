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

func CreateComposit(c *gin.Context, Name string, Voltage types.Voltage, CRating types.Current, SafeCapacity float32) (models.Composit, error) {
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

func UpdateComposit(c *gin.Context, Id uint, Name string, Voltage types.Voltage, CRating types.Current, SafeCapacity float32) (models.Composit, error) {
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", Id).First(&composit).Error; err != nil {
		exeptions.NotFound(c, "Запись не найдена!")
		return models.Composit{}, err
	}

	updateComposit := models.Composit{
		Id:   composit.Id,
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

	if err := postgres.DB.Save(&updateComposit).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return models.Composit{}, err
	}

	return updateComposit, nil
}

func DeleteComposit(c *gin.Context, Id string) error {
	if err := postgres.DB.Unscoped().Delete(&models.Composit{}, Id).Error; err != nil {
		exeptions.InternalServerError(c, err)
		return err
	}
	return nil
}
