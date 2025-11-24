package f1_2021

import (
	"f1telemetry/src/telemetry/versions"
	"f1telemetry/src/util"
	"fmt"
)

const (
	F12021_MOTION               versions.PacketType = "F12021_MOTION"
	F12021_SESSION              versions.PacketType = "F12021_SESSION"
	F12021_LAP                  versions.PacketType = "F12021_LAP"
	F12021_PARTICIPANTS         versions.PacketType = "F12021_PARTICIPANTS"
	F12021_CARSETUP             versions.PacketType = "F12021_CARSETUP"
	F12021_CARTELEMETRY         versions.PacketType = "F12021_CARTELEMETRY"
	F12021_CARSTATUS            versions.PacketType = "F12021_CARSTATUS"
	F12021_FINALCLASSIFICATION  versions.PacketType = "F12021_FINALCLASSIFICATION"
	F12021_LOBBYINFO            versions.PacketType = "F12021_LOBBYINFO"
	F12021_CARDAMAGE            versions.PacketType = "F12021_CARDAMAGE"
	F12021_SESSIONHISTORY       versions.PacketType = "F12021_SESSIONHISTORY"
	F12021_SESSION_STARTED      versions.PacketType = "F12021_SESSION_STARTED"
	F12021_SESSION_ENDED        versions.PacketType = "F12021_SESSION_ENDED"
	F12021_FASTEST_LAP          versions.PacketType = "F12021_FASTEST_LAP"
	F12021_RETIREMENT           versions.PacketType = "F12021_RETIREMENT"
	F12021_DRS_ENABLED          versions.PacketType = "F12021_DRS_ENABLED"
	F12021_DRS_DISABLED         versions.PacketType = "F12021_DRS_DISABLED"
	F12021_TEAMMATE_IN_PIT      versions.PacketType = "F12021_TEAMMATE_IN_PIT"
	F12021_CHEQUERED_FLAG       versions.PacketType = "F12021_CHEQUERED_FLAG"
	F12021_RACE_WINNER          versions.PacketType = "F12021_RACE_WINNER"
	F12021_PENALTY_ISSUED       versions.PacketType = "F12021_PENALTY_ISSUED"
	F12021_SPEED_TRAP_TRIGGERED versions.PacketType = "F12021_SPEED_TRAP_TRIGGERED"
	F12021_START_LIGHTS         versions.PacketType = "F12021_START_LIGHTS"
	F12021_LIGHTS_OUT           versions.PacketType = "F12021_LIGHTS_OUT"
	F12021_DRIVE_THROUGH_SERVED versions.PacketType = "F12021_DRIVE_THROUGH_SERVED"
	F12021_STOP_GO_SERVED       versions.PacketType = "F12021_STOP_GO_SERVED"
	F12021_FLASHBACK            versions.PacketType = "F12021_FLASHBACK"
	F12021_BUTTON_STATUS        versions.PacketType = "F12021_BUTTON_STATUS"
)

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

func init() {
	fmt.Println("F1 2021 telemetry version loaded")
}

func MapBytes(data []byte) *versions.PacketEnvelop {
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

func readPacket(ph PacketHeader, data []byte) (*versions.PacketEnvelop, error) {
	var p versions.Packet
	var err error
	var t versions.PacketType
	switch ph.PacketId {
	case Motion:
		t = F12021_MOTION
		p = new(PacketMotionData)
		break
	case Session:
		t = F12021_SESSION
		p = new(PacketSessionData)
		break
	case Lap:
		t = F12021_LAP
		p = new(PacketLapData)
		break
	case Event:
		p, t, err = selectEvent(data)
		if err != nil {
			return nil, err
		}
		break
	case Participants:
		t = F12021_PARTICIPANTS
		p = new(PacketParticipantsData)
		break
	case CarSetup:
		t = F12021_CARSETUP
		p = new(PacketCarSetupData)
		break
	case CarTelemetry:
		t = F12021_CARTELEMETRY
		p = new(PacketCarTelemetryData)
		break
	//case CarStatus:
	//	t = F12021_CARSTATUS
	//  p = new(PacketCarStatus)
	//break
	case FinalClassification:
		t = F12021_FINALCLASSIFICATION
		p = new(PacketFinalClassificationData)
		break
	case LobbyInfo:
		t = F12021_LOBBYINFO
		p = new(PacketLobbyInfoData)
		break
	case CarDamage:
		t = F12021_CARDAMAGE
		p = new(PacketCarDamageData)
		break
	case SessionHistory:
		t = F12021_SESSIONHISTORY
		p = new(PacketSessionHistoryData)
		break
	default:
		return nil, fmt.Errorf("unsupported packet id: %d", ph.PacketId)
	}

	err = util.MapBytesToPacket(data, p)

	if err != nil {
		return nil, err
	}

	return &versions.PacketEnvelop{
		EventType: t,
		Packet:    p,
	}, nil
}

func selectEvent(buf []byte) (versions.Packet, versions.PacketType, error) {
	p := new(PacketEventData)

	var resultPacket versions.Packet

	err := util.MapBytesToPacket(buf, p)
	if err != nil {
		return nil, "", err
	}
	var t versions.PacketType
	switch p.EventString() {
	case SessionStarted:
		t = F12021_SESSION_STARTED
		resultPacket = new(PacketSessionData)
		break
	case SessionEnded:
		t = F12021_SESSION_ENDED
		resultPacket = new(PacketSessionData)
		break
	case FastestLap:
		t = F12021_FASTEST_LAP
		resultPacket = new(FastestLapPacket)
		break
	case Retirement:
		t = F12021_RETIREMENT
		resultPacket = new(RetirementPacket)
		break
	case DRSEnabled:
		t = F12021_DRS_ENABLED
		resultPacket = new(PacketEventData)
		break
	case DRSDisabled:
		t = F12021_DRS_DISABLED
		resultPacket = new(PacketEventData)
		break
	case TeamMateInPit:
		t = F12021_TEAMMATE_IN_PIT
		resultPacket = new(TeamMateInPitsPacket)
		break
	case ChequeredFlag:
		t = F12021_CHEQUERED_FLAG
		resultPacket = new(PacketEventData)
		break
	case RaceWinner:
		t = F12021_RACE_WINNER
		resultPacket = new(RaceWinnerPacket)
		break
	case PenaltyIssued:
		t = F12021_PENALTY_ISSUED
		resultPacket = new(PacketPenalty)
		break
	//case SpeedTrapTriggered:
	//	t = F12021_SPEED_TRAP_TRIGGERED
	//	resultPacket = new(SpeedTrapPacket)
	//	break
	case StartLights:
		t = F12021_START_LIGHTS
		resultPacket = new(StartLightsPacket)
		break
	case LightsOut:
		t = F12021_LIGHTS_OUT
		resultPacket = new(StartLightsPacket)
		break
	case DriveThroughServed:
		t = F12021_DRIVE_THROUGH_SERVED
		resultPacket = new(DriveThroughPenaltyServedPacket)
		break
	case StopGoServed:
		t = F12021_STOP_GO_SERVED
		resultPacket = new(StopGoPenaltyServedPacket)
		break
	case Flashback:
		t = F12021_FLASHBACK
		resultPacket = new(FlashbackPacket)
		break
	case ButtonStatus:
		t = F12021_BUTTON_STATUS
		resultPacket = new(PacketButtons)
		break
	default:
		return nil, "", fmt.Errorf("unknown event type: %s", p.EventString())
	}

	return resultPacket, t, nil
}
