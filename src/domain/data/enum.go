package data

type BrakingAssist uint8

const (
	Off    BrakingAssist = 0
	Low    BrakingAssist = 1
	Medium BrakingAssist = 2
	High   BrakingAssist = 3
)

func (s BrakingAssist) String() string {
	switch s {
	case Off:
		return "Off"
	case Low:
		return "low"
	case Medium:
		return "medium"
	case High:
		return "high"
	}
	return "unknown"
}

type Weather uint8

const (
	Clear      Weather = 0
	LightCloud Weather = 1
	Overcast   Weather = 2
	LightRain  Weather = 3
	HeavyRain  Weather = 4
	Storm      Weather = 5
)

func (s Weather) String() string {
	switch s {
	case Clear:
		return "clear"
	case LightCloud:
		return "light cloud"
	case Overcast:
		return "overcast"
	case LightRain:
		return "light rain"
	case HeavyRain:
		return "heavy rain"
	case Storm:
		return "storm"
	}
	return "unknown"
}

type Formula int

const (
	Modern  Formula = 0
	Classic Formula = 1
	F2      Formula = 2
	F1      Formula = 3
)

func (s Formula) String() string {
	switch s {
	case Modern:
		return "F1 modern"
	case Classic:
		return "F1 classic"
	case F2:
		return "F2"
	case F1:
		return "F1 generic"
	}

	return "unknown"
}

type RacingLine int

const (
	None   RacingLine = 0
	Corner RacingLine = 1
	Full   RacingLine = 2
)

func (s RacingLine) String() string {
	switch s {
	case None:
		return "Off"
	case Corner:
		return "Corner"
	case Full:
		return "Full"
	}

	return "unknown"
}

type SessionType int

const (
	Unknown            SessionType = 0
	P1                 SessionType = 1
	P2                 SessionType = 2
	P3                 SessionType = 3
	ShortPractice      SessionType = 4
	Q1                 SessionType = 5
	Q2                 SessionType = 6
	Q3                 SessionType = 7
	ShortQualification SessionType = 8
	OSQ                SessionType = 9
	R                  SessionType = 10
	R2                 SessionType = 11
	TimeTrial          SessionType = 13
)

func (s SessionType) String() string {
	switch s {
	case P1:
		return "Practice 1"
	case P2:
		return "Practice 2"
	case P3:
		return "Practice 3"
	case ShortPractice:
		return "Short Practice"
	case Q1:
		return "Q1"
	case Q2:
		return "Q2"
	case Q3:
		return "Q3"
	case ShortQualification:
		return "Short Qualification"
	case OSQ:
		return "One Shot Qualification"
	case R:
		return "Race"
	case R2:
		return "Sprint race maybe?"
	case TimeTrial:
		return "Time Trial"
	}

	return "unknown"
}

type TelemetryStatus uint8

const (
	Restricted TelemetryStatus = 0
	Public     TelemetryStatus = 1
)

func (s TelemetryStatus) String() string {
	switch s {
	case Restricted:
		return "restricted"
	case Public:
		return "public"
	}

	return "unknown"
}
