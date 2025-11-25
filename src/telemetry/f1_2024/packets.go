package f1_2024

import (
	"f1telemetry/src/telemetry/common"
)

type PacketHeader struct {
	PacketFormat            uint16  `json:"m_packetFormat"`     // 2024
	GameYear                uint8   `json:"m_gameYear"`         // 24
	GameMajorVersion        uint8   `json:"m_gameMajorVersion"` // Game major version - "X.00"
	GameMinorVersion        uint8   `json:"m_gameMinorVersion"` // Game minor version - "1.XX"
	PacketVersion           uint8   `json:"m_packetVersion"`    // Version of this packet type, all start from 1
	PacketId                uint8   `json:"m_packetId"`         // Identifier for the packet type, see below
	SessionId               uint64  `json:"m_sessionUID"`       // Unique identifier for the session
	SessionTime             float32 `json:"m_sessionTime"`      // Session timestamp
	FrameIdentifier         uint32  `json:"m_frameIdentifier"`  // Identifier for the frame the data was retrieved
	OverallFrameIdentifier  uint32  `json:"m_overallFrameIdentifier"`
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

func (p *PacketHeader) GetPlayerCarIndex() uint8 {
	return p.PlayerCarIndex
}

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

func (p PacketMotionData) PacketHeader() common.Header {
	return &p.Header
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

type PacketSessionData struct {
	Header PacketHeader `json:"m_header"`

	Weather          uint8 `json:"m_weather"` // Weather - 0 = clear, 1 = light cloud, 2 = overcast, 3 = light rain, 4 = heavy rain, 5 = storm
	TrackTemperature int8  `json:"m_trackTemperature"`
	AirTemperature   int8  `json:"m_airTemperature"`

	TotalLaps   uint8  `json:"m_totalLaps"`
	TrackLength uint16 `json:"m_trackLength"`
	SessionType uint8  `json:"m_sessionType"` // 0 = unknown, 1 = P1, 2 = P2, 3 = P3, 4 = Short P, 5 = Q1, 6 = Q2, 7 = Q3, 8 = Short Q, 9 = OSQ, 10 = R, 11 = R2, 12 = R3, 13 = Time Trial
	TrackId     int8   `json:"m_trackId"`     // -1 for unknown, 0-21 for tracks, see appendix
	Formula     uint8  `json:"m_formula"`     // Formula, 0 = F1 Modern, 1 = F1 Classic, 2 = F2, 3 = F1 Generic

	SessionTimeLeft      uint16          `json:"m_sessionTimeLeft"`
	SessionDuration      uint16          `json:"m_sessionDuration"`
	PitSpeedLimit        uint8           `json:"m_pitSpeedLimit"`
	GamePaused           bool            `json:"m_gamePaused"`
	IsSpectating         uint8           `json:"m_isSpectating"`
	SpectatorCarIndex    uint8           `json:"m_spectatorCarIndex"`
	SLIProNativeSupport  uint8           `json:"m_sliProNativeSupport"`
	NumberOfMarshalZones uint8           `json:"m_numMarshalZones"`
	MarshalZones         [21]MarshalZone `json:"m_marshalZones"`
	SafetyCarStatus      uint8           `json:"m_safetyCarStatus"` // 0 = no safety car, 1 = full, 2 = virtual, 3 = formation lap
	NetworkGame          uint8           `json:"m_networkGame"`     // 0 = offline, 1 = online

	NumWeatherForecastSamples uint8                     `json:"m_numWeatherForecastSamples"`
	WeatherForecastSamples    [56]WeatherForecastSample `json:"m_weatherForecastSamples"`
	ForecastAccuracy          uint8                     `json:"m_forecastAccuracy"`

	AIDifficulty uint8 `json:"m_aiDifficulty"`

	SeasonLinkIdentifier            uint32    `json:"m_seasonLinkIdentifier"`            // Identifier for season - persists across saves
	WeekendLinkIdentifier           uint32    `json:"m_weekendLinkIdentifier"`           // Identifier for weekend - persists across saves
	SessionLinkIdentifier           uint32    `json:"m_sessionLinkIdentifier"`           // Identifier for session - persists across saves
	PitStopWindowIdealLap           uint8     `json:"m_pitStopWindowIdealLap"`           // Ideal lap to pit on for current strategy (player)
	PitStopWindowLatestLap          uint8     `json:"m_pitStopWindowLatestLap"`          // Latest lap to pit on for current strategy (player)
	PitStopRejoinPosition           uint8     `json:"m_pitStopRejoinPosition"`           // Predicted position to rejoin at (player)
	SteeringAssist                  bool      `json:"m_steeringAssist"`                  // 0 = off, 1 = on
	BrakingAssist                   uint8     `json:"m_brakingAssist"`                   // 0 = off, 1 = low, 2 = medium, 3 = high
	GearboxAssist                   uint8     `json:"m_gearboxAssist"`                   // 1 = manual, 2 = manual & suggested gear, 3 = auto
	PitAssist                       bool      `json:"m_pitAssist"`                       // 0 = off, 1 = on
	PitReleaseAssist                bool      `json:"m_pitReleaseAssist"`                // 0 = off, 1 = on
	ERSAssist                       bool      `json:"m_ERSAssist"`                       // 0 = off, 1 = on
	DRSAssist                       bool      `json:"m_DRSAssist"`                       // 0 = off, 1 = on
	DynamicRacingLine               uint8     `json:"m_dynamicRacingLine"`               // 0 = off, 1 = corners only, 2 = full
	DynamicRacingLineType           uint8     `json:"m_dynamicRacingLineType"`           // 0 = 2D, 1 = 3D
	GameMode                        uint8     `json:"m_gameMode"`                        // Game mode id - see appendix
	RuleSet                         uint8     `json:"m_ruleSet"`                         // Ruleset - see appendix
	TimeOfDay                       uint32    `json:"m_timeOfDay"`                       // Local time of day in minutes since midnight
	SessionLength                   uint8     `json:"m_sessionLength"`                   // 0 = None, 2 = Short, 3 = Medium, 4 = Full
	SpeedUnitsLeadPlayer            uint8     `json:"m_speedUnitsLeadPlayer"`            // 0 = KPH, 1 = MPH
	TemperatureUnitsLeadPlayer      uint8     `json:"m_temperatureUnitsLeadPlayer"`      // 0 = Celsius, 1 = Fahrenheit
	SpeedUnitsSecondaryPlayer       uint8     `json:"m_speedUnitsSecondaryPlayer"`       // 0 = KPH, 1 = MPH
	TemperatureUnitsSecondaryPlayer uint8     `json:"m_temperatureUnitsSecondaryPlayer"` // 0 = Celsius, 1 = Fahrenheit
	NumberOfSafetyCarPeriods        uint8     `json:"m_numSafetyCarPeriods"`             // Number of safety cars called during session
	NumberOfVirtualSafetyCarPeriods uint8     `json:"m_numVirtualSafetyCarPeriods"`      // Number of virtual safety cars called
	NumberOfRedFlagPeriods          uint8     `json:"m_numRedFlagPeriods"`               // Number of red flags called during session
	EqualCarPerformance             uint8     `json:"m_equalCarPerformance"`             // 0 = Off, 1 = On
	RecoveryMode                    uint8     `json:"m_recoveryMode"`                    // 0 = None, 1 = Flashbacks, 2 = Auto-recovery
	FlashbackLimit                  uint8     `json:"m_flashbackLimit"`                  // 0 = Low, 1 = Medium, 2 = High, 3 = Unlimited
	SurfaceType                     uint8     `json:"m_surfaceType"`                     // 0 = Simplified, 1 = Realistic
	LowFuelMode                     uint8     `json:"m_lowFuelMode"`                     // 0 = Easy, 1 = Hard
	RaceStarts                      uint8     `json:"m_raceStarts"`                      // 0 = Manual, 1 = Assisted
	TyreTemperature                 uint8     `json:"m_tyreTemperature"`                 // 0 = Surface only, 1 = Surface & Carcass
	PitLaneTyreSim                  uint8     `json:"m_pitLaneTyreSim"`                  // 0 = On, 1 = Off
	CarDamage                       uint8     `json:"m_carDamage"`                       // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Simulation
	CarDamageRate                   uint8     `json:"m_carDamageRate"`                   // 0 = Reduced, 1 = Standard, 2 = Simulation
	Collisios                       uint8     `json:"m_collisios"`                       // 0 = Off, 1 = Player-to-Player Off, 2 = On
	CollisionsOffForFirstLapOnly    uint8     `json:"m_collisionsOffForFirstLapOnly"`    // 0 = Disabled, 1 = Enabled
	MultiplayerUnsafePitRelease     uint8     `json:"m_mpUnsafePitRelease"`              // 0 = On, 1 = Off (Multiplayer)
	MultiplayerOffForGriefing       uint8     `json:"m_mpOffForGriefing"`                // 0 = Disabled, 1 = Enabled (Multiplayer)
	CornerCuttingStringency         uint8     `json:"m_cornerCuttingStringency"`         // 0 = Regular, 1 = Strict
	ParcFermeRules                  uint8     `json:"m_parcFermErules"`                  // 0 = Off, 1 = On
	PitStopExperience               uint8     `json:"m_pitStopExperience"`               // 0 = Automatic, 1 = Broadcast, 2 = Immersive
	SafetyCar                       uint8     `json:"m_safetyCar"`                       // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	SafetyCarExperience             uint8     `json:"m_safetyCarExperience"`             // 0 = Broadcast, 1 = Immersive
	FormationLap                    uint8     `json:"m_formationLap"`                    // 0 = Off, 1 = On
	FormationLapExperience          uint8     `json:"m_formationLapExperience"`          // 0 = Broadcast, 1 = Immersive
	RedFlags                        uint8     `json:"m_redFlags"`                        // 0 = Off, 1 = Reduced, 2 = Standard, 3 = Increased
	AffectsLicenceLevelSolo         uint8     `json:"m_affectsLicenceLevelSolo"`         // 0 = Off, 1 = On
	AffectsLicenceLevelMP           uint8     `json:"m_affectsLicenceLevelMP"`           // 0 = Off, 1 = On
	NumSessionsInWeekend            uint8     `json:"m_numSessionsInWeekend"`            // Number of sessions in following array
	WeekendStructure                [12]uint8 `json:"m_weekendStructure"`                // List of session types that show weekend structure
	Sector2LapDistanceStart         float32   `json:"m_sector2LapDistanceStart"`         // Distance (m) where sector 2 starts
	Sector3LapDistanceStart         float32   `json:"m_sector3LapDistanceStart"`         // Distance (m) where sector 3 starts
}

func (p PacketSessionData) PacketHeader() common.Header {
	return &p.Header
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

type PacketLapData struct {
	Header PacketHeader `json:"m_header"`

	LapData [22]LapData `json:"m_lapData"` // Lap data for all cars on track

	TimeTrialPBCarIdx    uint8 `json:"m_timeTrialPBCarIdx"`    // Index of Personal Best car in time trial (255 if invalid)
	TimeTrialRivalCarIdx uint8 `json:"m_timeTrialRivalCarIdx"` // Index of Rival car in time trial (255 if invalid)
}

func (p PacketLapData) PacketHeader() common.Header {
	return &p.Header
}

type LapData struct {
	LastLapTimeInMS              uint32  `json:"m_lastLapTimeInMS"`              // Last lap time in milliseconds
	CurrentLapTimeInMS           uint32  `json:"m_currentLapTimeInMS"`           // Current time around the lap in milliseconds
	Sector1TimeMSPart            uint16  `json:"m_sector1TimeMSPart"`            // Sector 1 time milliseconds part
	Sector1TimeMinutesPart       uint8   `json:"m_sector1TimeMinutesPart"`       // Sector 1 whole minute part
	Sector2TimeMSPart            uint16  `json:"m_sector2TimeMSPart"`            // Sector 2 time milliseconds part
	Sector2TimeMinutesPart       uint8   `json:"m_sector2TimeMinutesPart"`       // Sector 2 whole minute part
	DeltaToCarInFrontMSPart      uint16  `json:"m_deltaToCarInFrontMSPart"`      // Time delta to car in front milliseconds part
	DeltaToCarInFrontMinutesPart uint8   `json:"m_deltaToCarInFrontMinutesPart"` // Time delta to car in front whole minute part
	DeltaToRaceLeaderMSPart      uint16  `json:"m_deltaToRaceLeaderMSPart"`      // Time delta to race leader milliseconds part
	DeltaToRaceLeaderMinutesPart uint8   `json:"m_deltaToRaceLeaderMinutesPart"` // Time delta to race leader whole minute part
	LapDistance                  float32 `json:"m_lapDistance"`                  // Distance vehicle is around current lap in metres – could be negative if line hasn’t been crossed yet
	TotalDistance                float32 `json:"m_totalDistance"`                // Total distance travelled in session in metres – could be negative if line hasn’t been crossed yet
	SafetyCarDelta               float32 `json:"m_safetyCarDelta"`               // Delta in seconds for safety car
	CarPosition                  uint8   `json:"m_carPosition"`                  // Car race position
	CurrentLapNum                uint8   `json:"m_currentLapNum"`                // Current lap number
	PitStatus                    uint8   `json:"m_pitStatus"`                    // 0 = none, 1 = pitting, 2 = in pit area
	NumPitStops                  uint8   `json:"m_numPitStops"`                  // Number of pit stops taken in this race
	Sector                       uint8   `json:"m_sector"`                       // 0 = sector1, 1 = sector2, 2 = sector3
	CurrentLapInvalid            uint8   `json:"m_currentLapInvalid"`            // Current lap invalid - 0 = valid, 1 = invalid
	Penalties                    uint8   `json:"m_penalties"`                    // Accumulated time penalties in seconds to be added
	TotalWarnings                uint8   `json:"m_totalWarnings"`                // Accumulated number of warnings issued
	CornerCuttingWarnings        uint8   `json:"m_cornerCuttingWarnings"`        // Accumulated number of corner cutting warnings issued
	NumUnservedDriveThroughPens  uint8   `json:"m_numUnservedDriveThroughPens"`  // Num drive through pens left to serve
	NumUnservedStopGoPens        uint8   `json:"m_numUnservedStopGoPens"`        // Num stop go pens left to serve
	GridPosition                 uint8   `json:"m_gridPosition"`                 // Grid position the vehicle started the race in
	DriverStatus                 uint8   `json:"m_driverStatus"`                 // Status of driver - 0 = in garage, 1 = flying lap, 2 = in lap, 3 = out lap, 4 = on track
	ResultStatus                 uint8   `json:"m_resultStatus"`                 // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
	PitLaneTimerActive           uint8   `json:"m_pitLaneTimerActive"`           // Pit lane timing, 0 = inactive, 1 = active
	PitLaneTimeInLaneInMS        uint16  `json:"m_pitLaneTimeInLaneInMS"`        // If active, the current time spent in the pit lane in ms
	PitStopTimerInMS             uint16  `json:"m_pitStopTimerInMS"`             // Time of the actual pit stop in ms
	PitStopShouldServePen        uint8   `json:"m_pitStopShouldServePen"`        // Whether the car should serve a penalty at this stop
	SpeedTrapFastestSpeed        float32 `json:"m_speedTrapFastestSpeed"`        // Fastest speed through speed trap for this car in kmph
	SpeedTrapFastestLap          uint8   `json:"m_speedTrapFastestLap"`          // Lap no the fastest speed was achieved, 255 = not set

}

type PacketEventData struct {
	Header          PacketHeader `json:"m_packetHeader"`
	EventStringCode [4]uint8     `json:"m_eventStringCode"` // Event string code, see below
}

func (p PacketEventData) PacketHeader() common.Header {
	return &p.Header
}

func (p PacketEventData) EventString() string {
	return string(p.EventStringCode[0:])
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

type PacketPenalty struct {
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

type PacketButtons struct {
	EventData    PacketEventData `json:"m_packetEventData"`
	ButtonStatus uint32          `json:"m_buttonStatus"` // Bit flags specifying which buttons are being pressed currently - see appendices
}

type OvertakePacket struct {
	EventData                PacketEventData `json:"m_packetEventData"`
	OvertakingVehicleIdx     uint8           `json:"overtakingVehicleIdx"`     // Vehicle index of the vehicle overtaking
	BeingOvertakenVehicleIdx uint8           `json:"beingOvertakenVehicleIdx"` // Vehicle index of the vehicle being overtaken
}

type CollisionPacket struct {
	EventData   PacketEventData `json:"m_packetEventData"`
	Vehicle1Idx uint8           `json:"vehicle1Idx"` // Vehicle index of the first vehicle involved in the collision
	Vehicle2Idx uint8           `json:"vehicle2Idx"` // Vehicle index of the second vehicle involved in the collision
}

type SafetyCarPacket struct {
	EventData     PacketEventData `json:"m_packetEventData"`
	SafetyCarType uint8           `json:"safetyCarType"` // 0 = No Safety Car, 1 = Full Safety Car, 2 = Virtual Safety Car, 3 = Formation Lap Safety Car            // Vehicle index of the vehicle overtaking
	EventType     uint8           `json:"eventType"`     // 0 = Deployed, 1 = Returning, 2 = Returned, 3 = Resume Race
}

func (p SafetyCarPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}
func (p OvertakePacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p SpeedTrapPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p FlashbackPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p PacketPenalty) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p StopGoPenaltyServedPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p DriveThroughPenaltyServedPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p StartLightsPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p RaceWinnerPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p TeamMateInPitsPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p FastestLapPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p PacketButtons) PacketHeader() common.Header {
	return &p.EventData.Header
}

func (p RetirementPacket) PacketHeader() common.Header {
	return &p.EventData.Header
}

type PacketParticipantsData struct {
	Header          PacketHeader        `json:"m_header"`
	NumActiveCars   uint8               `json:"m_numActiveCars"` // Number of active cars in the data – should match number of cars on HUD
	ParticipantData [22]ParticipantData `json:"m_participants"`
}

func (p PacketParticipantsData) PacketHeader() common.Header {
	return &p.Header
}

type ParticipantData struct {
	AIControlled    bool     `json:"m_aiControlled"`    // Whether the vehicle is AI (1) or Human (0) controlled
	DriverId        uint8    `json:"m_driverId"`        // Driver id - see appendix, 255 if network human
	NetworkId       uint8    `json:"m_networkId"`       // Network id – unique identifier for network players
	TeamId          uint8    `json:"m_teamId"`          // Team id - see appendix
	MyTeam          bool     `json:"m_myTeam"`          // My team flag – 1 = My Team, 0 = otherwise
	RaceNumber      uint8    `json:"m_raceNumber"`      // Race number of the car
	Nationality     uint8    `json:"m_nationality"`     // Nationality of the driver
	Name            [48]byte `json:"m_name"`            // Name of participant in UTF-8 format – null terminated Will be truncated with … (U+2026) if too long
	YourTelemetry   uint8    `json:"m_yourTelemetry"`   // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8    `json:"m_showOnlineNames"` // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16   `json:"m_techLevel"`       // F1 World tech level
	Platform        uint8    `json:"m_platform"`        // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
}

type PacketCarSetupData struct {
	Header       PacketHeader     `json:"m_header"`
	CarSetupData [22]CarSetupData `json:"m_carSetups"`
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

func (p PacketCarSetupData) PacketHeader() common.Header {
	return &p.Header
}

type PacketCarStatus struct {
	Header             PacketHeader      `json:"m_header"`
	CarStatusData      [22]CarStatusData `json:"m_carStatusData"`
	NextFrontWingValue float32           `json:"m_nextFrontWingValue"` // Value of front wing after next pit stop - player only
}

func (p PacketCarStatus) PacketHeader() common.Header {
	return &p.Header
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

type PacketFinalClassificationData struct {
	Header                  PacketHeader                `json:"m_header"`
	NumCars                 uint8                       `json:"m_numCars"` // Number of cars in the final classification
	FinalClassificationData [22]FinalClassificationData `json:"m_classificationData"`
}

func (p *PacketFinalClassificationData) PacketHeader() common.Header {
	return &p.Header
}

type FinalClassificationData struct {
	Position          uint8    `json:"m_position"`          // Finishing position
	NumLaps           uint8    `json:"m_numLaps"`           // Number of laps completed
	GridPosition      uint8    `json:"m_gridPosition"`      // Grid position of the car
	Points            uint8    `json:"m_points"`            // Number of points scored
	NumPitStops       uint8    `json:"m_numPitStops"`       // Number of pit stops made
	ResultStatus      uint8    `json:"m_resultStatus"`      // Result status - 0 = invalid, 1 = inactive, 2 = active, 3 = finished, 4 = didnotfinish, 5 = disqualified, 6 = not classified, 7 = retired
	BestLapTimeInMS   uint32   `json:"m_bestLapTimeInMS"`   // Best lap time of the session in milliseconds
	TotalRaceTime     float32  `json:"m_totalRaceTime"`     // Total race time in seconds without penalties
	PenaltiesTime     uint8    `json:"m_penaltiesTime"`     // Total penalties accumulated in seconds
	NumPenalties      uint8    `json:"m_numPenalties"`      // Number of penalties applied to this driver
	NumTyreStints     uint8    `json:"m_numTyreStints"`     // Number of tyres stints up to maximum
	TyreStintsActual  [8]uint8 `json:"m_tyreStintsActual"`  // Actual tyres used by this driver
	TyreStintsVisual  [8]uint8 `json:"m_tyreStintsVisual"`  // Visual tyres used by this driver
	TyreStintsEndLaps [8]uint8 `json:"m_tyreStintsEndLaps"` // The lap number stints end on
}

type PacketLobbyInfoData struct {
	Header PacketHeader `json:"m_header"`

	// Packet specific data
	NumPlayers   uint8             `json:"m_numPlayers"` // Number of players in the lobby data
	LobbyPlayers [22]LobbyInfoData `json:"m_lobbyPlayers"`
}

func (p *PacketLobbyInfoData) PacketHeader() common.Header {
	return &p.Header
}

type LobbyInfoData struct {
	AIControlled    uint8      `json:"m_aiControlled"`    // Whether the vehicle is AI (1) or Human (0) controlled
	TeamId          uint8      `json:"m_teamId"`          // Team id - see appendix (255 if no team currently selected)
	Nationality     uint8      `json:"m_nationality"`     // Nationality of the driver
	Platform        uint8      `json:"m_platform"`        // 1 = Steam, 3 = PlayStation, 4 = Xbox, 6 = Origin, 255 = unknown
	Name            [48]string `json:"m_name"`            // Name of participant in UTF-8 format – null terminated Will be truncated with ... (U+2026) if too long
	CarNumber       uint8      `json:"m_carNumber"`       // Car number of the player
	YourTelemetry   uint8      `json:"m_yourTelemetry"`   // The player's UDP setting, 0 = restricted, 1 = public
	ShowOnlineNames uint8      `json:"m_showOnlineNames"` // The player's show online names setting, 0 = off, 1 = on
	TechLevel       uint16     `json:"m_techLevel"`       // F1 World tech level
	ReadyStatus     uint8      `json:"m_readyStatus"`     // 0 = not ready, 1 = ready, 2 = spectating
}

type PacketCarDamageData struct {
	Header PacketHeader `json:"m_header"`

	CarDamageData [22]CarDamageData `json:"m_carDamageData"`
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
	EngineBlown          uint8      `json:"m_engineBlown"`          // Engine blown, 0 = OK, 1 = fault
	EngineSeized         uint8      `json:"m_engineSeized"`         // Engine seized, 0 = OK, 1 = fault
}

func (p PacketCarDamageData) PacketHeader() common.Header {
	return &p.Header
}

type PacketSessionHistoryData struct {
	Header               PacketHeader            `json:"m_header"`
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

func (p PacketSessionHistoryData) PacketHeader() common.Header {
	return &p.Header
}

type TyreSetData struct {
	ActualTyreCompound uint8 `json:"m_actualTyreCompound"` // Actual tyre compound used
	VisualTyreCompound uint8 `json:"m_visualTyreCompound"` // Visual tyre compound used
	Wear               uint8 `json:"m_wear"`               // Tyre wear (percentage)
	Available          uint8 `json:"m_available"`          // Whether this set is currently available
	RecommendedSession uint8 `json:"m_recommendedSession"` // Recommended session for tyre set
	LifeSpan           uint8 `json:"m_lifeSpan"`           // Laps left in this tyre set
	UsableLife         uint8 `json:"m_usableLife"`         // Max number of laps recommended for this compound
	LapDeltaTime       int16 `json:"m_lapDeltaTime"`       // Lap delta time in milliseconds compared to fitted set
	Fitted             uint8 `json:"m_fitted"`             // Whether the set is fitted or not
}

type PacketTyreSetsData struct {
	Header      PacketHeader    `json:"m_header"`
	CarIdx      uint8           `json:"m_carIdx"`      // Index of the car this data relates to
	TyreSetData [20]TyreSetData `json:"m_tyreSetData"` // 13 (dry) + 7 (wet)
	FittedIdx   uint8           `json:"m_fittedIdx"`   // Index into array of fitted tyre
}

func (p PacketTyreSetsData) PacketHeader() common.Header {
	return &p.Header
}

type LapHistoryData struct {
	LapTimeInMS            uint32 `json:"m_lapTimeInMS"`            // Lap time in milliseconds
	Sector1TimeMSPart      uint16 `json:"m_sector1TimeMSPart"`      // Sector 1 milliseconds part
	Sector1TimeMinutesPart uint8  `json:"m_sector1TimeMinutesPart"` // Sector 1 whole minutes part
	Sector2TimeMSPart      uint16 `json:"m_sector2TimeMSPart"`      // Sector 2 milliseconds part
	Sector2TimeMinutesPart uint8  `json:"m_sector2TimeMinutesPart"` // Sector 2 whole minutes part
	Sector3TimeMSPart      uint16 `json:"m_sector3TimeMSPart"`      // Sector 3 milliseconds part
	Sector3TimeMinutesPart uint8  `json:"m_sector3TimeMinutesPart"` // Sector 3 whole minutes part
	LapValidBitFlags       uint8  `json:"m_lapValidBitFlags"`       // 0x01 lap valid, 0x02 sector1 valid, 0x04 sector2 valid, 0x08 sector3 valid
}

type TyreStintHistoryData struct {
	EndLap             uint8 `json:"m_endLap"`             // Lap the tyre usage ends on (255 of current tyre)
	TyreActualCompound uint8 `json:"m_tyreActualCompound"` // Actual tyres used by this driver
	TyreVisualCompound uint8 `json:"m_tyreVisualCompound"` // Visual tyres used by this driver
}

type PacketCarTelemetryData struct {
	Header                       PacketHeader         `json:"m_header"`
	CarTelemetryData             [22]CarTelemetryData `json:"m_carTelemetryData"`
	MFDPanelIndex                uint8                `json:"m_mfdPanelIndex"`                // Index of MFD panel open - 255 = MFD closed, Single player, race – 0 = Car setup, 1 = Pits, 2 = Damage, 3 =  Engine, 4 = Temperatures, May vary depending on game mode
	MFDPanelIndexSecondaryPlayer uint8                `json:"m_mfdPanelIndexSecondaryPlayer"` // See above
	SuggestedGear                int8                 `json:"m_suggestedGear"`                // Suggested gear for the player (1-8), 0 if no gear suggested
}

func (p PacketCarTelemetryData) PacketHeader() common.Header {
	return &p.Header
}

type CarTelemetryData struct {
	Speed                   uint16     `json:"m_speed"`                   // Speed of car in kilometres per hour
	Throttle                float32    `json:"m_throttle"`                // Amount of throttle applied (0.0 to 1.0)
	Steer                   float32    `json:"m_steer"`                   // Steering (-1.0 (full lock left) to 1.0 (full lock right))
	Brake                   float32    `json:"m_brake"`                   // Amount of brake applied (0.0 to 1.0)
	Clutch                  uint8      `json:"m_clutch"`                  // Amount of clutch applied (0 to 100)
	Gear                    int8       `json:"m_gear"`                    // Gear selected (1-8, N=0, R=-1)
	EngineRPM               uint16     `json:"m_engineRPM"`               // Engine RPM
	DRS                     bool       `json:"m_drs"`                     // 0 = off, 1 = on
	RevLightsPercent        uint8      `json:"m_revLightsPercent"`        // Rev lights indicator (percentage)
	RevLightsBitValue       uint16     `json:"m_revLightsBitValue"`       // Rev lights (bit 0 = leftmost LED, bit 14 = rightmost LED)
	BrakesTemperature       [4]uint16  `json:"m_brakesTemperature"`       // Brakes temperature (celsius)
	TyresSurfaceTemperature [4]uint8   `json:"m_tyresSurfaceTemperature"` // Tyres surface temperature (celsius)
	TyresInnerTemperature   [4]uint8   `json:"m_tyresInnerTemperature"`   // Tyres inner temperature (celsius)
	EngineTemperature       uint16     `json:"m_engineTemperature"`       // Engine temperature (celsius)
	TyresPressure           [4]float32 `json:"m_tyresPressure"`           // Tyres pressure (PSI)
	SurfaceType             [4]uint8   `json:"m_surfaceType"`             // Driving surface, see appendices
}

type PacketMotionExData struct {
	Header PacketHeader `json:"m_header"`

	// Extra player car ONLY data
	SuspensionPosition     [4]float32 `json:"m_suspensionPosition"`     // RL, RR, FL, FR
	SuspensionVelocity     [4]float32 `json:"m_suspensionVelocity"`     // RL, RR, FL, FR
	SuspensionAcceleration [4]float32 `json:"m_suspensionAcceleration"` // RL, RR, FL, FR
	WheelSpeed             [4]float32 `json:"m_wheelSpeed"`             // Speed of each wheel
	WheelSlipRatio         [4]float32 `json:"m_wheelSlipRatio"`         // Slip ratio for each wheel
	WheelSlipAngle         [4]float32 `json:"m_wheelSlipAngle"`         // Slip angles for each wheel
	WheelLatForce          [4]float32 `json:"m_wheelLatForce"`          // Lateral forces for each wheel
	WheelLongForce         [4]float32 `json:"m_wheelLongForce"`         // Longitudinal forces for each wheel
	HeightOfCOGAboveGround float32    `json:"m_heightOfCOGAboveGround"` // Height of centre of gravity above ground
	LocalVelocityX         float32    `json:"m_localVelocityX"`         // Velocity in local space – metres/s
	LocalVelocityY         float32    `json:"m_localVelocityY"`         // Velocity in local space
	LocalVelocityZ         float32    `json:"m_localVelocityZ"`         // Velocity in local space
	AngularVelocityX       float32    `json:"m_angularVelocityX"`       // Angular velocity x-component – radians/s
	AngularVelocityY       float32    `json:"m_angularVelocityY"`       // Angular velocity y-component
	AngularVelocityZ       float32    `json:"m_angularVelocityZ"`       // Angular velocity z-component
	AngularAccelerationX   float32    `json:"m_angularAccelerationX"`   // Angular acceleration x-component – radians/s/s
	AngularAccelerationY   float32    `json:"m_angularAccelerationY"`   // Angular acceleration y-component
	AngularAccelerationZ   float32    `json:"m_angularAccelerationZ"`   // Angular acceleration z-component
	FrontWheelsAngle       float32    `json:"m_frontWheelsAngle"`       // Current front wheels angle in radians
	WheelVertForce         [4]float32 `json:"m_wheelVertForce"`         // Vertical forces for each wheel
	FrontAeroHeight        float32    `json:"m_frontAeroHeight"`        // Front plank edge height above road surface
	RearAeroHeight         float32    `json:"m_rearAeroHeight"`         // Rear plank edge height above road surface
	FrontRollAngle         float32    `json:"m_frontRollAngle"`         // Roll angle of the front suspension
	RearRollAngle          float32    `json:"m_rearRollAngle"`          // Roll angle of the rear suspension
	ChassisYaw             float32    `json:"m_chassisYaw"`             // Yaw angle of the chassis relative to the direction of motion - radians
}

func (p PacketMotionExData) PacketHeader() common.Header {
	return &p.Header
}

type TimeTrialDataSet struct {
	CarIdx              uint8  `json:"m_carIdx"`              // Index of the car this data relates to
	TeamId              uint8  `json:"m_teamId"`              // Team id - see appendix
	LapTimeInMS         uint32 `json:"m_lapTimeInMS"`         // Lap time in milliseconds
	Sector1TimeInMS     uint32 `json:"m_sector1TimeInMS"`     // Sector 1 time in milliseconds
	Sector2TimeInMS     uint32 `json:"m_sector2TimeInMS"`     // Sector 2 time in milliseconds
	Sector3TimeInMS     uint32 `json:"m_sector3TimeInMS"`     // Sector 3 time in milliseconds
	TractionControl     uint8  `json:"m_tractionControl"`     // 0 = off, 1 = medium, 2 = full
	GearboxAssist       uint8  `json:"m_gearboxAssist"`       // 1 = manual, 2 = manual & suggested gear, 3 = auto
	AntiLockBrakes      uint8  `json:"m_antiLockBrakes"`      // 0 = off, 1 = on
	EqualCarPerformance uint8  `json:"m_equalCarPerformance"` // 0 = Realistic, 1 = Equal
	CustomSetup         uint8  `json:"m_customSetup"`         // 0 = No, 1 = Yes
	Valid               uint8  `json:"m_valid"`               // 0 = invalid, 1 = valid
}

type PacketTimeTrialData struct {
	Header                   PacketHeader     `json:"m_header"`
	PlayerSessionBestDataSet TimeTrialDataSet `json:"m_playerSessionBestDataSet"` // Player session best data set
	PersonalBestDataSet      TimeTrialDataSet `json:"m_personalBestDataSet"`      // Personal best data set
	RivalDataSet             TimeTrialDataSet `json:"m_rivalDataSet"`             // Rival data set
}

func (p PacketTimeTrialData) PacketHeader() common.Header {
	return &p.Header
}
