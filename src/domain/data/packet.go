package data

type Session struct {
	FrameId         uint32         `json:"frameId"`
	Assist          SettingsAssist `json:"assist"`
	Environment     Environment    `json:"environment"`
	AIDifficulty    uint8          `json:"AIDifficulty"`
	SessionType     string         `json:"sessionType"`
	Formula         string         `json:"formula"`
	SessionTimeLeft uint16         `json:"sessionTimeLeft"`
	SessionDuration uint16         `json:"sessionDuration"`
	GamePaused      bool           `json:"isGamePaused"`
}

type StartLight struct {
	NumberOfLights uint8 `json:"NumberOfLights"`
}
type LobbyInfo struct {
}

/*
	type PacketSessionData struct {
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
		WeatherForecastSamples    [56]WeatherForecastSample `json:"m_weatherForecastSamples"`
		ForecastAccuracy          uint8                     `json:"m_forecastAccuracy"`

		AIDifficulty uint8 `json:"m_aiDifficulty"`

		SeasonLinkIdentifier   uint32 `json:"m_seasonLinkIdentifier"`   // Identifier for season - persists across saves
		WeekendLinkIdentifier  uint32 `json:"m_weekendLinkIdentifier"`  // Identifier for weekend - persists across saves
		SessionLinkIdentifier  uint32 `json:"m_sessionLinkIdentifier"`  // Identifier for session - persists across saves
		PitStopWindowIdealLap  uint8  `json:"m_pitStopWindowIdealLap"`  // Ideal lap to pit on for current strategy (player)
		PitStopWindowLatestLap uint8  `json:"m_pitStopWindowLatestLap"` // Latest lap to pit on for current strategy (player)
		PitStopRejoinPosition  uint8  `json:"m_pitStopRejoinPosition"`  // Predicted position to rejoin at (player)
		SteeringAssist         bool   `json:"m_steeringAssist"`         // 0 = off, 1 = on
		BrakingAssist          uint8  `json:"m_brakingAssist"`          // 0 = off, 1 = low, 2 = medium, 3 = high
		GearboxAssist          uint8  `json:"m_gearboxAssist"`          // 1 = manual, 2 = manual & suggested gear, 3 = auto
		PitAssist              bool   `json:"m_pitAssist"`              // 0 = off, 1 = on
		PitReleaseAssist       bool   `json:"m_pitReleaseAssist"`       // 0 = off, 1 = on
		ERSAssist              bool   `json:"m_ERSAssist"`              // 0 = off, 1 = on
		DRSAssist              bool   `json:"m_DRSAssist"`              // 0 = off, 1 = on
		DynamicRacingLine      uint8  `json:"m_dynamicRacingLine"`      // 0 = off, 1 = corners only, 2 = full
		DynamicRacingLineType  uint8  `json:"m_dynamicRacingLineType"`  // 0 = 2D, 1 = 3D
	}
*/
type Environment struct {
	AirTemperature   int8   `json:"airTemperature"`
	TrackTemperature int8   `json:"trackTemperature"`
	Weather          string `json:"weather"`
}

type Track struct {
	Id          int8   `json:"id"`
	Name        string `json:"name"`
	TotalLaps   uint8  `json:"totalLaps"`
	TrackLength uint16 `json:"trackLength"`
}

type SettingsAssist struct {
	SteeringAssist        bool   `json:"isSteeringAssisted"`
	BrakingAssist         string `json:"brakingAssist"`
	AntiLockBrakes        bool   `json:"isABSEnabled"`
	TractionControl       uint8  `json:"tractionControlLevel"`
	RacingLine            string `json:"RacingLineLevel"`
	DynamicRacingLineType uint8  `json:"dynamic_racing_line_type"`
	GearboxAssist         uint8  `json:"gearbox_assist"`
	PitAssist             bool   `json:"pit_assist"`
	PitReleaseAssist      bool   `json:"pit_release_assist"`
	DRSAssist             bool   `json:"drs_assist"`
	ERSAssist             bool   `json:"ers_assist"`
}

type Button struct {
	FrameId uint32 `json:"frameId"`
	Status  uint32 `json:"status"`
}

type Participants struct {
	FrameId      uint32   `json:"frameId"`
	ActiveCarNum uint8    `json:"activeCarNum"`
	Drivers      []Driver `json:"drivers"`
}

type Driver struct {
	AI              bool   `json:"isAI"`
	Id              uint8  `json:"id"`
	NetworkId       uint8  `json:"networkId"`
	TeamId          uint8  `json:"teamId"`
	IsTeamMate      bool   `json:"isTeamMate"`
	RaceNumber      uint8  `json:"raceNumber"`
	Nationality     string `json:"nationality"`
	Name            string `json:"name"`
	TelemetryStatus string `json:"telemetryStatus"`
}

type CarSetups struct {
	FrameId uint32     `json:"frameId"`
	Setups  []CarSetup `json:"carSetups"`
}

type CarSetup struct {
	FrontWing              uint8   `json:"frontWing"`
	RearWing               uint8   `json:"rearWing"`
	OnThrottle             uint8   `json:"onThrottle"`
	OffThrottle            uint8   `json:"offThrottle"`
	FrontCamber            float32 `json:"frontCamber"`
	RearCamber             float32 `json:"rearCamber"`
	FrontToe               float32 `json:"frontToe"`
	RearToe                float32 `json:"rearToe"`
	FrontSuspension        uint8   `json:"frontSuspension"`
	RearSuspension         uint8   `json:"rearSuspension"`
	FrontAntiRollBar       uint8   `json:"frontAntiRollBar"`
	RearAntiRollBar        uint8   `json:"rearAntiRollBar"`
	FrontSuspensionHeight  uint8   `json:"frontSuspensionHeight"`
	RearSuspensionHeight   uint8   `json:"rearSuspensionHeight"`
	BrakePressure          uint8   `json:"brakePressure"`
	BrakeBias              uint8   `json:"brakeBias"`
	RearLeftTyrePressure   float32 `json:"rearLeftTyrePressure"`
	RearRightTyrePressure  float32 `json:"rearRightTyrePressure"`
	FrontLeftTyrePressure  float32 `json:"frontLeftTyrePressure"`
	FrontRightTyrePressure float32 `json:"frontRightTyrePressure"`
	Ballast                uint8   `json:"ballast"`
	FuelLoad               float32 `json:"fuelLoad"`
}

type LapDatas struct {
	FrameId uint32    `json:"frameId"`
	Datas   []LapData `json:"lapDatas"`
}

type LapData struct {
	LastLapTimeInMS             uint32  `json:"lastLapTimeInMS"`
	CurrentLapTimeInMS          uint32  `json:"currentLapTimeInMS"`
	Sector1TimeInMS             uint16  `json:"sector1TimeInMS"`
	Sector2TimeInMS             uint16  `json:"sector2TimeInMS"`
	LapDistance                 float32 `json:"lapDistance"`
	TotalDistance               float32 `json:"totalDistance"`
	SafetyCarDelta              float32 `json:"safetyCarDelta"`
	CarPosition                 uint8   `json:"carPosition"`
	CurrentLapNum               uint8   `json:"currentLapNum"`
	PitStatus                   uint8   `json:"pitStatus"`
	NumPitStops                 uint8   `json:"numPitStops"`
	Sector                      uint8   `json:"sector"`
	CurrentLapInvalid           uint8   `json:"currentLapInvalid"`
	Penalties                   uint8   `json:"penalties"`
	Warnings                    uint8   `json:"warnings"`
	NumUnservedDriveThroughPens uint8   `json:"numUnservedDriveThroughPens"`
	NumUnservedStopGoPens       uint8   `json:"numUnservedStopGoPens"`
	GridPosition                uint8   `json:"gridPosition"`
	DriverStatus                uint8   `json:"driverStatus"`
	ResultStatus                uint8   `json:"resultStatus"`
	PitLaneTimerActive          uint8   `json:"pitLaneTimerActive"`
	PitLaneTimeInLaneInMS       uint16  `json:"pitLaneTimeInLaneInMS"`
	PitStopTimerInMS            uint16  `json:"pitStopTimerInMS"`
	PitStopShouldServePen       uint8   `json:"pitStopShouldServePen"`
}

type MotionData struct {
	FrameId uint32 `json:"frameId"`
	// RL, RR, FL, FR
	SuspensionPosition     [4]float32 `json:"suspensionPosition"`
	SuspensionVelocity     [4]float32 `json:"suspensionVelocity"`
	SuspensionAcceleration [4]float32 `json:"suspensionAcceleration"`
	WheelSpeed             [4]float32 `json:"wheelSpeed"`
	WheelSlip              [4]float32 `json:"wheelSlip"`

	LocalVelocityX float32 `json:"localVelocityX"`
	LocalVelocityY float32 `json:"localVelocityY"`
	LocalVelocityZ float32 `json:"localVelocityZ"`

	AngularVelocityX float32 `json:"angularVelocityX"`
	AngularVelocityY float32 `json:"angularVelocityY"`
	AngularVelocityZ float32 `json:"angularVelocityZ"`

	AngularAccelerationX float32 `json:"angularAccelerationX"`
	AngularAccelerationY float32 `json:"angularAccelerationY"`
	AngularAccelerationZ float32 `json:"angularAccelerationZ"`

	FrontWheelsAngle float32 `json:"frontWheelsAngle"`
}

type Telemetry struct {
	FrameId     uint32         `json:"frameId"`
	Telemetries []CarTelemetry `json:"telemetries"`
}

type CarTelemetry struct {
	Speed                   uint16     `json:"speed"`
	Throttle                float32    `json:"throttle"`
	Steer                   float32    `json:"steer"`
	Brake                   float32    `json:"brake"`
	Clutch                  uint8      `json:"clutch"`
	Gear                    int8       `json:"gear"`
	EngineRPM               uint16     `json:"engineRPM"`
	DRS                     bool       `json:"drs"`
	RevLightsPercent        uint8      `json:"revLightsPercent"`
	RevLightsBitValue       uint16     `json:"revLightsBitValue"`
	BrakesTemperature       [4]uint16  `json:"brakesTemperature"`
	TyresSurfaceTemperature [4]uint8   `json:"tyresSurfaceTemperature"`
	TyresInnerTemperature   [4]uint8   `json:"tyresInnerTemperature"`
	EngineTemperature       uint16     `json:"engineTemperature"`
	TyresPressure           [4]float32 `json:"tyresPressure"`
	SurfaceType             [4]uint8   `json:"surfaceType"`
}

type SessionHistory struct {
	FrameId           uint32             `json:"frameId"`
	CarId             uint8              `json:"carId"`
	NumLaps           uint8              `json:"numLaps"`
	BestLapTimeLapNum uint8              `json:"bestLapTimeLapNum"`
	BestSector1LapNum uint8              `json:"bestSector1LapNum"`
	BestSector2LapNum uint8              `json:"bestSector2LapNum"`
	BestSector3LapNum uint8              `json:"bestSector3LapNum"`
	LapHistory        []LapHistory       `json:"lapHistoryData"`
	TyreStintHistory  []TyreStintHistory `json:"tyreStintsHistory"`
}

type LapHistory struct {
	LapTimeInMS     uint32 `json:"lapTimeInMS"`
	Sector1TimeInMS uint16 `json:"sector1TimeInMS"`
	Sector2TimeInMS uint16 `json:"sector2TimeInMS"`
	Sector3TimeInMS uint16 `json:"sector3TimeInMS"`
	IsValid         bool   `json:"isValid"`
	ValidSectors    []bool `json:"validSectors"`
}

type TyreStintHistory struct {
	EndLap             uint8 `json:"m_endLap"`
	TyreActualCompound uint8 `json:"m_tyreActualCompound"`
	TyreVisualCompound uint8 `json:"m_tyreVisualCompound"`
}

type CarDamages struct {
	FrameId uint32      `json:"frameId"`
	Cars    []CarDamage `json:"cars"`
}

type CarDamage struct {
	TyresWear            [4]float32 `json:"tyresWear"`
	TyresDamage          [4]uint8   `json:"tyresDamage"`
	BrakesDamage         [4]uint8   `json:"brakesDamage"`
	FrontLeftWingDamage  uint8      `json:"frontLeftWingDamage"`
	FrontRightWingDamage uint8      `json:"frontRightWingDamage"`
	RearWingDamage       uint8      `json:"rearWingDamage"`
	FloorDamage          uint8      `json:"floorDamage"`
	DiffuserDamage       uint8      `json:"diffuserDamage"`
	SidepodDamage        uint8      `json:"sidepodDamage"`
	DRSFault             uint8      `json:"drsFault"`
	GearBoxDamage        uint8      `json:"gearBoxDamage"`
	EngineDamage         uint8      `json:"engineDamage"`
	EngineMGUHWear       uint8      `json:"engineMGUHWear"`
	EngineESWear         uint8      `json:"engineESWear"`
	EngineCEWear         uint8      `json:"engineCEWear"`
	EngineICEWear        uint8      `json:"engineICEWear"`
	EngineMGUKWear       uint8      `json:"engineMGUKWear"`
	EngineTCWear         uint8      `json:"engineTCWear"`
}
