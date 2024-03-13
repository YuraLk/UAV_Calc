package models

import (
	"github.com/YuraLk/teca_server/internal/types"
	"gorm.io/gorm"
)

type Controller struct {
	gorm.Model
	Id                 uint        `json:"ID" gorm:"primary_key"`
	Mass               float32     `json:"mass" gorm:"type:float;not null"`               // Масса контроллера(-ов). Общую масссу вычисляем на клиенте, исходя из выбранного типа (Кг)
	Voltage            float32     `json:"voltage" gorm:"type:float;not null"`            // Максимальное напряжение контроллера. Для расчета на стороне клиента берется напрожение на банку и перемножается на S - коэффицент регулятора (В)
	InternalResistance float32     `json:"internalResistance" gorm:"type:float;not null"` // Внутреннее сопротивление регулятора, (Ом)
	Current            types.JSONB `json:"currentRange" gorm:"type:jsonb;not null"`       // Диапазон сил тока. Номинальная сила тока - 80% от максимальной. Ограничивает подаваемый ток на двигатель. В случае объединенного контроллера, сила тока на один выход делится на количество выходов под один двигатель
	LayoutID           uint        `json:"layoutID" gorm:"not null"`
	Assembly           Assembly
}
