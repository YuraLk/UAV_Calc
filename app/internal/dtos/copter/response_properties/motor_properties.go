package response_properties

type MotorProperties struct {
	Mass                   float32 `json:"mass"`
	PhaseValueOfEMFConst   float64 `json:"phaseValueOfEMFConst"`
	MaxPowerOfMotorOnBoard float64 `json:"maxPowerOfMotorOnBoard"`
	CharacteristicLength   float64 `json:"characteristicLength"`
}
