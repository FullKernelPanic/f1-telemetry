package listener

import (
	"encoding/json"
	"fmt"
	"log"
	"reflect"

	"f1telemetry/src/domain"
	"f1telemetry/src/domain/data"
)

func NewListener(b domain.Broadcaster) Listener {
	return Listener{b}
}

type Listener struct {
	broadcaster domain.Broadcaster
}

func (l Listener) OnSession(data data.Session) {
	l.broadcast("session", data)
}

func (l Listener) OnButton(data data.Button) {
	l.broadcast("button", data)
}

func (l Listener) OnParticipants(data data.Participants) {
	l.broadcast("participants", data)
}

func (l Listener) OnCarSetupData(data data.CarSetups) {
	l.broadcast("carSetup", data)
}

func (l Listener) OnLapData(data data.LapDatas) {
	l.broadcast("lapData", data)
}

func (l Listener) OnMotionData(data data.MotionData) {
	l.broadcast("motion", data)
}

func (l Listener) OnTelemetryData(data data.Telemetry) {
	l.broadcast("telemetry", data)
}

func (l Listener) OnSessionHistory(data data.SessionHistory) {
	l.broadcast("sessionHistory", data)
}

func (l Listener) OnCarDamage(data data.CarDamages) {
	l.broadcast("carDamage", data)
}

func (l Listener) OnStartLight(data data.StartLight) {
	fmt.Println("OnStartLight")
	l.broadcast("startLight", data)
}

func (l Listener) OnLobbyInfoData(data data.LobbyInfo) {
	fmt.Println("OnLobbyInfoData")
	l.broadcast("lobbyInfo", data)
}

func (l Listener) broadcast(messageType string, data interface{}) {
	message, err := json.Marshal(JsonDataFormat{messageType, data})

	if err != nil {
		log.Println("Json error >> "+err.Error(), reflect.TypeOf(data).String())
		return
	}

	l.broadcaster.Broadcast(message)
}

type PacketGateway interface {
	OnSession(data data.Session)
	OnButton(data data.Button)
	OnParticipants(data data.Participants)
	OnCarSetupData(data data.CarSetups)
	OnLapData(data data.LapDatas)
	OnMotionData(data data.MotionData)
	OnTelemetryData(data data.Telemetry)
	OnSessionHistory(data data.SessionHistory)
	OnCarDamage(data data.CarDamages)
	OnStartLight(packet data.StartLight)
	OnLobbyInfoData(packet data.LobbyInfo)
}
