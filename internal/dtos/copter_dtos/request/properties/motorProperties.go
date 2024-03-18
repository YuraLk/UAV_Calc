package properties

type MotorProperties struct {
	KvConst            uint     `json:"kvConst" binding:"required,gt=0"`            // Констаннта количества оборотов в минуту, которые мотор может развить на вольт поданного напряжения без нагрузки
	WindingResistance  float64  `json:"windingResistance" binding:"required,gt=0"`  // Сопротивление обмоток двигателя (Ом)
	Mass               float32  `json:"mass" binding:"required,gt=0"`               // Масса двигателя (Кг)
	Currents           Currents `json:"currents" binding:"required"`                // Токи двигателя
	TorqueConst        float64  `json:"torqueConst" binding:"required,gt=0"`        // Константа крутящего момента, (Ньютон-метр/Ампер)
	Voltage            float32  `json:"voltage" binding:"required,gt=0"`            // Номинальное напряжение двигателя. Конечные цифры напряжения зависят от типа аккумлятора.
	Efficiency         float32  `json:"efficiency" binding:"required,gt=0"`         // КПД электродвигателя
	MomentInertia      float64  `json:"momentInertia" binding:"required,gt=0"`      // Момент инерции двигателя, (Кг/м^2)
	ElectricInductance float64  `json:"electricInductance" binding:"required,gt=0"` // Электрическая индуктивность, (Генри)
	MaxPower           uint     `json:"maxPower" binding:"required,gt=0"`           // Максимальная мощность, (Вт)
}

type Currents struct {
	NoLoadConst float64 `json:"noLoadConst" binding:"required,gt=0"` // Ток двигателя без нагрузки (А)
	Max         uint8   `json:"max" binding:"required,gt=0"`         // Максимальный ток двигателя (А)
}
