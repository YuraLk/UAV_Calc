package properties

type AttachmentsProperties struct {
	Mass             float32 `json:"mass" binding:"gte=0"`             // Массса навесного оборудования, (Кг)
	PowerConsumption float32 `json:"powerConsumption" binding:"gte=0"` // Энергопотребление навесного оборудования, (Вт)
}
