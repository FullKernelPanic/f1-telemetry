package f12021

type PacketEventData struct {
	Header          PacketHeader `json:"m_packetHeader"`    // Header
	EventStringCode [4]uint8     `json:"m_eventStringCode"` // Event string code, see below
}

func (p PacketEventData) PacketHeader() Header {
	return &p.Header
}

func (p PacketEventData) EventString() string {
	return string(p.EventStringCode[0:])
}

type FastestLapPacket struct {
	EventData PacketEventData `json:"m_packetEventData"`

	VehicleIdx uint8   `json:"vehicleIdx"` // Vehicle index of car achieving fastest lap
	LapTime    float32 `json:"lapTime"`    // Lap time is in seconds
}

type RetirementPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of car retiring
}

type TeamMateInPitsPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of team mate
}

type RaceWinnerPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the race winner
}

type PacketPenalty struct {
	EventData        PacketEventData `json:"m_packetEventData"`
	PenaltyType      uint8           `json:"penaltyType"`      // Penalty type – see Appendices
	InfringementType uint8           `json:"infringementType"` // Infringement type – see Appendices
	VehicleIndex     uint8           `json:"vehicleIdx"`       // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8           `json:"otherVehicleIdx"`  // Vehicle index of the other car involved
	Time             uint8           `json:"time"`             // Time gained, or time spent doing action in seconds
	LapNum           uint8           `json:"lapNum"`           // Lap the penalty occurred on
	PlacesGained     uint8           `json:"placesGained"`     // Number of places gained by this
}

type SpeedTrapPacket struct {
	EventData               PacketEventData `json:"m_packetEventData"`
	VehicleIndex            uint8           `json:"vehicleIdx"`              // Vehicle index of the vehicle triggering speed trap
	Speed                   float32         `json:"speed"`                   // Top speed achieved in kilometres per hour
	OverallFastestInSession uint8           `json:"overallFastestInSession"` // Overall fastest speed in session = 1, otherwise 0
	DriverFastestInSession  uint8           `json:"driverFastestInSession"`  // Fastest speed for driver in session = 1, otherwise 0
}

type StartLightsPacket struct {
	EventData PacketEventData `json:"m_packetEventData"`
	NumLights uint8           `json:"numLights"` // Number of lights showing
}

type DriveThroughPenaltyServedPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the vehicle serving drive through
}

type StopGoPenaltyServedPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the vehicle serving stop go
}

type FlashbackPacket struct {
	EventData                PacketEventData `json:"m_packetEventData"`
	FlashbackFrameIdentifier uint32          `json:"flashbackFrameIdentifier"` // Frame identifier flashed back to
	FlashbackSessionTime     float32         `json:"flashbackSessionTime"`     // Session time flashed back to
}

type PacketButtons struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	ButtonStatus uint32          `json:"m_buttonStatus"` // Bit flags specifying which buttons are being pressed currently - see appendices
}

func (p SpeedTrapPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p FlashbackPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p PacketPenalty) PacketHeader() Header {
	return &p.EventData.Header
}

func (p StopGoPenaltyServedPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p DriveThroughPenaltyServedPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p StartLightsPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p RaceWinnerPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p TeamMateInPitsPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p FastestLapPacket) PacketHeader() Header {
	return &p.EventData.Header
}

func (p PacketButtons) PacketHeader() Header {
	return &p.EventData.Header
}

func (p RetirementPacket) PacketHeader() Header {
	return &p.EventData.Header
}

const (
	SessionStarted     string = "SSTA" // Sent when the session starts
	SessionEnded       string = "SEND" // Sent when the session ends
	FastestLap         string = "FTLP" // When a driver achieves the fastest lap
	Retirement         string = "RTMT" // When a driver retires
	DRSEnabled         string = "DRSE" // Race control have enabled DRS
	DRSDisabled        string = "DRSD" // Race control have disabled DRS
	TeamMateInPit      string = "TMPT" // Your team mate has entered the pits
	ChequeredFlag      string = "CHQF" // The chequered flag has been waved
	RaceWinner         string = "RCWN" // The race winner is announced
	PenaltyIssued      string = "PENA" // A penalty has been issued – details in event
	SpeedTrapTriggered string = "SPTP" // Speed trap has been triggered by fastest speed
	StartLights        string = "STLG" // Start lights – number shown
	LightsOut          string = "LGOT" // Lights out
	DriveThroughServed string = "DTSV" // Drive through penalty served
	StopGoServed       string = "SGSV" // Stop go penalty served
	Flashback          string = "FLBK" // Flashback activated
	ButtonStatus       string = "BUTN" // Button status changed
)
