package versions

import "f1telemetry/src/domain"

type Packet interface {
	PacketHeader() Header
}

type Header interface {
	Format() uint16
	Id() uint8
	FrameId() uint32
	MajorVersion() uint8
	MinorVersion() uint8
	PackVersion() uint8
	GetPlayerCarIndex() uint8
}

type MessageEnvelopeHeader struct {
	Type domain.MessageType
}

type MessageEnvelope struct {
	Type domain.MessageType
	Data interface{}
}

func (e *MessageEnvelope) GetType() domain.MessageType {
	return e.Type
}

type PacketType string

type PacketEnvelop struct {
	EventType PacketType `json:"event_type"`
	Packet    Packet     `json:"packet"`
}
