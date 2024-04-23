package f12021

import (
	"bytes"
	"fmt"
	"reflect"

	"github.com/FullKernelPanic/f1-telemetry/domain/data"
	"github.com/FullKernelPanic/f1-telemetry/listener"
)

func MapBytes(b []byte, gateway listener.PacketGateway) bool {
	header := readheader(b)

	if header.Format() != 2021 {
		return false
	}

	packet := readpacket(header.Id(), b)

	callListener(packet, gateway)

	return true
}

func callListener(p Packet, gateway listener.PacketGateway) {
	switch reflect.TypeOf(p).String() {
	case "*f12021.PacketSessionData":
		gateway.OnSession(mapSession(p.(*PacketSessionData)))
	case "*f12021.PacketButtons":
		gateway.OnButton(mapButton(p.(*PacketButtons)))
	case "*f12021.PacketParticipantsData":
		gateway.OnParticipants(mapParticipants(p.(*PacketParticipantsData)))
	case "*f12021.PacketCarSetupData":
		gateway.OnCarSetupData(mapCarSetupData(p.(*PacketCarSetupData)))
	case "*f12021.PacketLapData":
		gateway.OnLapData(mapLapData(p.(*PacketLapData)))
	case "*f12021.PacketMotionData":
		gateway.OnMotionData(mapMotionData(p.(*PacketMotionData)))
	case "*f12021.PacketCarTelemetryData":
		gateway.OnTelemetryData(mapTelemetryData(p.(*PacketCarTelemetryData)))
	case "*f12021.PacketSessionHistoryData":
		gateway.OnSessionHistory(mapSessionHistory(p.(*PacketSessionHistoryData)))
	case "*f12021.PacketCarDamageData":
		gateway.OnCarDamage(mapCarDamage(p.(*PacketCarDamageData)))
	case "*f12021.PacketPenalty":
		mapPenalty(p.(*PacketPenalty))
	//*f12021.StartLightsPacket
	//*f12021.SpeedTrapPacket
	default:
		fmt.Println(reflect.TypeOf(p).String())
	}
}

func mapSession(pack *PacketSessionData) *data.Session {
	assist := data.SettingsAssist{
		SteeringAssist:        pack.SteeringAssist,
		BrakingAssist:         data.BrakingAssist(pack.BrakingAssist).String(),
		AntiLockBrakes:        false,
		TractionControl:       0,
		RacingLine:            data.RacingLine(pack.DynamicRacingLine).String(),
		DynamicRacingLineType: pack.DynamicRacingLineType,
		GearboxAssist:         pack.GearboxAssist,
		PitAssist:             pack.PitAssist,
		PitReleaseAssist:      pack.PitReleaseAssist,
		DRSAssist:             pack.DRSAssist,
		ERSAssist:             pack.ERSAssist}

	env := data.Environment{
		AirTemperature:   pack.AirTemperature,
		TrackTemperature: pack.TrackTemperature,
		Weather:          data.Weather(pack.Weather).String()}

	return &data.Session{
		FrameId:         pack.Header.FrameIdentifier,
		Assist:          assist,
		Environment:     env,
		AIDifficulty:    pack.AIDifficulty,
		SessionType:     data.SessionType(pack.SessionType).String(),
		Formula:         data.Formula(pack.Formula).String(),
		SessionTimeLeft: pack.SessionTimeLeft,
		SessionDuration: pack.SessionDuration,
		GamePaused:      pack.GamePaused}
}

func mapButton(pack *PacketButtons) *data.Button {
	return &data.Button{Status: pack.ButtonStatus}
}

func mapParticipants(pack *PacketParticipantsData) *data.Participants {
	drivers := make([]data.Driver, pack.NumActiveCars)

	for i := 0; i < int(pack.NumActiveCars); i++ {
		p := pack.ParticipantData[i]

		nameBytes := bytes.Trim(p.Name[:], "\x00")
		d := data.Driver{
			AI:              p.AIControlled,
			Id:              p.DriverId,
			NetworkId:       p.NetworkId,
			TeamId:          p.TeamId,
			IsTeamMate:      p.MyTeam,
			RaceNumber:      p.RaceNumber,
			Nationality:     countryCodeById(p.Nationality),
			Name:            string(nameBytes),
			TelemetryStatus: data.TelemetryStatus(p.YourTelemetry).String()}

		drivers[i] = d
	}

	return &data.Participants{ActiveCarNum: pack.NumActiveCars, Drivers: drivers}
}

func mapCarSetupData(pack *PacketCarSetupData) *data.CarSetups {
	setups := make([]data.CarSetup, len(pack.CarSetupData))

	for i, _ := range pack.CarSetupData {
		setup := data.CarSetup{
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

	return &data.CarSetups{Setups: setups}
}

func mapLapData(pack *PacketLapData) *data.LapDatas {
	datas := make([]data.LapData, len(pack.LapData))

	for i, ld := range pack.LapData {
		data := data.LapData{
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

		datas[i] = data
	}

	return &data.LapDatas{Datas: datas}
}

func mapMotionData(pack *PacketMotionData) *data.MotionData {
	return &data.MotionData{
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

func mapTelemetryData(pack *PacketCarTelemetryData) *data.Telemetry {
	datas := make([]data.CarTelemetry, len(pack.CarTelemetryData))

	for i, ct := range pack.CarTelemetryData {
		data := data.CarTelemetry{
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

		datas[i] = data
	}

	return &data.Telemetry{
		FrameId:     pack.Header.FrameIdentifier,
		Telemetries: datas}
}

func mapSessionHistory(pack *PacketSessionHistoryData) *data.SessionHistory {
	return &data.SessionHistory{
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

func mapLapHistory(lhd [100]LapHistoryData) []data.LapHistory {
	result := make([]data.LapHistory, 0)

	for _, lh := range lhd {
		isValid := false
		varlidSectors := make([]bool, 0)

		result = append(result, data.LapHistory{
			LapTimeInMS:     lh.LapTimeInMS,
			Sector1TimeInMS: lh.Sector1TimeInMS,
			Sector2TimeInMS: lh.Sector2TimeInMS,
			Sector3TimeInMS: lh.Sector3TimeInMS,
			IsValid:         isValid,
			ValidSectors:    varlidSectors,
		})
	}

	return result
}

func mapTyreStintHistory(tsh [8]TyreStintHistoryData) []data.TyreStintHistory {
	result := make([]data.TyreStintHistory, 0)

	for _, t := range tsh {
		result = append(result, data.TyreStintHistory{
			EndLap:             t.EndLap,
			TyreActualCompound: t.TyreActualCompound,
			TyreVisualCompound: t.TyreVisualCompound,
		})
	}

	return result
}

func mapCarDamage(pack *PacketCarDamageData) data.CarDamages {
	cars := make([]data.CarDamage, len(pack.CarDamageData))

	for i, cd := range pack.CarDamageData {
		car := data.CarDamage{
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

	return data.CarDamages{Cars: cars}
}

func mapPenalty(pack *PacketPenalty) {
	fmt.Println(pack.InfringementType)
}
