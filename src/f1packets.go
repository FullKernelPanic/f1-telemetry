package main

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
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

// MOTION
type PacketMotionData struct {
	Header        PacketHeader      `json:"m_header"`
	CarMotionData [22]CarMotionData `json:"m_carMotionData"`

	// RL, RR, FL, FR
	SuspensionPosition     [4]float32 `json:"m_suspensionPosition"`
	SuspensionVelocity     [4]float32 `json:"m_suspensionVelocity"`
	SuspensionAcceleration [4]float32 `json:"m_suspensionAcceleration"`
	WheelSpeed             [4]float32 `json:"m_wheelSpeed"`
	WheelSlip              [4]float32 `json:"m_wheelSlip"`

	LocalVelocityX float32 `json:"m_localVelocityX"`
	LocalVelocityY float32 `json:"m_localVelocityY"`
	LocalVelocityZ float32 `json:"m_localVelocityZ"`

	AngularVelocityX float32 `json:"m_angularVelocityX"`
	AngularVelocityY float32 `json:"m_angularVelocityY"`
	AngularVelocityZ float32 `json:"m_angularVelocityZ"`

	AngularAccelerationX float32 `json:"m_angularAccelerationX"`
	AngularAccelerationY float32 `json:"m_angularAccelerationY"`
	AngularAccelerationZ float32 `json:"m_angularAccelerationZ"`

	FrontWheelsAngle float32 `json:"m_frontWheelsAngle"`
}

// SESSION
type PacketSessionData struct {
	Header PacketHeader `json:"m_header"`

	Weather          uint8 `json:"m_weather"` // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature int8  `json:"m_trackTemperature"`
	AirTemperature   int8  `json:"m_airTemperature"`

	TotalLaps   int8   `json:"m_totalLaps"`
	TrackLength uint16 `json:"m_trackLength"`
	SessionType uint16 `json:"m_sessionType"` // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2, 12 = R3, 13 = Time Trial
	TrackId     int8   `json:"m_trackId"`     // -1 for unknown, 0-21 for tracks, see appendix
	Formula     uint8  `json:"m_formula"`     // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, 3 = F1 Generic

	SessionTimeLeft      uint16          `json:"m_sessionTimeLeft"`
	SessionDuration      uint16          `json:"m_sessionDuration"`
	PitSpeedLimit        uint8           `json:"m_pitSpeedLimit"`
	GamePaused           uint8           `json:"m_gamePaused"`
	IsSpectating         uint8           `json:"m_isSpectating"`
	SpectatorCarIndex    uint8           `json:"m_spectatorCarIndex"`
	SLIProNativeSupport  uint8           `json:"m_sliProNativeSupport"`
	NumberOfMarshalZones uint8           `json:"m_numMarshalZones"`
	MarshalZones         [21]MarshalZone `json:"m_marshalZones"`
	SafetyCarStatus      uint8           `json:"m_safetyCarStatus"` // 0 = no safety car, 1 = full, 2 = virtual, 3 = formation lap
	NetworkGame          uint8           `json:"m_networkGame"`     // 0 = offline, 1 = online

	NumWeatherForecastSamples uint8                     `json:"m_numWeatherForecastSamples"`
	WeatherForecastSamples    [55]WeatherForecastSample `json:"m_weatherForecastSamples"`
	ForecastAccuracy          uint8                     `json:"m_forecastAccuracy"`

	AIDifficulty uint8 `json:"m_aiDifficulty"`

	SeasonLinkIdentifier   uint32 `json:"m_seasonLinkIdentifier"`   // Identifier for season - persists across saves
	WeekendLinkIdentifier  uint32 `json:"m_weekendLinkIdentifier"`  // Identifier for weekend - persists across saves
	SessionLinkIdentifier  uint32 `json:"m_sessionLinkIdentifier"`  // Identifier for session - persists across saves
	PitStopWindowIdealLap  uint8  `json:"m_pitStopWindowIdealLap"`  // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap uint8  `json:"m_pitStopWindowLatestLap"` // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition  uint8  `json:"m_pitStopRejoinPosition"`  // Predicted position to rejoin at (player)
	SteeringAssist         uint8  `json:"m_steeringAssist"`         // 0 = off, 1 = on
	BrakingAssist          uint8  `json:"m_brakingAssist"`          // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist          uint8  `json:"m_gearboxAssist"`          // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist              uint8  `json:"m_pitAssist"`              // 0 = off, 1 = on
	PitReleaseAssist       uint8  `json:"m_pitReleaseAssist"`       // 0 = off, 1 = on
	ERSAssist              uint8  `json:"m_ERSAssist"`              // 0 = off, 1 = on
	DRSAssist              uint8  `json:"m_DRSAssist"`              // 0 = off, 1 = on
	DynamicRacingLine      uint8  `json:"m_dynamicRacingLine"`      // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType  uint8  `json:"m_dynamicRacingLineType"`  // 0 = 2D, 1 = 3D
}

// LAPDATA
type PacketLapData struct {
	Header PacketHeader `json:"m_header"` // Header

	LapData [22]LapData `json:"m_lapData"` // Lap data for all cars on track
}

// EVENT
type PacketEventData struct {
	Header          PacketHeader `json:"m_packetHeader"`    // Header
	EventStringCode [4]uint8     `json:"m_eventStringCode"` // Event string code, see below
}

// PARTICIPANTS
type PacketParticipantsData struct {
	Header          PacketHeader        `json:"m_header"`        // Header
	NumActiveCars   uint8               `json:"m_numActiveCars"` // Number of active cars in the data – should match number of cars on HUD
	ParticipantData [22]ParticipantData `json:"m_participants"`
}

// CARSETUP
type PacketCarSetupData struct {
	Header       PacketHeader     `json:"m_header"` // Header
	CarSetupData [22]CarSetupData `json:"m_carSetups"`
}

// CARTELEMETRY
type PacketCarTelemetryData struct {
	Header                       PacketHeader         `json:"m_header"` // Header
	CarTelemetryData             [22]CarTelemetryData `json:"m_carTelemetryData"`
	MFDPanelIndex                uint8                `json:"m_mfdPanelIndex"`                // Index of MFD panel open - 255 = MFD closed, Single player, race – 0 = Car setup, 1 = Pits, 2 = Damage, 3 =  Engine, 4 = Temperatures, May vary depending on game mode
	MFDPanelIndexSecondaryPlayer uint8                `json:"m_mfdPanelIndexSecondaryPlayer"` // See above
	SuggestedGear                int8                 `json:"m_suggestedGear"`                // Suggested gear for the player (1-8), 0 if no gear suggested
}

//CARSTATUS

// FINALCLASSIFICATION
type PacketFinalClassificationData struct {
	Header                  PacketHeader                `json:"m_header"`  // Header
	NumCars                 uint8                       `json:"m_numCars"` // Number of cars in the final classification
	FinalClassificationData [22]FinalClassificationData `json:"m_classificationData"`
}

// LOBBYINFO
type PacketLobbyInfoData struct {
	Header PacketHeader `json:"m_header"` // Header

	// Packet specific data
	NumPlayers   uint8             `json:"m_numPlayers"` // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData `json:"m_lobbyPlayers"`
}

// CARDAMAGE
type PacketCarDamageData struct {
	Header PacketHeader `json:"m_header"` // Header

	CarDamageData [22]CarDamageData `json:"m_carDamageData"`
}

// SESSIONHISTORY
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

type CarMotionData struct {
	WorldPositionX float32 `json:"m_worldPositionX"`
	WorldPositionY float32 `json:"m_worldPositionY"`
	WorldPositionZ float32 `json:"m_worldPositionZ"`

	WorldVelocityX float32 `json:"m_worldVelocityX"`
	WorldVelocityY float32 `json:"m_worldVelocityY"`
	WorldVelocityZ float32 `json:"m_worldVelocityZ"`

	WorldForwardDirX int16 `json:"m_worldForwardDirX"`
	WorldForwardDirY int16 `json:"m_worldForwardDirY"`
	WorldForwardDirZ int16 `json:"m_worldForwardDirZ"`

	WorldRightDirX int16 `json:"m_worldRightDirX"`
	WorldRightDirY int16 `json:"m_worldRightDirY"`
	WorldRightDirZ int16 `json:"m_worldRightDirZ"`

	GForceLateral      float32 `json:"m_gForceLateral"`
	GForceLongitudinal float32 `json:"m_gForceLongitudinal"`
	GForceVertical     float32 `json:"m_gForceVertical"`

	Yaw   float32 `json:"m_yaw"`
	Pitch float32 `json:"m_pitch"`
	Roll  float32 `json:"m_roll"`
}

type MarshalZone struct {
	ZoneStart float32 `json:"m_zoneStart"` // Fraction (0..1) of way through the lap the marshal zone starts
	ZoneFlag  int8    `json:"m_zoneFlag"`  // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
}

type WeatherForecastSample struct {
	SessionType            uint8 `json:"m_sessionType"`            // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2 12 = Time Trial
	TimeOffset             uint8 `json:"m_timeOffset"`             // Time in minutes the forecast is for
	Weather                uint8 `json:"m_weather"`                // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature       int8  `json:"m_trackTemperature"`       // Track temp. in degrees Celsius
	TrackTemperatureChange int8  `json:"m_trackTemperatureChange"` // Track temp. change – 0 = up, 1 = down, 2 = no change
	AirTemperature         int8  `json:"m_airTemperature"`         // Air temp. in degrees celsius
	AirTemperatureChange   int8  `json:"m_airTemperatureChange"`   // Air temp. change – 0 = up, 1 = down, 2 = no change
	RainPercentage         uint8 `json:"m_rainPercentage"`         // Rain percentage (0-100)
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

type FastestLapPacket struct {
	EventData PacketEventData `json:"m_packetEventData"`

	VehicleIdx uint8   `json:"vehicleIdx"` // Vehicle index of car achieving fastest lap
	LapTime    float32 `json:"lapTime"`    // Lap time is in seconds
}

type RetirementPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of car retiring
}

type TeamMateInPitsPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of team mate
}

type RaceWinnerPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the race winner
}

type PenaltyPacket struct {
	EventData        PacketEventData `json:"m_packetEventData"`
	PenaltyType      uint8           `json:"penaltyType"`      // Penalty type – see Appendices
	InfringementType uint8           `json:"infringementType"` // Infringement type – see Appendices
	VehicleIndex     uint8           `json:"vehicleIdx"`       // Vehicle index of the car the penalty is applied to
	OtherVehicleIdx  uint8           `json:"otherVehicleIdx"`  // Vehicle index of the other car involved
	Time             uint8           `json:"time"`             // Time gained, or time spent doing action in seconds
	LapNum           uint8           `json:"lapNum"`           // Lap the penalty occurred on
	PlacesGained     uint8           `json:"placesGained"`     // Number of places gained by this
}

type SpeedTrapPacket struct {
	EventData               PacketEventData `json:"m_packetEventData"`
	VehicleIndex            uint8           `json:"vehicleIdx"`              // Vehicle index of the vehicle triggering speed trap
	Speed                   float32         `json:"speed"`                   // Top speed achieved in kilometres per hour
	OverallFastestInSession uint8           `json:"overallFastestInSession"` // Overall fastest speed in session = 1, otherwise 0
	DriverFastestInSession  uint8           `json:"driverFastestInSession"`  // Fastest speed for driver in session = 1, otherwise 0
}

type StartLightsPacket struct {
	EventData PacketEventData `json:"m_packetEventData"`
	NumLights uint8           `json:"numLights"` // Number of lights showing
}

type DriveThroughPenaltyServedPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the vehicle serving drive through
}

type StopGoPenaltyServedPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	VehicleIndex uint8           `json:"vehicleIdx"` // Vehicle index of the vehicle serving stop go
}

type FlashbackPacket struct {
	EventData                PacketEventData `json:"m_packetEventData"`
	FlashbackFrameIdentifier uint32          `json:"flashbackFrameIdentifier"` // Frame identifier flashed back to
	FlashbackSessionTime     float32         `json:"flashbackSessionTime"`     // Session time flashed back to
}

type ButtonsPacket struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	ButtonStatus uint32          `json:"m_buttonStatus"` // Bit flags specifying which buttons are being pressed currently - see appendices
}

type ParticipantData struct {
	AIControlled  uint8      `json:"m_aiControlled"`  // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId      uint8      `json:"m_driverId"`      // Driver id - see appendix, 255 if network human
	NetworkId     uint8      `json:"m_networkId"`     // Network id – unique identifier for network players
	TeamId        uint8      `json:"m_teamId"`        // Team id - see appendix
	MyTeam        uint8      `json:"m_myTeam"`        // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber    uint8      `json:"m_raceNumber"`    // Race number of the car
	Nationality   uint8      `json:"m_nationality"`   // Nationality of the driver
	Name          [48]string `json:"m_name"`          // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
	YourTelemetry uint8      `json:"m_yourTelemetry"` // The player's UDP setting, 0 = restricted, 1 = public
}

type CarSetupData struct {
	FrontWing              uint8   `json:"m_frontWing"`              // Front wing aero
	RearWing               uint8   `json:"m_rearWing"`               // Rear wing aero
	OnThrottle             uint8   `json:"m_onThrottle"`             // Differential adjustment on throttle (percentage)
	OffThrottle            uint8   `json:"m_offThrottle"`            // Differential adjustment off throttle (percentage)
	FrontCamber            float32 `json:"m_frontCamber"`            // Front camber angle (suspension geometry)
	RearCamber             float32 `json:"m_rearCamber"`             // Rear camber angle (suspension geometry)
	FrontToe               float32 `json:"m_frontToe"`               // Front toe angle (suspension geometry)
	RearToe                float32 `json:"m_rearToe"`                // Rear toe angle (suspension geometry)
	FrontSuspension        uint8   `json:"m_frontSuspension"`        // Front suspension
	RearSuspension         uint8   `json:"m_rearSuspension"`         // Rear suspension
	FrontAntiRollBar       uint8   `json:"m_frontAntiRollBar"`       // Front anti-roll bar
	RearAntiRollBar        uint8   `json:"m_rearAntiRollBar"`        // Front anti-roll bar
	FrontSuspensionHeight  uint8   `json:"m_frontSuspensionHeight"`  // Front ride height
	RearSuspensionHeight   uint8   `json:"m_rearSuspensionHeight"`   // Rear ride height
	BrakePressure          uint8   `json:"m_brakePressure"`          // Brake pressure (percentage)
	BrakeBias              uint8   `json:"m_brakeBias"`              // Brake bias (percentage)
	RearLeftTyrePressure   float32 `json:"m_rearLeftTyrePressure"`   // Rear left tyre pressure (PSI)
	RearRightTyrePressure  float32 `json:"m_rearRightTyrePressure"`  // Rear right tyre pressure (PSI)
	FrontLeftTyrePressure  float32 `json:"m_frontLeftTyrePressure"`  // Front left tyre pressure (PSI)
	FrontRightTyrePressure float32 `json:"m_frontRightTyrePressure"` // Front right tyre pressure (PSI)
	Ballast                uint8   `json:"m_ballast"`                // Ballast
	FuelLoad               float32 `json:"m_fuelLoad"`               // Fuel load
}

type CarTelemetryData struct {
	Speed                   uint16     `json:"m_speed"`                   // Speed of car in kilometres per hour
	Throttle                float32    `json:"m_throttle"`                // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    `json:"m_steer"`                   // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    `json:"m_brake"`                   // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      `json:"m_clutch"`                  // Amount of clutch applied (0 to 100)
	Gear                    int8       `json:"m_gear"`                    // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     `json:"m_engineRPM"`               // Engine RPM
	DRS                     uint8      `json:"m_drs"`                     // 0 = off, 1 = on
	RevLightsPercent        uint8      `json:"m_revLightsPercent"`        // Rev lights indicator (percentage)
	RevLightsBitValue       uint16     `json:"m_revLightsBitValue"`       // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       [4]uint16  `json:"m_brakesTemperature"`       // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   `json:"m_tyresSurfaceTemperature"` // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   `json:"m_tyresInnerTemperature"`   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     `json:"m_engineTemperature"`       // Engine temperature (celsius)
	TyresPressure           [4]float32 `json:"m_tyresPressure"`           // Tyres pressure (PSI)
	SurfaceType             [4]uint8   `json:"m_surfaceType"`             // Driving surface, see appendices
}

type CarStatusData struct {
	TractionControl       uint8   `json:"m_tractionControl"`       // Traction control - 0 = off, 1 = medium, 2 = full
	AntiLockBrakes        uint8   `json:"m_antiLockBrakes"`        // 0 (off) - 1 (on)
	FuelMix               uint8   `json:"m_fuelMix"`               // Fuel mix - 0 = lean, 1 = standard, 2 = rich, 3 = max
	FrontBrakeBias        uint8   `json:"m_frontBrakeBias"`        // Front brake bias (percentage)
	PitLimiterStatus      uint8   `json:"m_pitLimiterStatus"`      // Pit limiter status - 0 = off, 1 = on
	FuelInTank            float32 `json:"m_fuelInTank"`            // Current fuel mass
	FuelCapacity          float32 `json:"m_fuelCapacity"`          // Fuel capacity
	FuelRemainingLaps     float32 `json:"m_fuelRemainingLaps"`     // Fuel remaining in terms of laps (value on MFD)
	MaxRPM                uint16  `json:"m_maxRPM"`                // Cars max RPM, point of rev limiter
	IdleRPM               uint16  `json:"m_idleRPM"`               // Cars idle RPM
	MaxGears              uint8   `json:"m_maxGears"`              // Maximum number of gears
	DRSAllowed            uint8   `json:"m_drsAllowed"`            // 0 = not allowed, 1 = allowed
	DRSActivationDistance uint16  `json:"m_drsActivationDistance"` // 0 = DRS not available, non-zero - DRS will be available in [X] metres
	ActualTyreCompound    uint8   `json:"m_actualTyreCompound"`    // F1 Modern - 16 = C5, 17 = C4, 18 = C3, 19 = C2, 20 = C1, 7 = inter, 8 = wet, F1 Classic - 9 = dry, 10 = wet, F2 – 11 = super soft, 12 = soft, 13 = medium, 14 = hard, 15 = wet
	VisualTyreCompound    uint8   `json:"m_visualTyreCompound"`    // F1 visual (can be different from actual compound), 16 = soft, 17 = medium, 18 = hard, 7 = inter, 8 = wet, F1 Classic – same as above, F2 ‘19, 15 = wet, 19 – super soft, 20 = soft, 21 = medium , 22 = hard
	TyresAgeLaps          uint8   `json:"m_tyresAgeLaps"`          // Age in laps of the current set of tyres
	VehicleFiaFlags       int8    `json:"m_vehicleFiaFlags"`       // -1 = invalid/unknown, 0 = none, 1 = green, 2 = blue, 3 = yellow, 4 = red
	ERSStoreEnergy        float32 `json:"m_ersStoreEnergy"`        // ERS energy store in Joules
	ERSDeployMode         uint8   `json:"m_ersDeployMode"`         // ERS deployment mode, 0 = none, 1 = medium
	// 2 = hotlap, 3 = overtake
	ERSHarvestedThisLapMGUK float32 `json:"m_ersHarvestedThisLapMGUK"` // ERS energy harvested this lap by MGU-K
	ERSHarvestedThisLapMGUH float32 `json:"m_ersHarvestedThisLapMGUH"` // ERS energy harvested this lap by MGU-H
	ERSDeployedThisLap      float32 `json:"m_ersDeployedThisLap"`      // ERS energy deployed this lap
	NetworkPaused           uint8   `json:"m_networkPaused"`           // Whether the car is paused in a network game
}

type FinalClassificationData struct {
	Position         uint8    `json:"m_position"`         // Finishing position
	NumLaps          uint8    `json:"m_numLaps"`          // Number of laps completed
	GridPosition     uint8    `json:"m_gridPosition"`     // Grid position of the car
	Points           uint8    `json:"m_points"`           // Number of points scored
	NumPitStops      uint8    `json:"m_numPitStops"`      // Number of pit stops made
	ResultStatus     uint8    `json:"m_resultStatus"`     // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
	BestLapTimeInMS  uint32   `json:"m_bestLapTimeInMS"`  // Best lap time of the session in milliseconds
	TotalRaceTime    float32  `json:"m_totalRaceTime"`    // Total race time in seconds without penalties
	PenaltiesTime    uint8    `json:"m_penaltiesTime"`    // Total penalties accumulated in seconds
	NumPenalties     uint8    `json:"m_numPenalties"`     // Number of penalties applied to this driver
	NumTyreStints    uint8    `json:"m_numTyreStints"`    // Number of tyres stints up to maximum
	TyreStintsActual [8]uint8 `json:"m_tyreStintsActual"` // Actual tyres used by this driver
	TyreStintsVisual [8]uint8 `json:"m_tyreStintsVisual"` // Visual tyres used by this driver
}

type LobbyInfoData struct {
	AIControlled uint8      `json:"m_aiControlled"` // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId       uint8      `json:"m_teamId"`       // Team id - see appendix (255 if no team currently selected)
	Nationality  uint8      `json:"m_nationality"`  // Nationality of the driver
	Name         [48]string `json:"m_name"`         // Name of participant in UTF-8 format – null terminated Will be truncated with ... (U+2026) if too long
	CarNumber    uint8      `json:"m_carNumber"`    // Car number of the player
	ReadyStatus  uint8      `json:"m_readyStatus"`  // 0 = not ready, 1 = ready, 2 = spectating
}

type CarDamageData struct {
	TyresWear            [4]float32 `json:"m_tyresWear"`            // Tyre wear (percentage)
	TyresDamage          [4]uint8   `json:"m_tyresDamage"`          // Tyre damage (percentage)
	BrakesDamage         [4]uint8   `json:"m_brakesDamage"`         // Brakes damage (percentage)
	FrontLeftWingDamage  uint8      `json:"m_frontLeftWingDamage"`  // Front left wing damage (percentage)
	FrontRightWingDamage uint8      `json:"m_frontRightWingDamage"` // Front right wing damage (percentage)
	RearWingDamage       uint8      `json:"m_rearWingDamage"`       // Rear wing damage (percentage)
	FloorDamage          uint8      `json:"m_floorDamage"`          // Floor damage (percentage)
	DiffuserDamage       uint8      `json:"m_diffuserDamage"`       // Diffuser damage (percentage)
	SidepodDamage        uint8      `json:"m_sidepodDamage"`        // Sidepod damage (percentage)
	DRSFault             uint8      `json:"m_drsFault"`             // Indicator for DRS fault, 0 = OK, 1 = fault
	GearBoxDamage        uint8      `json:"m_gearBoxDamage"`        // Gear box damage (percentage)
	EngineDamage         uint8      `json:"m_engineDamage"`         // Engine damage (percentage)
	EngineMGUHWear       uint8      `json:"m_engineMGUHWear"`       // Engine wear MGU-H (percentage)
	EngineESWear         uint8      `json:"m_engineESWear"`         // Engine wear ES (percentage)
	EngineCEWear         uint8      `json:"m_engineCEWear"`         // Engine wear CE (percentage)
	EngineICEWear        uint8      `json:"m_engineICEWear"`        // Engine wear ICE (percentage)
	EngineMGUKWear       uint8      `json:"m_engineMGUKWear"`       // Engine wear MGU-K (percentage)
	EngineTCWear         uint8      `json:"m_engineTCWear"`         // Engine wear TC (percentage)
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

type DataSourceConnection struct {
	addr *net.UDPAddr
	conn *net.UDPConn
}

/*
	func main() {
		app.CreateNewApp()
		datasource.Proba()
		fmt.Println("hello world")

		connection, err := createDataSourceConnection("127.0.0.1", ":20777")

		if err != nil {
			fmt.Println("Could Not connect")
			return
		}

		for {
			readData(connection.conn)
		}

		//udpConn := datasource.Connect("127.0.0.1", ":20777")

		//datasource.Connect("127.0.0.1", ":20777")
	}
*/
func createDataSourceConnection(ip string, port string) (*DataSourceConnection, error) {
	protocol := "udp"

	addr, err := net.ResolveUDPAddr(protocol, port)

	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP(protocol, addr)

	if err != nil {
		return nil, err
	}

	dataSourceConnection := DataSourceConnection{addr, conn}

	return &dataSourceConnection, nil
}

func readData(conn *net.UDPConn) (*PacketHeader, interface{}, error) {
	buf := make([]byte, 1024+1024/2)

	_, _, err := conn.ReadFromUDP(buf)

	if err != nil {
		return nil, nil, fmt.Errorf("read error: %s", err)
	}

	header := new(PacketHeader)
	if err = ReadPacket(buf, header); err != nil {
		return nil, nil, err
	}

	pack := newPacketById(header.PacketId)
	if pack == nil {
		return nil, nil, fmt.Errorf("invalid packet: %d", header.PacketId)
	}

	if err = ReadPacket(buf, pack); err != nil {
		return nil, nil, fmt.Errorf("%d: %s", header.PacketId, err)
	}

	if header.PacketId == PacketEvent {
		//		details := resolveEventDetails(pack.(*PacketEventData))

		/*pre := pack.(*PacketEventData)
		if details != nil {
			err = ReadPacket(EventDetails[:unsafe.Sizeof(details)], details)
			if err != nil {
				return nil, nil, fmt.Errorf("event packet details read error: %s", err)
			}
		}
		pack = &PacketEventData{
			Header:          pre.Header,
			EventStringCode: pre.EventStringCode,
			EventDetails:    details,
		}*/
	}

	return header, pack, nil
}

func ReadPacket(buf []byte, pack interface{}) error {
	reader := bytes.NewReader(buf)
	if err := binary.Read(reader, binary.LittleEndian, pack); err != nil {
		return err
	}

	return nil
}

func newPacketById(packetId uint8) interface{} {
	switch packetId {
	case PacketMotion:
		return new(PacketMotionData)
	case PacketSession:
		return new(PacketSessionData)
	case PacketLap:
		return new(PacketLapData)
	case PacketEvent:
		return new(PacketEventData)
	case PacketParticipants:
		return new(PacketParticipantsData)
	case PacketCarSetup:
		return new(PacketCarSetupData)
	case PacketCarTelemetry:
		return new(PacketCarTelemetryData)
	case PacketCarStatus:
		return new(PacketCarSetupData)
	case PacketFinalClassification:
		return new(PacketFinalClassificationData)
	case PacketLobbyInfo:
		return new(PacketLobbyInfoData)
	case PacketCarDamage:
		return new(PacketCarDamageData)
	case PacketSessionHistory:
		return new(PacketSessionHistoryData)
	}

	return nil
}

const (
	PacketMotion              uint8 = 0  // Contains all motion data for player’s car – only sent while player is in control
	PacketSession             uint8 = 1  // Data about the session – track, time left
	PacketLap                 uint8 = 2  // Data about all the lap times of cars in the session
	PacketEvent               uint8 = 3  // Various notable events that happen during a session
	PacketParticipants        uint8 = 4  // List of participants in the session, mostly relevant for multiplayer
	PacketCarSetup            uint8 = 5  // Packet detailing car setups for cars in the race
	PacketCarTelemetry        uint8 = 6  // Telemetry data for all cars
	PacketCarStatus           uint8 = 7  // Status data for all cars such as damage
	PacketFinalClassification uint8 = 8  // Final classification confirmation at the end of a race
	PacketLobbyInfo           uint8 = 9  // Information about players in a multiplayer lobby
	PacketCarDamage           uint8 = 10 // Damage status for all cars
	PacketSessionHistory      uint8 = 11 // Lap and tyre data for session
)

func resolveEventDetails(p *PacketEventData) interface{} {
	switch string(p.EventStringCode[:]) {
	case FastestLap:
		return new(FastestLapPacket)
	case Retirement:
		return new(RetirementPacket)
	case TeamMateInPit:
		return new(TeamMateInPitsPacket)
	case RaceWinner:
		return new(RaceWinnerPacket)
	case PenaltyIssued:
		return new(PenaltyPacket)
	case SpeedTrapTriggered:
		return new(SpeedTrapPacket)
	case StartLights:
		return new(StartLightsPacket)
	case LightsOut:
		return new(StartLightsPacket)
	case DriveThroughServed:
		return new(DriveThroughPenaltyServedPacket)
	case StopGoServed:
		return new(StopGoPenaltyServedPacket)
	case Flashback:
		return new(FlashbackPacket)
	case ButtonStatus:
		return new(ButtonsPacket)
	}

	return nil
}
