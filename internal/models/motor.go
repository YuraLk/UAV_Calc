package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Motor struct {
	gorm.Model
	Id                         uint        `json:"ID" gorm:"primary_key"`
	KvConst                    uint8       `json:"kvConst" gorm:"type:uint;not null"`            // Констаннта количества оборотов в минуту, которые мотор может развить на вольт поданного напряжения без нагрузки
	WindingResistance          float32     `json:"windingResistance" gorm:"type:float;not null"` // Сопротивление обмоток двигателя (Ом)
	MagnetsNumber              uint8       `json:"magnetsNumber" gorm:"type:uint;not null"`      // Количество магнитов
	Mass                       float32     `json:"mass" gorm:"type:float;not null"`              // Масса двигателя (Кг)
	Current                    types.JSONB `json:"current" gorm:"type:jsonb;not null"`
	TorqueProportionalityConst float32     `json:"torqueProportionalityConst" gorm:"type:float;not null"` // Константа пропорциональности крутящего момента двигателя
	Voltage                    float32     `json:"voltage" gorm:"type:float;not null"`                    // Номинальное напряжение двигателя. Конечные цифры напряжения зависят от типа аккумлятора.
	Efficiency                 float32     `json:"efficiency" gorm:"type:float;not null"`                 // КПД электродвигателя
	MomentInertia              float32     `json:"momentInertia" gorm:"type:float;not null"`              // Момент инерции двигателя, (Кг/м^2)
	ElectricInductance         float32     `json:"electricInductance" gorm:"type:float;not null"`         // Электрическая индуктивность, (Генри)
	MaxPower                   uint        `json:"maxPower" gorm:"type:uint;not null"`                    // Максимальная мощность
	ModelID                    uint        `json:"modelID" gorm:"not null"`
	Assembly                   Assembly
}
