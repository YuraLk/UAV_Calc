package service

import (
	"encoding/csv"
	"errors"

	// "fmt"
	"mime/multipart"
	"strconv"
	"strings"
)

type BatteryData struct {
	ChargePercentage uint8
	SmoothedVoltage  float64
	LoadVoltage      float64
}

func replaceCommaWithDot(str string) string {
	return strings.ReplaceAll(str, ",", ".")
}

func ParseTableFromFile(file *multipart.FileHeader) ([]BatteryData, error) {

	// Открываем файл
	src, err := file.Open()

	if err != nil {
		return []BatteryData{}, err
	}
	defer src.Close()

	// Создаем reader для CSV
	reader := csv.NewReader(src)
	reader.Comma = ';'
	reader.FieldsPerRecord = 3
	reader.Comment = '#'

	// Массив данных
	var CVC []BatteryData

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
		SmoothedVoltage, err := strconv.ParseFloat(replaceCommaWithDot(record[1]), 64) // Преобразуем в число
		if err != nil {
			continue
		}

		LoadVoltage, err := strconv.ParseFloat(replaceCommaWithDot(record[2]), 64) // Преобразуем в число
		if err != nil {
			continue
		}

		CVC = append(CVC, BatteryData{
			ChargePercentage: uint8(ChargePercentage),
			SmoothedVoltage:  SmoothedVoltage,
			LoadVoltage:      LoadVoltage,
		})
	}
	// fmt.Println(len(CVC))

	// Проверяем, что длина массива CVC равна 100, то есть все записи были успешно прочитаны
	if len(CVC) != 101 {
		return []BatteryData{}, errors.New("not all records were read correctly")
	}

	return CVC, nil
}
