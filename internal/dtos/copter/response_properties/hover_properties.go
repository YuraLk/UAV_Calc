package response_properties

type HoverProperties struct {
	PropellerHangingLift  float64 `json:"propellerHangingLift"`
	PropellerSpeed        float64 `json:"propellerSpeed"`
	PropellerEfficiency   float64 `json:"propellerEfficiency"`
	RPM                   float64 `json:"rpm"`
	PropellerAngularSpeed float64 `json:"propellerAngularSpeed"`
	MotorBackEMF          float64 `json:"motorBackEmf"`
	AirFlowSpeed          float64 `json:"airFlowSpeed"`
	PowerForPropeller     float64 `json:"powerForPropeller"`
}
