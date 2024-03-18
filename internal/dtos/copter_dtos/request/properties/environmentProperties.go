package properties

type EnvironmentProperties struct {
	AltitudeRange  AltitudeRange `json:"altitudeRange" binding:"required"`  // Диапазон высот полета ЛА
	AirTemperature float64       `json:"airTemperature" binding:"required"` // Температура воздуха на высоте запуска ЛА, (K)
	AirHumidity    float64       `json:"airHumidity" binding:"gte=0,lte=1"` // Влажность воздуха
}

type AltitudeRange struct {
	Start  float64 `json:"start" binding:"required"`  // Высота запуска ЛА над уровнем моря, (М)
	Flight float64 `json:"flight" binding:"required"` // Высота полета ЛА над уровнем моря, (М)
}
