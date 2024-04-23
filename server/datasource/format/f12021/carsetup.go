package f12021

type PacketCarSetupData struct {
	Header       PacketHeader     `json:"m_header"` // Header
	CarSetupData [22]CarSetupData `json:"m_carSetups"`
}

type CarSetupData struct {
	FrontWing              uint8   `json:"m_frontWing"`              // Front wing aero
	RearWing               uint8   `json:"m_rearWing"`               // Rear wing aero
	OnThrottle             uint8   `json:"m_onThrottle"`             // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   `json:"m_offThrottle"`            // Differential adjustment off throttle (percentage)
	FrontCamber            float32 `json:"m_frontCamber"`            // Front camber angle (suspension geometry)
	RearCamber             float32 `json:"m_rearCamber"`             // Rear camber angle (suspension geometry)
	FrontToe               float32 `json:"m_frontToe"`               // Front toe angle (suspension geometry)
	RearToe                float32 `json:"m_rearToe"`                // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   `json:"m_frontSuspension"`        // Front suspension
	RearSuspension         uint8   `json:"m_rearSuspension"`         // Rear suspension
	FrontAntiRollBar       uint8   `json:"m_frontAntiRollBar"`       // Front anti-roll bar
	RearAntiRollBar        uint8   `json:"m_rearAntiRollBar"`        // Front anti-roll bar
	FrontSuspensionHeight  uint8   `json:"m_frontSuspensionHeight"`  // Front ride height
	RearSuspensionHeight   uint8   `json:"m_rearSuspensionHeight"`   // Rear ride height
	BrakePressure          uint8   `json:"m_brakePressure"`          // Brake pressure (percentage)
	BrakeBias              uint8   `json:"m_brakeBias"`              // Brake bias (percentage)
	RearLeftTyrePressure   float32 `json:"m_rearLeftTyrePressure"`   // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 `json:"m_rearRightTyrePressure"`  // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 `json:"m_frontLeftTyrePressure"`  // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 `json:"m_frontRightTyrePressure"` // Front right tyre pressure (PSI)
	Ballast                uint8   `json:"m_ballast"`                // Ballast
	FuelLoad               float32 `json:"m_fuelLoad"`               // Fuel load
}

func (p *PacketCarSetupData) PacketHeader() Header {
	return &p.Header
}
