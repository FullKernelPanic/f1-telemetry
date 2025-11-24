package mapper

import (
	"bytes"
	"f1telemetry/src/domain"
	"f1telemetry/src/telemetry/versions"
	"f1telemetry/src/telemetry/versions/f1_2021"
)

func Map2021(envelop *versions.PacketEnvelop) *versions.MessageEnvelope {
	switch envelop.EventType {
	case f1_2021.F12021_BUTTON_STATUS:
		return &versions.MessageEnvelope{
			Type: domain.BUTTON_STATUS,
			Data: mapButtonData2021(envelop.Packet.(*f1_2021.PacketButtons)),
		}
	case f1_2021.F12021_PARTICIPANTS:
		return &versions.MessageEnvelope{
			Type: domain.PARTICIPANTS,
			Data: mapParticipantsData2021(envelop.Packet.(*f1_2021.PacketParticipantsData)),
		}
	case f1_2021.F12021_START_LIGHTS:
		return &versions.MessageEnvelope{
			Type: domain.START_LIGHTS,
			Data: mapStartLightPacket(envelop.Packet.(*f1_2021.StartLightsPacket)),
		}
	case f1_2021.F12021_LIGHTS_OUT:
		return &versions.MessageEnvelope{
			Type: domain.START_LIGHTS,
			Data: mapStartLightPacket(envelop.Packet.(*f1_2021.StartLightsPacket)),
		}
	case f1_2021.F12021_SESSION:
		return &versions.MessageEnvelope{
			Type: domain.SESSION,
			Data: mapSession(envelop.Packet.(*f1_2021.PacketSessionData)),
		}
	case f1_2021.F12021_CARSETUP:
		return &versions.MessageEnvelope{
			Type: domain.CAR_SETUPS,
			Data: mapCarSetupData(envelop.Packet.(*f1_2021.PacketCarSetupData)),
		}
	case f1_2021.F12021_LAP:
		return &versions.MessageEnvelope{
			Type: domain.LAP,
			Data: mapLapData(envelop.Packet.(*f1_2021.PacketLapData)),
		}
	case f1_2021.F12021_MOTION:
		return &versions.MessageEnvelope{
			Type: domain.MOTION,
			Data: mapMotionData(envelop.Packet.(*f1_2021.PacketMotionData)),
		}
	case f1_2021.F12021_CARTELEMETRY:
		return &versions.MessageEnvelope{
			Type: domain.TELEMETRY,
			Data: mapTelemetryData(envelop.Packet.(*f1_2021.PacketCarTelemetryData)),
		}
	case f1_2021.F12021_SESSIONHISTORY:
		return &versions.MessageEnvelope{
			Type: domain.SESSION_HISTORY,
			Data: mapSessionHistory(envelop.Packet.(*f1_2021.PacketSessionHistoryData)),
		}
	case f1_2021.F12021_CARDAMAGE:
		return &versions.MessageEnvelope{
			Type: domain.CAR_DAMAGE,
			Data: mapCarDamage(envelop.Packet.(*f1_2021.PacketCarDamageData)),
		}
	case f1_2021.F12021_PENALTY_ISSUED:
		return &versions.MessageEnvelope{
			Type: domain.PENALTY_ISSUED,
			Data: mapPenalty(envelop.Packet.(*f1_2021.PacketPenalty)),
		}
	default:
		return nil
	}
}

func mapButtonData2021(packet *f1_2021.PacketButtons) domain.Button {
	return domain.Button{
		Status:  packet.ButtonStatus,
		FrameId: packet.PacketHeader().FrameId(),
	}
}

func mapParticipantsData2021(pack *f1_2021.PacketParticipantsData) domain.Participants {
	drivers := make([]domain.Driver, pack.NumActiveCars)

	for i := 0; i < int(pack.NumActiveCars); i++ {
		p := pack.ParticipantData[i]

		nameBytes := bytes.Trim(p.Name[:], "\x00")
		d := domain.Driver{
			AI:              p.AIControlled,
			Id:              p.DriverId,
			NetworkId:       p.NetworkId,
			TeamId:          p.TeamId,
			IsTeamMate:      p.MyTeam,
			RaceNumber:      p.RaceNumber,
			Nationality:     countryCodeById(p.Nationality),
			Name:            string(nameBytes),
			TelemetryStatus: domain.TelemetryStatus(p.YourTelemetry).String()}

		drivers[i] = d
	}

	return domain.Participants{ActiveCarNum: pack.NumActiveCars, Drivers: drivers}
}

func mapStartLightPacket(packet *f1_2021.StartLightsPacket) domain.StartLight {
	var numLights = packet.NumLights
	if numLights > 5 {
		numLights = 0
	}
	return domain.StartLight{
		NumberOfLights: numLights,
	}
}

func mapSession(pack *f1_2021.PacketSessionData) domain.Session {
	assist := domain.SettingsAssist{
		SteeringAssist:        pack.SteeringAssist,
		BrakingAssist:         domain.BrakingAssist(pack.BrakingAssist).String(),
		AntiLockBrakes:        false,
		TractionControl:       0,
		RacingLine:            domain.RacingLine(pack.DynamicRacingLine).String(),
		DynamicRacingLineType: pack.DynamicRacingLineType,
		GearboxAssist:         pack.GearboxAssist,
		PitAssist:             pack.PitAssist,
		PitReleaseAssist:      pack.PitReleaseAssist,
		DRSAssist:             pack.DRSAssist,
		ERSAssist:             pack.ERSAssist}

	env := domain.Environment{
		AirTemperature:   pack.AirTemperature,
		TrackTemperature: pack.TrackTemperature,
		Weather:          domain.Weather(pack.Weather).String()}

	return domain.Session{
		FrameId:         pack.Header.FrameIdentifier,
		Assist:          assist,
		Environment:     env,
		AIDifficulty:    pack.AIDifficulty,
		SessionType:     domain.SessionType(pack.SessionType).String(),
		Formula:         domain.Formula(pack.Formula).String(),
		SessionTimeLeft: pack.SessionTimeLeft,
		SessionDuration: pack.SessionDuration,
		GamePaused:      pack.GamePaused,
		TrackInfo:       trackInfo(pack),
	}
}

func trackInfo(pack *f1_2021.PacketSessionData) domain.TrackInfo {
	track := domain.Track(pack.TrackId)
	return domain.TrackInfo{
		Id:          int8(track),
		Name:        track.String(),
		TotalLaps:   pack.TotalLaps,
		TrackLength: pack.TrackLength,
	}
}

func mapCarSetupData(pack *f1_2021.PacketCarSetupData) domain.CarSetups {
	setups := make([]domain.CarSetup, len(pack.CarSetupData))

	for i, _ := range pack.CarSetupData {
		setup := domain.CarSetup{
			/*FrontWing:              cs.FrontWing,
			RearWing:               cs.RearWing,
			OnThrottle:             cs.OnThrottle,
			OffThrottle:            cs.OffThrottle,
			FrontCamber:            cs.FrontCamber,
			RearCamber:             cs.RearCamber,
			FrontToe:               cs.FrontToe,
			RearToe:                cs.RearToe,
			FrontSuspension:        cs.FrontSuspension,
			RearSuspension:         cs.RearSuspension,
			FrontAntiRollBar:       cs.FrontAntiRollBar,
			RearAntiRollBar:        cs.RearAntiRollBar,
			FrontSuspensionHeight:  cs.FrontSuspensionHeight,
			RearSuspensionHeight:   cs.RearSuspensionHeight,
			BrakePressure:          cs.BrakePressure,
			BrakeBias:              cs.BrakeBias,
			RearLeftTyrePressure:   cs.RearLeftTyrePressure,
			RearRightTyrePressure:  cs.RearRightTyrePressure,
			FrontLeftTyrePressure:  cs.FrontLeftTyrePressure,
			FrontRightTyrePressure: cs.FrontRightTyrePressure,
			Ballast:                cs.Ballast,
			FuelLoad:               cs.FuelLoad*/}

		setups[i] = setup
	}

	return domain.CarSetups{Setups: setups}
}

func mapLapData(pack *f1_2021.PacketLapData) domain.LapDatas {
	datas := make([]domain.LapData, len(pack.LapData))

	for i, ld := range pack.LapData {
		dld := domain.LapData{
			CurrentLapTimeInMS:          ld.CurrentLapTimeInMS,
			LastLapTimeInMS:             ld.LastLapTimeInMS,
			Sector1TimeInMS:             ld.Sector1TimeInMS,
			Sector2TimeInMS:             ld.Sector2TimeInMS,
			LapDistance:                 ld.LapDistance,
			TotalDistance:               ld.TotalDistance,
			SafetyCarDelta:              ld.SafetyCarDelta,
			CarPosition:                 ld.CarPosition,
			CurrentLapNum:               ld.CurrentLapNum,
			PitStatus:                   ld.PitStatus,
			NumPitStops:                 ld.NumPitStops,
			Sector:                      ld.Sector,
			CurrentLapInvalid:           ld.CurrentLapInvalid,
			Penalties:                   ld.Penalties,
			Warnings:                    ld.Warnings,
			NumUnservedDriveThroughPens: ld.NumUnservedDriveThroughPens,
			NumUnservedStopGoPens:       ld.NumUnservedStopGoPens,
			GridPosition:                ld.GridPosition,
			DriverStatus:                ld.DriverStatus,
			ResultStatus:                ld.ResultStatus,
			PitLaneTimerActive:          ld.PitLaneTimerActive,
			PitLaneTimeInLaneInMS:       ld.PitLaneTimeInLaneInMS,
			PitStopTimerInMS:            ld.PitStopTimerInMS,
			PitStopShouldServePen:       ld.PitStopShouldServePen}

		datas[i] = dld
	}

	return domain.LapDatas{FrameId: pack.Header.FrameIdentifier, Datas: datas}
}

func mapMotionData(pack *f1_2021.PacketMotionData) domain.MotionData {
	return domain.MotionData{
		SuspensionPosition:     pack.SuspensionPosition,
		SuspensionVelocity:     pack.SuspensionVelocity,
		SuspensionAcceleration: pack.SuspensionAcceleration,
		WheelSpeed:             pack.WheelSpeed,
		WheelSlip:              pack.WheelSlip,
		LocalVelocityX:         pack.LocalVelocityX,
		LocalVelocityY:         pack.LocalVelocityY,
		LocalVelocityZ:         pack.LocalVelocityZ,
		AngularVelocityX:       pack.AngularVelocityX,
		AngularVelocityY:       pack.AngularVelocityY,
		AngularVelocityZ:       pack.AngularVelocityZ,
		AngularAccelerationX:   pack.AngularAccelerationX,
		AngularAccelerationY:   pack.AngularAccelerationY,
		AngularAccelerationZ:   pack.AngularAccelerationZ}
}

func mapTelemetryData(pack *f1_2021.PacketCarTelemetryData) domain.Telemetry {
	datas := make([]domain.CarTelemetry, len(pack.CarTelemetryData))

	for i, ct := range pack.CarTelemetryData {
		dct := domain.CarTelemetry{
			Speed:                   ct.Speed,
			Throttle:                ct.Throttle,
			Steer:                   ct.Steer,
			Brake:                   ct.Brake,
			Clutch:                  ct.Clutch,
			Gear:                    ct.Gear,
			EngineRPM:               ct.EngineRPM,
			DRS:                     ct.DRS,
			RevLightsPercent:        ct.RevLightsPercent,
			RevLightsBitValue:       ct.RevLightsBitValue,
			BrakesTemperature:       ct.BrakesTemperature,
			TyresSurfaceTemperature: ct.TyresSurfaceTemperature,
			TyresInnerTemperature:   ct.TyresInnerTemperature,
			EngineTemperature:       ct.EngineTemperature,
			TyresPressure:           ct.TyresPressure,
			SurfaceType:             ct.SurfaceType}

		datas[i] = dct
	}

	return domain.Telemetry{
		FrameId:     pack.Header.FrameIdentifier,
		Telemetries: datas}
}

func mapSessionHistory(pack *f1_2021.PacketSessionHistoryData) domain.SessionHistory {
	return domain.SessionHistory{
		CarId:             pack.CarIdx,
		NumLaps:           pack.NumLaps,
		BestLapTimeLapNum: pack.BestLapTimeLapNum,
		BestSector1LapNum: pack.BestSector1LapNum,
		BestSector2LapNum: pack.BestSector2LapNum,
		BestSector3LapNum: pack.BestSector3LapNum,
		LapHistory:        mapLapHistory(pack.LapHistoryData),
		TyreStintHistory:  mapTyreStintHistory(pack.TyreStintHistoryData),
	}
}

func mapLapHistory(lhd [100]f1_2021.LapHistoryData) []domain.LapHistory {
	result := make([]domain.LapHistory, 0)

	for _, lh := range lhd {
		varlidSectors := make([]bool, 0)

		result = append(result, domain.LapHistory{
			LapTimeInMS:     lh.LapTimeInMS,
			Sector1TimeInMS: lh.Sector1TimeInMS,
			Sector2TimeInMS: lh.Sector2TimeInMS,
			Sector3TimeInMS: lh.Sector3TimeInMS,
			IsValid:         false,
			ValidSectors:    varlidSectors,
		})
	}

	return result
}

func mapTyreStintHistory(tsh [8]f1_2021.TyreStintHistoryData) []domain.TyreStintHistory {
	result := make([]domain.TyreStintHistory, 0)

	for _, t := range tsh {
		result = append(result, domain.TyreStintHistory{
			EndLap:             t.EndLap,
			TyreActualCompound: t.TyreActualCompound,
			TyreVisualCompound: t.TyreVisualCompound,
		})
	}

	return result
}

func mapCarDamage(pack *f1_2021.PacketCarDamageData) domain.CarDamages {
	cars := make([]domain.CarDamage, len(pack.CarDamageData))

	for i, cd := range pack.CarDamageData {
		car := domain.CarDamage{
			TyresWear:            cd.TyresWear,
			TyresDamage:          cd.TyresDamage,
			BrakesDamage:         cd.BrakesDamage,
			FrontLeftWingDamage:  cd.FrontLeftWingDamage,
			FrontRightWingDamage: cd.FrontRightWingDamage,
			RearWingDamage:       cd.RearWingDamage,
			FloorDamage:          cd.FloorDamage,
			DiffuserDamage:       cd.DiffuserDamage,
			SidepodDamage:        cd.SidepodDamage,
			DRSFault:             cd.DRSFault,
			GearBoxDamage:        cd.GearBoxDamage,
			EngineDamage:         cd.EngineDamage,
			EngineMGUHWear:       cd.EngineMGUHWear,
			EngineESWear:         cd.EngineESWear,
			EngineCEWear:         cd.EngineCEWear,
			EngineICEWear:        cd.EngineICEWear,
			EngineMGUKWear:       cd.EngineMGUKWear,
			EngineTCWear:         cd.EngineTCWear}
		cars[i] = car
	}

	return domain.CarDamages{Cars: cars}
}

func mapPenalty(pack *f1_2021.PacketPenalty) domain.Penalty {
	return domain.Penalty{
		PenaltyType:      pack.PenaltyType,
		InfringementType: pack.InfringementType,
		VehicleIndex:     pack.VehicleIndex,
		OtherVehicleIdx:  pack.OtherVehicleIdx,
		Time:             pack.Time,
		LapNum:           pack.LapNum,
		PlacesGained:     pack.PlacesGained,
	}
}

func countryCodeById(id uint8) string {
	if int(id) >= len(countries) {
		id = 0
	}

	return countries[id]
}

var countries = [88]string{
	"N/A",
	"American",
	"Argentinean",
	"Australian",
	"Austrian",
	"Azerbaijani",
	"Bahraini",
	"Belgian",
	"Bolivian",
	"Brazilian",
	"British",
	"Bulgarian",
	"Cameroonian",
	"Canadian",
	"Chilean",
	"Chinese",
	"Colombian",
	"CostaRican",
	"Croatian",
	"Cypriot",
	"Czech",
	"Danish",
	"Dutch",
	"Ecuadorian",
	"English",
	"Emirian",
	"Estonian",
	"Finnish",
	"French",
	"German",
	"Ghanaian",
	"Greek",
	"Guatemalan",
	"Honduran",
	"HongKonger",
	"Hungarian",
	"Icelander",
	"Indian",
	"Indonesian",
	"Irish",
	"Israeli",
	"Italian",
	"Jamaican",
	"Japanese",
	"Jordanian",
	"Kuwaiti",
	"Latvian",
	"Lebanese",
	"Lithuanian",
	"Luxembourger",
	"Malaysian",
	"Maltese",
	"Mexican",
	"Monegasque",
	"NewZealander",
	"Nicaraguan",
	"NorthernIrish",
	"Norwegian",
	"Omani",
	"Pakistani",
	"Panamanian",
	"Paraguayan",
	"Peruvian",
	"Polish",
	"Portuguese",
	"Qatari",
	"Romanian",
	"Russian",
	"Salvadoran",
	"Saudi",
	"Scottish",
	"Serbian",
	"Singaporean",
	"Slovakian",
	"Slovenian",
	"SouthKorean",
	"SouthAfrican",
	"Spanish",
	"Swedish",
	"Swiss",
	"Thai",
	"Turkish",
	"Uruguayan",
	"Ukrainian",
	"Venezuelan",
	"Welsh",
	"Barbadian",
	"Vietnamese",
}
