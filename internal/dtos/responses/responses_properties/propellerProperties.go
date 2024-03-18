package responses_properties

type PropellerProperties struct {
	SweptArea          float64 `json:"sweptArea"`
	Mass               float32 `json:"mass"`
	RelativePitch      float32 `json:"relativePitch"`
	MomentOfInertia    float64 `json:"momentOfInertia"`
	AerodynamicQuality float64 `json:"aerodynamicQuality"`
}
