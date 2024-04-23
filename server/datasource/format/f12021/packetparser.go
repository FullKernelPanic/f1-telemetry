package f12021

import (
	"bytes"
	"encoding/binary"
)

func readheader(buf []byte) Header {
	h := PacketHeader{}

	fillPacket(buf, &h)

	return &h
}

func readpacket(packetId uint8, buf []byte) Packet {
	var p Packet

	switch packetId {
	case Motion:
		p = new(PacketMotionData)
	case Session:
		p = new(PacketSessionData)
	case Lap:
		p = new(PacketLapData)
	case Event:
		p = selectEvent(buf)
	case Participants:
		p = new(PacketParticipantsData)
	case CarSetup:
		p = new(PacketCarSetupData)
	case CarTelemetry:
		p = new(PacketCarTelemetryData)
	case CarStatus:
		p = new(PacketCarSetupData)
	case FinalClassification:
		p = new(PacketFinalClassificationData)
	case LobbyInfo:
		p = new(PacketLobbyInfoData)
	case CarDamage:
		p = new(PacketCarDamageData)
	case SessionHistory:
		p = new(PacketSessionHistoryData)
	default:
		p = new(NullPacket)
	}

	fillPacket(buf, p)

	return p
}

func selectEvent(buf []byte) Packet {
	p := new(PacketEventData)

	var resultPacket Packet

	fillPacket(buf, p)

	switch p.EventString() {
	case SessionStarted:
		resultPacket = new(PacketSessionData)
	case SessionEnded:
		resultPacket = new(PacketSessionData)
	case FastestLap:
		resultPacket = new(FastestLapPacket)
	case Retirement:
		resultPacket = new(RetirementPacket)
	case DRSEnabled:
		resultPacket = new(PacketEventData)
	case DRSDisabled:
		resultPacket = new(PacketEventData)
	case TeamMateInPit:
		resultPacket = new(TeamMateInPitsPacket)
	case ChequeredFlag:
		resultPacket = new(PacketEventData)
	case RaceWinner:
		resultPacket = new(RaceWinnerPacket)
	case PenaltyIssued:
		resultPacket = new(PacketPenalty)
	case SpeedTrapTriggered:
		resultPacket = new(SpeedTrapPacket)
	case StartLights:
		resultPacket = new(StartLightsPacket)
	case LightsOut:
		resultPacket = new(StartLightsPacket)
	case DriveThroughServed:
		resultPacket = new(DriveThroughPenaltyServedPacket)
	case StopGoServed:
		resultPacket = new(StopGoPenaltyServedPacket)
	case Flashback:
		resultPacket = new(FlashbackPacket)
	case ButtonStatus:
		resultPacket = new(PacketButtons)
	}

	return resultPacket
}

func fillPacket(buf []byte, pack any) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

type PacketMapper interface {
	getHeader() Header
	getPacket() Packet
}

type Packet interface {
	PacketHeader() Header
}

type Header interface {
	Format() uint16
	Id() uint8
	FrameId() uint32
}
