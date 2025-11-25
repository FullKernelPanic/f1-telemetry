package f1_2024

import (
	"f1telemetry/src/telemetry/common"
	"f1telemetry/src/util"
	"fmt"
)

const (
	F12024_MOTION               common.PacketType = "F12024_MOTION"
	F12024_SESSION              common.PacketType = "F12024_SESSION"
	F12024_LAP                  common.PacketType = "F12024_LAP"
	F12024_PARTICIPANTS         common.PacketType = "F12024_PARTICIPANTS"
	F12024_CARSETUP             common.PacketType = "F12024_CARSETUP"
	F12024_CARTELEMETRY         common.PacketType = "F12024_CARTELEMETRY"
	F12024_CARSTATUS            common.PacketType = "F12024_CARSTATUS"
	F12024_FINALCLASSIFICATION  common.PacketType = "F12024_FINALCLASSIFICATION"
	F12024_LOBBYINFO            common.PacketType = "F12024_LOBBYINFO"
	F12024_CARDAMAGE            common.PacketType = "F12024_CARDAMAGE"
	F12024_SESSIONHISTORY       common.PacketType = "F12024_SESSIONHISTORY"
	F12024_SESSION_STARTED      common.PacketType = "F12024_SESSION_STARTED"
	F12024_SESSION_ENDED        common.PacketType = "F12024_SESSION_ENDED"
	F12024_FASTEST_LAP          common.PacketType = "F12024_FASTEST_LAP"
	F12024_RETIREMENT           common.PacketType = "F12024_RETIREMENT"
	F12024_DRS_ENABLED          common.PacketType = "F12024_DRS_ENABLED"
	F12024_DRS_DISABLED         common.PacketType = "F12024_DRS_DISABLED"
	F12024_TEAMMATE_IN_PIT      common.PacketType = "F12024_TEAMMATE_IN_PIT"
	F12024_CHEQUERED_FLAG       common.PacketType = "F12024_CHEQUERED_FLAG"
	F12024_RACE_WINNER          common.PacketType = "F12024_RACE_WINNER"
	F12024_PENALTY_ISSUED       common.PacketType = "F12024_PENALTY_ISSUED"
	F12024_SPEED_TRAP_TRIGGERED common.PacketType = "F12024_SPEED_TRAP_TRIGGERED"
	F12024_START_LIGHTS         common.PacketType = "F12024_START_LIGHTS"
	F12024_LIGHTS_OUT           common.PacketType = "F12024_LIGHTS_OUT"
	F12024_DRIVE_THROUGH_SERVED common.PacketType = "F12024_DRIVE_THROUGH_SERVED"
	F12024_STOP_GO_SERVED       common.PacketType = "F12024_STOP_GO_SERVED"
	F12024_FLASHBACK            common.PacketType = "F12024_FLASHBACK"
	F12024_BUTTON_STATUS        common.PacketType = "F12024_BUTTON_STATUS"
)

const (
	Motion              uint8 = 0  // Contains all motion data for player’s car – only sent while player is in control
	Session             uint8 = 1  // Data about the session – track, time left
	Lap                 uint8 = 2  // Data about all the lap times of cars in the session
	Event               uint8 = 3  // Various notable events that happen during a session
	Participants        uint8 = 4  // List of participants in the session, mostly relevant for multiplayer
	CarSetups           uint8 = 5  // Packet detailing car setups for cars in the race
	CarTelemetry        uint8 = 6  // Telemetry data for all cars
	CarStatus           uint8 = 7  // Status data for all cars such as damage
	FinalClassification uint8 = 8  // Final classification confirmation at the end of a race
	LobbyInfo           uint8 = 9  // Information about players in a multiplayer lobby
	CarDamage           uint8 = 10 // Damage status for all cars
	SessionHistory      uint8 = 11 // Lap and tyre data for session
	TyreSets            uint8 = 12 // Extended tyre set data
	MotionEx            uint8 = 13 // Extended motion data for player car
	TimeTrial           uint8 = 14 // Time trial specificdata
)
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
	RedFlag            string = "RDFL" // Red flag shown
	Overtake           string = "OVTK" // Overtake occurred
	SafetyCar          string = "SCAR" // Safety car event - details in event
	Collision          string = "COLL" // Collision between two vehicles has occurred
)

func init() {
	fmt.Println("F1 2024 telemetry version loaded")
}

func MapBytes(data []byte) *common.PacketEnvelop {
	ph := PacketHeader{}
	_ = util.MapBytesToPacket(data, &ph)
	p, err := readPacket(ph, data)

	if err != nil {
		fmt.Println(err)
		return nil
	}
	fmt.Println("Packet header: ", p.EventType)

	return p
}

func readPacket(ph PacketHeader, data []byte) (*common.PacketEnvelop, error) {
	var p common.Packet
	var err error
	var t common.PacketType
	fmt.Println(ph.PacketId)
	switch ph.PacketId {
	case Motion:
		t = F12024_MOTION
		p = new(PacketMotionData)
		break
	case Session:
		t = F12024_SESSION
		p = new(PacketSessionData)
		break
	case Lap:
		t = F12024_LAP
		p = new(PacketLapData)
		break
	case Event:
		p, t, err = selectEvent(data)
		if err != nil {
			return nil, err
		}
		break
	case Participants:
		t = F12024_PARTICIPANTS
		p = new(PacketParticipantsData)
		break
	case CarSetups:
		t = F12024_CARSETUP
		p = new(PacketCarSetupData)
		break
	case CarTelemetry:
		t = F12024_CARTELEMETRY
		p = new(PacketCarTelemetryData)
		break
	//case CarStatus:
	//	t = F12024_CARSTATUS
	//  p = new(PacketCarStatus)
	//break
	case FinalClassification:
		t = F12024_FINALCLASSIFICATION
		p = new(PacketFinalClassificationData)
		break
	case LobbyInfo:
		t = F12024_LOBBYINFO
		p = new(PacketLobbyInfoData)
		break
	case CarDamage:
		t = F12024_CARDAMAGE
		p = new(PacketCarDamageData)
		break
	case SessionHistory:
		t = F12024_SESSIONHISTORY
		p = new(PacketSessionHistoryData)
		break
	default:
		return nil, fmt.Errorf("unsupported packet id: %d", ph.PacketId)
	}

	err = util.MapBytesToPacket(data, p)

	if err != nil {
		return nil, err
	}

	return &common.PacketEnvelop{
		EventType: t,
		Packet:    p,
	}, nil
}

func selectEvent(buf []byte) (common.Packet, common.PacketType, error) {
	p := new(PacketEventData)

	var resultPacket common.Packet

	err := util.MapBytesToPacket(buf, p)
	if err != nil {
		return nil, "", err
	}
	var t common.PacketType
	switch p.EventString() {
	case SessionStarted:
		t = F12024_SESSION_STARTED
		resultPacket = new(PacketSessionData)
		break
	case SessionEnded:
		t = F12024_SESSION_ENDED
		resultPacket = new(PacketSessionData)
		break
	case FastestLap:
		t = F12024_FASTEST_LAP
		resultPacket = new(FastestLapPacket)
		break
	case Retirement:
		t = F12024_RETIREMENT
		resultPacket = new(RetirementPacket)
		break
	case DRSEnabled:
		t = F12024_DRS_ENABLED
		resultPacket = new(PacketEventData)
		break
	case DRSDisabled:
		t = F12024_DRS_DISABLED
		resultPacket = new(PacketEventData)
		break
	case TeamMateInPit:
		t = F12024_TEAMMATE_IN_PIT
		resultPacket = new(TeamMateInPitsPacket)
		break
	case ChequeredFlag:
		t = F12024_CHEQUERED_FLAG
		resultPacket = new(PacketEventData)
		break
	case RaceWinner:
		t = F12024_RACE_WINNER
		resultPacket = new(RaceWinnerPacket)
		break
	case PenaltyIssued:
		t = F12024_PENALTY_ISSUED
		resultPacket = new(PacketPenalty)
		break
	//case SpeedTrapTriggered:
	//	t = F12024_SPEED_TRAP_TRIGGERED
	//	resultPacket = new(SpeedTrapPacket)
	//	break
	case StartLights:
		t = F12024_START_LIGHTS
		resultPacket = new(StartLightsPacket)
		break
	case LightsOut:
		t = F12024_LIGHTS_OUT
		resultPacket = new(StartLightsPacket)
		break
	case DriveThroughServed:
		t = F12024_DRIVE_THROUGH_SERVED
		resultPacket = new(DriveThroughPenaltyServedPacket)
		break
	case StopGoServed:
		t = F12024_STOP_GO_SERVED
		resultPacket = new(StopGoPenaltyServedPacket)
		break
	case Flashback:
		t = F12024_FLASHBACK
		resultPacket = new(FlashbackPacket)
		break
	case ButtonStatus:
		t = F12024_BUTTON_STATUS
		resultPacket = new(PacketButtons)
		break
	default:
		return nil, "", fmt.Errorf("unknown event type: %s", p.EventString())
	}

	return resultPacket, t, nil
}
