package f12021

type PacketLapData struct {
	Header PacketHeader `json:"m_header"` // Header

	LapData [22]LapData `json:"m_lapData"` // Lap data for all cars on track
}

func (p *PacketLapData) PacketHeader() Header {
	return &p.Header
}

type LapData struct {
	LastLapTimeInMS             uint32  `json:"m_lastLapTimeInMS"`             // Last lap time in milliseconds
	CurrentLapTimeInMS          uint32  `json:"m_currentLapTimeInMS"`          // Current time around the lap in milliseconds
	Sector1TimeInMS             uint16  `json:"m_sector1TimeInMS"`             // Sector 1 time in milliseconds
	Sector2TimeInMS             uint16  `json:"m_sector2TimeInMS"`             // Sector 2 time in milliseconds
	LapDistance                 float32 `json:"m_lapDistance"`                 // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance               float32 `json:"m_totalDistance"`               // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta              float32 `json:"m_safetyCarDelta"`              // Delta in seconds for safety car
	CarPosition                 uint8   `json:"m_carPosition"`                 // Car race position
	CurrentLapNum               uint8   `json:"m_currentLapNum"`               // Current lap number
	PitStatus                   uint8   `json:"m_pitStatus"`                   // 0 = none, 1 = pitting, 2 = in pit area
	NumPitStops                 uint8   `json:"m_numPitStops"`                 // Number of pit stops taken in this race
	Sector                      uint8   `json:"m_sector"`                      // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid           uint8   `json:"m_currentLapInvalid"`           // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                   uint8   `json:"m_penalties"`                   // Accumulated time penalties in seconds to be added
	Warnings                    uint8   `json:"m_warnings"`                    // Accumulated number of warnings issued
	NumUnservedDriveThroughPens uint8   `json:"m_numUnservedDriveThroughPens"` // Num drive through pens left to serve
	NumUnservedStopGoPens       uint8   `json:"m_numUnservedStopGoPens"`       // Num stop go pens left to serve
	GridPosition                uint8   `json:"m_gridPosition"`                // Grid position the vehicle started the race in
	DriverStatus                uint8   `json:"m_driverStatus"`                // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus                uint8   `json:"m_resultStatus"`                // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
	PitLaneTimerActive          uint8   `json:"m_pitLaneTimerActive"`          // Pit lane timing, 0 = inactive, 1 = active
	PitLaneTimeInLaneInMS       uint16  `json:"m_pitLaneTimeInLaneInMS"`       // If active, the current time spent in the pit lane in ms
	PitStopTimerInMS            uint16  `json:"m_pitStopTimerInMS"`            // Time of the actual pit stop in ms
	PitStopShouldServePen       uint8   `json:"m_pitStopShouldServePen"`       // Whether the car should serve a penalty at this stop
}
