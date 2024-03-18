package types

type BatteryData struct {
	ChargePercentage uint8   `json:"chargePercentage"`
	SmoothedVoltage  float64 `json:"smoothedVoltage"`
	LoadVoltage      float64 `json:"loadVoltage"`
}
