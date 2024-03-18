package requests_properties

type PropellerProperties struct {
	Diameter                    float32 `json:"diameter" binding:"required,gt=0"`                    // Диаметр пропеллера, (М)
	TorsionAngle                uint8   `json:"torsionAngle" binding:"gte=0"`                        // Угол кручения, (Град)
	Pitch                       float32 `json:"pitch" binding:"required,gt=0"`                       // Шаг винта, (М)
	BladesNumber                uint8   `json:"bladesNumber" binding:"required,gte=2"`               // Количество лопастей пропеллера.
	DimensionlessPowerConstant  float32 `json:"dimensionlessPowerConstant" binding:"required,gt=0"`  // Безразмерная константа мощности
	DimensionlessThrustConstant float32 `json:"dimensionlessThrustConstant" binding:"required,gt=0"` // Безразмерная константа тяги
	GearRatio                   float32 `json:"gearRatio" binding:"required,gt=0"`                   // Передаточное число
	Mass                        float32 `json:"mass" binding:"required,gt=0"`                        // Масса пропеллера, (Кг)
}
