package f12021

type PacketHeader struct {
	PacketFormat            uint16  `json:"m_packetFormat"`            // 2021
	GameMajorVersion        uint8   `json:"m_gameMajorVersion"`        // Game major version - "X.00"
	GameMinorVersion        uint8   `json:"m_gameMinorVersion"`        // Game minor version - "1.XX"
	PacketVersion           uint8   `json:"m_packetVersion"`           // Version of this packet type, all start from 1
	PacketId                uint8   `json:"m_packetId"`                // Identifier for the packet type, see below
	SessionId               uint64  `json:"m_sessionUID"`              // Unique identifier for the session
	SessionTime             float32 `json:"m_sessionTime"`             // Session timestamp
	FrameIdentifier         uint32  `json:"m_frameIdentifier"`         // Identifier for the frame the data was retrieved
	PlayerCarIndex          uint8   `json:"m_playerCarIndex"`          // Index of player's car in the array
	SecondaryPlayerCarIndex uint8   `json:"m_secondaryPlayerCarIndex"` // Index of secondary player's car in the array (splitscreen) 255 if no second player
}

func (p *PacketHeader) Format() uint16 {
	return p.PacketFormat
}

func (p *PacketHeader) MajorVersion() uint8 {
	return p.GameMajorVersion
}

func (p *PacketHeader) MinorVersion() uint8 {
	return p.GameMajorVersion
}

func (p *PacketHeader) PackVersion() uint8 {
	return p.PacketVersion
}

func (p *PacketHeader) Id() uint8 {
	return p.PacketId
}

func (p *PacketHeader) FrameId() uint32 {
	return p.FrameIdentifier
}

const (
	Motion              uint8 = 0  // Contains all motion data for player’s car – only sent while player is in control
	Session             uint8 = 1  // Data about the session – track, time left
	Lap                 uint8 = 2  // Data about all the lap times of cars in the session
	Event               uint8 = 3  // Various notable events that happen during a session
	Participants        uint8 = 4  // List of participants in the session, mostly relevant for multiplayer
	CarSetup            uint8 = 5  // Packet detailing car setups for cars in the race
	CarTelemetry        uint8 = 6  // Telemetry data for all cars
	CarStatus           uint8 = 7  // Status data for all cars such as damage
	FinalClassification uint8 = 8  // Final classification confirmation at the end of a race
	LobbyInfo           uint8 = 9  // Information about players in a multiplayer lobby
	CarDamage           uint8 = 10 // Damage status for all cars
	SessionHistory      uint8 = 11 // Lap and tyre data for session
)
