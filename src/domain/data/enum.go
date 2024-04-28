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
	F1Modern  Formula = 0
	F1Classic Formula = 1
	F2        Formula = 2
)

func (s Formula) String() string {
	switch s {
	case F1Modern:
		return "F1 modern"
	case F1Classic:
		return "F1 classic"
	case F2:
		return "F2"
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
	TimeTrial          SessionType = 12
)

func (s SessionType) String() string {
	switch s {
	case Unknown:
		return "Unknown"
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

	return "N/A"
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

type Track uint8

const (
	Melbourne        Track = 0
	LeCastellet      Track = 1
	Shanghai         Track = 2
	Bahrain          Track = 3
	Catalunya        Track = 4
	Monaco           Track = 5
	Montreal         Track = 6
	Silverstone      Track = 7
	Hockenheim       Track = 8
	Hungaroring      Track = 9
	Spa              Track = 10
	Monza            Track = 11
	Singapore        Track = 12
	Suzuka           Track = 13
	AbuDhabi         Track = 14
	Texas            Track = 15
	Brazil           Track = 16
	Austria          Track = 17
	Sochi            Track = 18
	Mexico           Track = 19
	Azerbaijan       Track = 20
	SakhirShort      Track = 21
	SilverstoneShort Track = 22
	TexasShort       Track = 23
	SuzukaShort      Track = 24
	Hanoi            Track = 25
	Zandvoort        Track = 26
	Imola            Track = 27
	Portimao         Track = 28
	Jeddah           Track = 29
	Miami            Track = 30
)

func (s Track) String() string {
	var result = ""
	switch s {
	case Bahrain:
		result = "Bahrain International Circuit"
		break
	case Jeddah:
		result = "Jeddah Corniche Circuit"
		break
	case Melbourne:
		result = "Albert Park Circuit"
		break
	case Imola:
		result = "Autodromo Internazionale Enzo E Dino Ferrari"
		break
	case Miami:
		result = "Miami International Autodrome"
		break
	case Catalunya:
		result = "Circuit De Barcelona-Catalunya"
		break
	case Monaco:
		result = "Circuit De Monaco"
		break
	case Azerbaijan:
		result = "Baku City Circuit"
		break
	case Montreal:
		result = "Circuit Gilles-Villeneuve"
		break
	case Silverstone:
		result = "SilverStone Circuit"
		break
	case SilverstoneShort:
		result = "Silverstone Circuit (Short)"
		break
	case Austria:
		result = "Red Bull Ring"
		break
	case LeCastellet:
		result = "Circuit Paul Ricard"
		break
	case Hungaroring:
		result = "Hungaroring"
		break
	case Spa:
		result = "Circuit De Spa-Francorchamps"
		break
	case Zandvoort:
		result = "Circuit Zandvoort"
		break
	case Monza:
		result = "Autodromo Nazionale Monza"
		break
	case Singapore:
		result = "Marina Bay Street Circuit"
		break
	case Suzuka:
		result = "Suzuka International Racing Course"
		break
	case SuzukaShort:
		result = "Suzuka International Racing Course (Short)"
		break
	case Texas:
		result = "Circuit Of The Americas"
		break
	case TexasShort:
		result = "Circuit Of The Americas (Short)"
		break
	case Mexico:
		result = "Autódromo Hermanos Rodríguez"
		break
	case Brazil:
		result = "Autódromo José Carlos Pace"
		break
	case AbuDhabi:
		result = "Yas Marina Circuit"
		break
	case Portimao:
		result = "Algarve International Circuit"
		break
	case Shanghai:
		result = "Shanghai International Circuit"
		break
	case Hockenheim:
		result = "Hockenheim"
		break
	case Sochi:
		result = "Sochi"
		break
	case SakhirShort:
		result = "Sakhir (Short)"
		break
	case Hanoi:
		result = "Hanoi"
		break
	default:
		result = "N/A"
	}

	return result
}
