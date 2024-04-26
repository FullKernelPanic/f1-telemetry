package f12021

type PacketCarTelemetryData struct {
	Header                       PacketHeader         `json:"m_header"` // Header
	CarTelemetryData             [22]CarTelemetryData `json:"m_carTelemetryData"`
	MFDPanelIndex                uint8                `json:"m_mfdPanelIndex"`                // Index of MFD panel open - 255 = MFD closed, Single player, race – 0 = Car setup, 1 = Pits, 2 = Damage, 3 =  Engine, 4 = Temperatures, May vary depending on game mode
	MFDPanelIndexSecondaryPlayer uint8                `json:"m_mfdPanelIndexSecondaryPlayer"` // See above
	SuggestedGear                int8                 `json:"m_suggestedGear"`                // Suggested gear for the player (1-8), 0 if no gear suggested
}

func (p PacketCarTelemetryData) PacketHeader() Header {
	return &p.Header
}

type CarTelemetryData struct {
	Speed                   uint16     `json:"m_speed"`                   // Speed of car in kilometres per hour
	Throttle                float32    `json:"m_throttle"`                // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    `json:"m_steer"`                   // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    `json:"m_brake"`                   // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      `json:"m_clutch"`                  // Amount of clutch applied (0 to 100)
	Gear                    int8       `json:"m_gear"`                    // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     `json:"m_engineRPM"`               // Engine RPM
	DRS                     bool       `json:"m_drs"`                     // 0 = off, 1 = on
	RevLightsPercent        uint8      `json:"m_revLightsPercent"`        // Rev lights indicator (percentage)
	RevLightsBitValue       uint16     `json:"m_revLightsBitValue"`       // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       [4]uint16  `json:"m_brakesTemperature"`       // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   `json:"m_tyresSurfaceTemperature"` // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   `json:"m_tyresInnerTemperature"`   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     `json:"m_engineTemperature"`       // Engine temperature (celsius)
	TyresPressure           [4]float32 `json:"m_tyresPressure"`           // Tyres pressure (PSI)
	SurfaceType             [4]uint8   `json:"m_surfaceType"`             // Driving surface, see appendices
}
