package requests_properties

type BatteryProperties struct {
	CellCapacity         float32 `json:"cellCapacity" binding:"required,gt=0"`          // Емкость банки, (Ампер * Час)
	CellMass             float32 `json:"cellMass" binding:"required,gt=0"`              // Масса банки аккумулятора, (Кг)
	S                    uint8   `json:"s" binding:"required,gt=0"`                     // Кол-во последовательно соединенных ячеек
	P                    uint8   `json:"p" binding:"required,gt=0"`                     // Кол-во банок аккумулятора
	CRating              CRating `json:"cRating" binding:"required"`                    // С - рейтинг аккумулятора
	InternalResistance   float64 `json:"internalResistance" binding:"required,gt=0"`    // Внутреннее сопротивление аккумулятора, (Ом)
	MaxDischargePercent  float32 `json:"maxDischargePercent" binding:"gte=0.05,lte=1"`  // Максимальный процент разряда, от 5% до 100%
	InitialStateOfCharge float32 `json:"initialStateOfCharge" binding:"gte=0.05,lte=1"` // Изначальное состояние заряда аккумулятора, от 5% до 100%
	CompositId           uint64  `json:"compositId" binding:"required,gt=0"`            // Индентификатор химических свойств аккумулятора
}

type CRating struct {
	Per uint8 `json:"per" binding:"required,gt=0"`
	Max uint8 `json:"max" binding:"required,gt=0"`
}
