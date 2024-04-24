package f12021

type PacketLobbyInfoData struct {
	Header PacketHeader `json:"m_header"` // Header

	// Packet specific data
	NumPlayers   uint8             `json:"m_numPlayers"` // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData `json:"m_lobbyPlayers"`
}

func (p *PacketLobbyInfoData) PacketHeader() Header {
	return &p.Header
}

type LobbyInfoData struct {
	AIControlled uint8      `json:"m_aiControlled"` // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId       uint8      `json:"m_teamId"`       // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8      `json:"m_nationality"`  // Nationality of the driver
	Name         [48]string `json:"m_name"`         // Name of participant in UTF-8 format – null terminated Will be truncated with ... (U+2026) if too long
	CarNumber    uint8      `json:"m_carNumber"`    // Car number of the player
	ReadyStatus  uint8      `json:"m_readyStatus"`  // 0 = not ready, 1 = ready, 2 = spectating
}
