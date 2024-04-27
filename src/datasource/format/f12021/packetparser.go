package f12021

import (
	"bytes"
	"encoding/binary"
	"fmt"
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
		break
	case Session:
		p = new(PacketSessionData)
		break
	case Lap:
		p = new(PacketLapData)
		break
	case Event:
		p = selectEvent(buf)
		break
	case Participants:
		p = new(PacketParticipantsData)
		break
	case CarSetup:
		p = new(PacketCarSetupData)
		break
	case CarTelemetry:
		p = new(PacketCarTelemetryData)
		break
	case CarStatus:
		p = new(PacketCarSetupData)
		break
	case FinalClassification:
		p = new(PacketFinalClassificationData)
		break
	case LobbyInfo:
		p = new(PacketLobbyInfoData)
		break
	case CarDamage:
		p = new(PacketCarDamageData)
		break
	case SessionHistory:
		p = new(PacketSessionHistoryData)
		break
	default:
		p = new(NullPacket)
		break
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
		break
	case SessionEnded:
		resultPacket = new(PacketSessionData)
		break
	case FastestLap:
		resultPacket = new(FastestLapPacket)
		break
	case Retirement:
		resultPacket = new(RetirementPacket)
		break
	case DRSEnabled:
		resultPacket = new(PacketEventData)
		break
	case DRSDisabled:
		resultPacket = new(PacketEventData)
		break
	case TeamMateInPit:
		resultPacket = new(TeamMateInPitsPacket)
		break
	case ChequeredFlag:
		resultPacket = new(PacketEventData)
		break
	case RaceWinner:
		resultPacket = new(RaceWinnerPacket)
		break
	case PenaltyIssued:
		resultPacket = new(PacketPenalty)
		break
	case SpeedTrapTriggered:
		resultPacket = new(SpeedTrapPacket)
		break
	case StartLights:
		resultPacket = new(StartLightsPacket)
		break
	case LightsOut:
		resultPacket = new(StartLightsPacket)
		break
	case DriveThroughServed:
		resultPacket = new(DriveThroughPenaltyServedPacket)
		break
	case StopGoServed:
		resultPacket = new(StopGoPenaltyServedPacket)
		break
	case Flashback:
		resultPacket = new(FlashbackPacket)
		break
	case ButtonStatus:
		resultPacket = new(PacketButtons)
		break
	default:
		fmt.Println("Unknown Event >> " + p.EventString())
	}

	return resultPacket
}

func fillPacket(buf []byte, packet any) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, packet); err != nil {
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
	MajorVersion() uint8
	MinorVersion() uint8
	PackVersion() uint8
}
