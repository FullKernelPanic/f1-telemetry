package f12021

type PacketParticipantsData struct {
	Header          PacketHeader        `json:"m_header"`        // Header
	NumActiveCars   uint8               `json:"m_numActiveCars"` // Number of active cars in the data – should match number of cars on HUD
	ParticipantData [22]ParticipantData `json:"m_participants"`
}

func (p PacketParticipantsData) PacketHeader() Header {
	return &p.Header
}

type ParticipantData struct {
	AIControlled  bool     `json:"m_aiControlled"`  // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId      uint8    `json:"m_driverId"`      // Driver id - see appendix, 255 if network human
	NetworkId     uint8    `json:"m_networkId"`     // Network id – unique identifier for network players
	TeamId        uint8    `json:"m_teamId"`        // Team id - see appendix
	MyTeam        bool     `json:"m_myTeam"`        // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8    `json:"m_raceNumber"`    // Race number of the car
	Nationality   uint8    `json:"m_nationality"`   // Nationality of the driver
	Name          [48]byte `json:"m_name"`          // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
	YourTelemetry uint8    `json:"m_yourTelemetry"` // The player's UDP setting, 0 = restricted, 1 = public
}
