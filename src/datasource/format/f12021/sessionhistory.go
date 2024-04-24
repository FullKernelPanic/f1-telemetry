package f12021

type PacketSessionHistoryData struct {
	Header               PacketHeader            `json:"m_header"`            // Header
	CarIdx               uint8                   `json:"m_carIdx"`            // Index of the car this lap data relates to
	NumLaps              uint8                   `json:"m_numLaps"`           // Num laps in the data (including current partial lap)
	NumTyreStints        uint8                   `json:"m_numTyreStints"`     // Number of tyre stints in the data
	BestLapTimeLapNum    uint8                   `json:"m_bestLapTimeLapNum"` // Lap the best lap time was achieved on
	BestSector1LapNum    uint8                   `json:"m_bestSector1LapNum"` // Lap the best Sector 1 time was achieved on
	BestSector2LapNum    uint8                   `json:"m_bestSector2LapNum"` // Lap the best Sector 2 time was achieved on
	BestSector3LapNum    uint8                   `json:"m_bestSector3LapNum"` // Lap the best Sector 3 time was achieved on
	LapHistoryData       [100]LapHistoryData     `json:"m_lapHistoryData"`    // 100 laps of data max
	TyreStintHistoryData [8]TyreStintHistoryData `json:"m_tyreStintsHistoryData"`
}

func (p *PacketSessionHistoryData) PacketHeader() Header {
	return &p.Header
}

type LapHistoryData struct {
	LapTimeInMS      uint32 `json:"m_lapTimeInMS"`      // Lap time in milliseconds
	Sector1TimeInMS  uint16 `json:"m_sector1TimeInMS"`  // Sector 1 time in milliseconds
	Sector2TimeInMS  uint16 `json:"m_sector2TimeInMS"`  // Sector 2 time in milliseconds
	Sector3TimeInMS  uint16 `json:"m_sector3TimeInMS"`  // Sector 3 time in milliseconds
	LapValidBitFlags uint8  `json:"m_lapValidBitFlags"` // 0x01 bit set-lap valid,      0x02 bit set-sector 1 valid 0x04 bit set-sector 2 valid, 0x08 bit set-sector 3 valid
}

type TyreStintHistoryData struct {
	EndLap             uint8 `json:"m_endLap"`             // Lap the tyre usage ends on (255 of current tyre)
	TyreActualCompound uint8 `json:"m_tyreActualCompound"` // Actual tyres used by this driver
	TyreVisualCompound uint8 `json:"m_tyreVisualCompound"` // Visual tyres used by this driver
}
