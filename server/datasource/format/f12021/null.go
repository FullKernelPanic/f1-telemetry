package f12021

type NullPacket struct {
	Header PacketHeader
}

func (p *NullPacket) PacketHeader() Header {
	return &p.Header
}
