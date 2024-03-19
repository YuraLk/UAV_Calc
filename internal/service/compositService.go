package service

import (
	"encoding/json"
	"errors"
	"fmt"
	"mime/multipart"

	"github.com/YuraLk/teca_server/internal/database/postgres"
	"github.com/YuraLk/teca_server/internal/exeptions"
	"github.com/YuraLk/teca_server/internal/models"
	"github.com/gin-gonic/gin"
)

type CompositService struct {
	C *gin.Context
}

func (S CompositService) Get() ([]models.Composit, error) {
	composits := []models.Composit{}
	if err := postgres.DB.Find(&composits).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return []models.Composit{}, err
	}

	return composits, nil
}

func (S CompositService) Create(Name string, File *multipart.FileHeader) (models.Composit, error) {
	var exist models.Composit // Сюда помещаем рузультаты поиска
	// Проверяем уникальность названия
	if err := postgres.DB.Where("name = ?", Name).First(&exist).Error; err == nil {
		err := errors.New("value is not unique")
		exeptions.BadRequest(S.C, fmt.Sprintf("Название %s уже существует!", Name), err)
		return models.Composit{}, err
	}

	CVC, err := FileService{}.ParseTableFromFile(File)

	if err != nil {
		exeptions.BadRequest(S.C, "Ошибка чтения файла!", err)
		return models.Composit{}, err
	}

	// Преобразование массива структур в JSON
	CVCJson, err := json.Marshal(CVC)
	if err != nil {
		exeptions.BadRequest(S.C, "Ошибка преобразования данных в JSON!", err)
		return models.Composit{}, err
	}

	// Создаем экземпляр
	composit := models.Composit{
		Name: Name,
		CVC:  CVCJson,
	}

	// Сохраняем ВАХ аккумулятора в БД
	if err := postgres.DB.Create(&composit).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return models.Composit{}, err
	}

	return composit, nil
}

func (S CompositService) Update(Id string, Name string, File *multipart.FileHeader) (models.Composit, error) {
	var composit models.Composit
	if err := postgres.DB.Where("id = ?", Id).First(&composit).Error; err != nil {
		exeptions.NotFound(S.C, "Запись не найдена!")
		return models.Composit{}, err
	}

	var exist models.Composit // Сюда помещаем рузультаты поиска
	// Проверяем уникальность названия
	if err := postgres.DB.Where("name = ?", Name).First(&exist).Error; err == nil {
		err := errors.New("value is not unique")
		exeptions.BadRequest(S.C, fmt.Sprintf("Название %s уже существует!", Name), err)
		return models.Composit{}, err
	}

	CVC, err := FileService{}.ParseTableFromFile(File)

	if err != nil {
		exeptions.BadRequest(S.C, "Ошибка чтения файла!", err)
		return models.Composit{}, err
	}

	// Преобразование массива структур в JSON
	CVCJson, err := json.Marshal(CVC)
	if err != nil {
		exeptions.BadRequest(S.C, "Ошибка преобразования данных в JSON!", err)
		return models.Composit{}, err
	}

	updateComposit := models.Composit{
		Id:   composit.Id,
		Name: Name,
		CVC:  CVCJson,
	}

	// Сохраняем ВАХ аккумулятора в БД
	if err := postgres.DB.Save(&updateComposit).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return models.Composit{}, err
	}

	return updateComposit, nil
}

func (S CompositService) Delete(Id string) error {
	if err := postgres.DB.Unscoped().Delete(&models.Composit{}, Id).Error; err != nil {
		exeptions.InternalServerError(S.C, err)
		return err
	}
	return nil
}
