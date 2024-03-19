package service

import (
	"encoding/csv"
	"errors"

	"mime/multipart"
	"strconv"
	"strings"

	dtos "github.com/YuraLk/teca_server/internal/dtos/battery_dtos"
)

type FileService struct{}

func (FileService) ParseTableFromFile(file *multipart.FileHeader) ([]dtos.BatteryDto, error) {

	// Открываем файл
	src, err := file.Open()

	if err != nil {
		return []dtos.BatteryDto{}, err
	}
	defer src.Close()

	// Создаем reader для CSV
	reader := csv.NewReader(src)
	reader.Comma = ';'
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	// Массив данных
	var CVC []dtos.BatteryDto

	for {
		record, err := reader.Read()
		if err != nil {
			break
		}
		// Конвертируем каждое значение в число и добавляем в массив
		ChargePercentageStr := strings.TrimSuffix(record[0], "%")              // Убираем знак %
		ChargePercentage, err := strconv.ParseUint(ChargePercentageStr, 10, 8) // Преобразуем в число
		if err != nil {
			continue
		}

		// Если в строке есть запятая, то она не может быть преобразована в число
		SmoothedVoltage, err := strconv.ParseFloat(strings.ReplaceAll(record[1], ",", "."), 64) // Преобразуем в число
		if err != nil {
			continue
		}

		LoadVoltage, err := strconv.ParseFloat(strings.ReplaceAll(record[2], ",", "."), 64) // Преобразуем в число
		if err != nil {
			continue
		}

		CVC = append(CVC, dtos.BatteryDto{
			ChargePercentage: uint8(ChargePercentage),
			SmoothedVoltage:  SmoothedVoltage,
			LoadVoltage:      LoadVoltage,
		})
	}
	// fmt.Println(len(CVC))

	// Проверяем, что длина массива CVC равна 100, то есть все записи были успешно прочитаны
	if len(CVC) != 101 {
		return []dtos.BatteryDto{}, errors.New("not all records were read correctly")
	}

	return CVC, nil
}
