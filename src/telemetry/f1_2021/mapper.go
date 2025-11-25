package f1_2021

import (
	"bytes"
	"f1telemetry/src/domain"
	"f1telemetry/src/telemetry/common"
)

func Map2021(envelop *common.PacketEnvelop) *common.MessageEnvelope {
	switch envelop.EventType {
	case F12021_BUTTON_STATUS:
		return &common.MessageEnvelope{
			Type: domain.BUTTON_STATUS,
			Data: mapButtonData2021(envelop.Packet.(*PacketButtons)),
		}
	case F12021_PARTICIPANTS:
		return &common.MessageEnvelope{
			Type: domain.PARTICIPANTS,
			Data: mapParticipantsData2021(envelop.Packet.(*PacketParticipantsData)),
		}
	case F12021_START_LIGHTS:
		return &common.MessageEnvelope{
			Type: domain.START_LIGHTS,
			Data: mapStartLightPacket2021(envelop.Packet.(*StartLightsPacket)),
		}
	case F12021_LIGHTS_OUT:
		return &common.MessageEnvelope{
			Type: domain.START_LIGHTS,
			Data: mapStartLightPacket2021(envelop.Packet.(*StartLightsPacket)),
		}
	case F12021_SESSION:
		return &common.MessageEnvelope{
			Type: domain.SESSION,
			Data: mapSession2021(envelop.Packet.(*PacketSessionData)),
		}
	case F12021_CARSETUP:
		return &common.MessageEnvelope{
			Type: domain.CAR_SETUPS,
			Data: mapCarSetupData2021(envelop.Packet.(*PacketCarSetupData)),
		}
	case F12021_LAP:
		return &common.MessageEnvelope{
			Type: domain.LAP,
			Data: mapLapData2021(envelop.Packet.(*PacketLapData)),
		}
	case F12021_MOTION:
		return &common.MessageEnvelope{
			Type: domain.MOTION,
			Data: mapMotionData2021(envelop.Packet.(*PacketMotionData)),
		}
	case F12021_CARTELEMETRY:
		return &common.MessageEnvelope{
			Type: domain.TELEMETRY,
			Data: mapTelemetryData2021(envelop.Packet.(*PacketCarTelemetryData)),
		}
	case F12021_SESSIONHISTORY:
		return &common.MessageEnvelope{
			Type: domain.SESSION_HISTORY,
			Data: mapSessionHistory2021(envelop.Packet.(*PacketSessionHistoryData)),
		}
	case F12021_CARDAMAGE:
		return &common.MessageEnvelope{
			Type: domain.CAR_DAMAGE,
			Data: mapCarDamage2021(envelop.Packet.(*PacketCarDamageData)),
		}
	case F12021_PENALTY_ISSUED:
		return &common.MessageEnvelope{
			Type: domain.PENALTY_ISSUED,
			Data: mapPenalty2021(envelop.Packet.(*PacketPenalty)),
		}
	default:
		return nil
	}
}

func mapButtonData2021(packet *PacketButtons) domain.Button {
	return domain.Button{
		Status:  packet.ButtonStatus,
		FrameId: packet.PacketHeader().FrameId(),
	}
}

func mapParticipantsData2021(pack *PacketParticipantsData) domain.Participants {
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
			Nationality:     countryCodeById2021(p.Nationality),
			Name:            string(nameBytes),
			TelemetryStatus: domain.TelemetryStatus(p.YourTelemetry).String()}

		drivers[i] = d
	}

	return domain.Participants{ActiveCarNum: pack.NumActiveCars, Drivers: drivers}
}

func mapStartLightPacket2021(packet *StartLightsPacket) domain.StartLight {
	var numLights = packet.NumLights
	if numLights > 5 {
		numLights = 0
	}
	return domain.StartLight{
		NumberOfLights: numLights,
	}
}

func mapSession2021(pack *PacketSessionData) domain.Session {
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
		TrackInfo:       trackInfo2021(pack),
	}
}

func trackInfo2021(pack *PacketSessionData) domain.TrackInfo {
	track := domain.Track(pack.TrackId)
	return domain.TrackInfo{
		Id:          int8(track),
		Name:        track.String(),
		TotalLaps:   pack.TotalLaps,
		TrackLength: pack.TrackLength,
	}
}

func mapCarSetupData2021(pack *PacketCarSetupData) domain.CarSetups {
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

func mapLapData2021(pack *PacketLapData) domain.LapDatas {
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

func mapMotionData2021(pack *PacketMotionData) domain.MotionData {
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

func mapTelemetryData2021(pack *PacketCarTelemetryData) domain.Telemetry {
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

func mapSessionHistory2021(pack *PacketSessionHistoryData) domain.SessionHistory {
	return domain.SessionHistory{
		CarId:             pack.CarIdx,
		NumLaps:           pack.NumLaps,
		BestLapTimeLapNum: pack.BestLapTimeLapNum,
		BestSector1LapNum: pack.BestSector1LapNum,
		BestSector2LapNum: pack.BestSector2LapNum,
		BestSector3LapNum: pack.BestSector3LapNum,
		LapHistory:        mapLapHistory2021(pack.LapHistoryData),
		TyreStintHistory:  mapTyreStintHistory(pack.TyreStintHistoryData),
	}
}

func mapLapHistory2021(lhd [100]LapHistoryData) []domain.LapHistory {
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

func mapTyreStintHistory(tsh [8]TyreStintHistoryData) []domain.TyreStintHistory {
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

func mapCarDamage2021(pack *PacketCarDamageData) domain.CarDamages {
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

func mapPenalty2021(pack *PacketPenalty) domain.Penalty {
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

func countryCodeById2021(id uint8) string {
	if int(id) >= len(countries2021) {
		id = 0
	}

	return countries2021[id]
}

var countries2021 = [88]string{
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
