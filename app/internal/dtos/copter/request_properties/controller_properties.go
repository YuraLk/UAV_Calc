package request_properties

type ControllerProperties struct {
	Mass               float32      `json:"mass" binding:"required,gt=0"`               // Масса контроллера(-ов). Общую масссу вычисляем на клиенте, исходя из выбранного типа (Кг)
	Voltage            float32      `json:"voltage" binding:"required,gt=0"`            // Максимальное напряжение контроллера. Для расчета на стороне клиента берется напрожение на банку и перемножается на S - коэффицент регулятора (В)
	InternalResistance float32      `json:"internalResistance" binding:"required,gt=0"` // Внутреннее сопротивление регулятора, (Ом)
	Current            CurrentRange `json:"currentRange" binding:"required"`            // Диапазон сил тока. Номинальная сила тока - 80% от максимальной. Ограничивает подаваемый ток на двигатель. В случае объединенного контроллера, сила тока на один выход делится на количество выходов под один двигатель
}

type CurrentRange struct {
	Per uint8 `json:"per" binding:"required,gt=0"` // Постоянное значение тока, (А)
	Max uint8 `json:"max" binding:"required,gt=0"` // Максимальное значениие тока, (А)
}
