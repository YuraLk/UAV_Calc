package properties

type FrameProperties struct {
	Mass                float32 `json:"mass" binding:"required,gt=0"`                        // Масса рамы, (Кг)
	PropellersNumber    uint8   `json:"propellersNumber" binding:"required,gt=0"`            // Кол-во пропеллеров
	DiagonalSize        float32 `json:"diagonalSize" binding:"required,gt=0"`                // Размер рамы по диагонали, (М)
	RollAngleLimitation uint8   `json:"rollAngleLimitation" binding:"required,min=0,max=80"` // Ограничение угла крена, (Град)
}
