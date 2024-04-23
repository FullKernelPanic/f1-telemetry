package listener

import (
	"encoding/json"
	"log"
	"reflect"

	"github.com/FullKernelPanic/f1-telemetry/domain"
	"github.com/FullKernelPanic/f1-telemetry/domain/data"
)

func NewListener(b domain.Broadcaster) listener {
	return listener{b}
}

type listener struct {
	broadcaster domain.Broadcaster
}

func (l listener) OnSession(data *data.Session) {
	l.broadcast("session", data)
}

func (l listener) OnButton(data *data.Button) {
	l.broadcast("button", data)
}

func (l listener) OnParticipants(data *data.Participants) {
	l.broadcast("participants", data)
}

func (l listener) OnCarSetupData(data *data.CarSetups) {
	l.broadcast("carSetup", data)
}

func (l listener) OnLapData(data *data.LapDatas) {
	l.broadcast("lapData", data)
}

func (l listener) OnMotionData(data *data.MotionData) {
	l.broadcast("motion", data)
}

func (l listener) OnTelemetryData(data *data.Telemetry) {
	l.broadcast("telemetry", data)
}

func (l listener) OnSessionHistory(data *data.SessionHistory) {
	l.broadcast("sessionHistory", data)
}

func (l listener) OnCarDamage(data data.CarDamages) {
	l.broadcast("carDamage", data)
}

func (l listener) broadcast(messageType string, data interface{}) {
	message, err := json.Marshal(JsonDataFormat{messageType, data})

	if err != nil {
		log.Println("Json error >> "+err.Error(), reflect.TypeOf(data).String())
		return
	}

	l.broadcaster.Broadcast(message)
}

type PacketGateway interface {
	OnSession(data *data.Session)
	OnButton(data *data.Button)
	OnParticipants(data *data.Participants)
	OnCarSetupData(data *data.CarSetups)
	OnLapData(data *data.LapDatas)
	OnMotionData(data *data.MotionData)
	OnTelemetryData(data *data.Telemetry)
	OnSessionHistory(data *data.SessionHistory)
	OnCarDamage(data data.CarDamages)
}
