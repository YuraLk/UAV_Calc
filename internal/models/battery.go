package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Battery struct {
	gorm.Model
	Id                     uint        `json:"ID" gorm:"primary_key"`
	CellCapacity           float32     `json:"cellCapacity" gorm:"type:float;not null"`           // Емкость банки, (Ампер * Час)
	CellMass               float32     `json:"cellMass" gorm:"type:float;not null"`               // Масса банки аккумулятора, (Кг)
	S                      uint8       `json:"s" gorm:"type:uint;not null"`                       // Кол-во последовательно соединенных ячеек
	P                      uint8       `json:"p" gorm:"type:uint;not null"`                       // Кол-во банок аккумулятора
	CRating                types.JSONB `json:"cRating" gorm:"type:jsonb;not null"`                // С - рейтинг аккумулятора
	InternalResistance     float64     `json:"internalResistance" gorm:"type:float;not null"`     // Внутреннее сопротивление аккумулятора, (Ом)
	CellChemicalProperties types.JSONB `json:"cellChemicalProperties" gorm:"type:jsonb;not null"` // Химические свойства аккумулятора
	MaxDischargePercent    uint8       `json:"maxDischargePercent" gorm:"type:uint;not null"`     // Максимальный процент разряда, от 5% до 100%
	InitialStateOfCharge   uint8       `json:"initialStateOfCharge" gorm:"type:uint;not null"`    // Изначальное состояние заряда аккумулятора,(%)
	CompositID             uint        `json:"compositID" gorm:"not null"`
	Assembly               Assembly
}
