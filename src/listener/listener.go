package listener

import (
	"encoding/json"
	"f1telemetry/src/web"
	"fmt"
	"log"
	"reflect"

	"f1telemetry/src/domain"
)

type JsonDataFormat struct {
	Type string      `json:"type"`
	Data interface{} `json:"data"`
}

func NewListener(b web.Broadcaster) Listener {
	return Listener{b}
}

type Listener struct {
	broadcaster web.Broadcaster
}

func (l Listener) OnSession(data domain.Session) {
	l.broadcast("session", data)
}

func (l Listener) OnButton(data domain.Button) {
	l.broadcast("button", data)
}

func (l Listener) OnParticipants(data domain.Participants) {
	l.broadcast("participants", data)
}

func (l Listener) OnCarSetupData(data domain.CarSetups) {
	l.broadcast("carSetup", data)
}

func (l Listener) OnLapData(data domain.LapDatas) {
	l.broadcast("lapData", data)
}

func (l Listener) OnMotionData(data domain.MotionData) {
	l.broadcast("motion", data)
}

func (l Listener) OnTelemetryData(data domain.Telemetry) {
	l.broadcast("telemetry", data)
}

func (l Listener) OnSessionHistory(data domain.SessionHistory) {
	l.broadcast("sessionHistory", data)
}

func (l Listener) OnCarDamage(data domain.CarDamages) {
	l.broadcast("carDamage", data)
}

func (l Listener) OnStartLight(data domain.StartLight) {
	l.broadcast("startLight", data)
}

func (l Listener) OnLobbyInfoData(data domain.LobbyInfo) {
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
	OnSession(data domain.Session)
	OnButton(data domain.Button)
	OnParticipants(data domain.Participants)
	OnCarSetupData(data domain.CarSetups)
	OnLapData(data domain.LapDatas)
	OnMotionData(data domain.MotionData)
	OnTelemetryData(data domain.Telemetry)
	OnSessionHistory(data domain.SessionHistory)
	OnCarDamage(data domain.CarDamages)
	OnStartLight(packet domain.StartLight)
	OnLobbyInfoData(packet domain.LobbyInfo)
}
