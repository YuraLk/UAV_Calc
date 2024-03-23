package response_properties

type HoverProperties struct {
	PropellerHangingLift  float64 `json:"propellerHangingLift"`
	PropellerSpeed        float64 `json:"propellerSpeed"`
	RPM                   float64 `json:"rpm"`
	PropellerAngularSpeed float64 `json:"propellerAngularSpeed"`
	MotorBackEMF          float64 `json:"motorBackEmf"`
	PowerForPropeller     float64 `json:"powerForPropeller"`
	AerodynamicCleanness  float64 `json:"aerodynamicCleanness"` // В то же время и условный КПД пропеллера в режиме висения
	AirSuctionSpeed       float64 `json:"airSuctionSpeed"`
	AirEjectionSpeed      float64 `json:"airEjectionSpeed"`
	MaximumEfficiency     float64 `json:"maximumEfficiency"`
	RealEfficiency        float64 `json:"realEfficiency"`
	MotorMechanicalPower  float64 `json:"motorMechanicalPower"`
	MotorElectricalPower  float64 `json:"motorElectricalPower"`
	GasLinear             float64 `json:"gasLinear"`
	MotorAverageCurrent   float64 `json:"motorAverageCurrent"`
	MotorVoltageUnderLoad float64 `json:"motorVoltageUnderLoad"`
	TimeOfFlight          float64 `json:"timeOfFlight"`
}
