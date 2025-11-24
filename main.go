package main

import (
	"bytes"
	"encoding/gob"
	"f1telemetry/src/datasource"
	"f1telemetry/src/domain"
	"f1telemetry/src/listener"
	"f1telemetry/src/logger"
	"f1telemetry/src/telemetry"
	"f1telemetry/src/telemetry/versions"
	"f1telemetry/src/util"
	"f1telemetry/src/web"
	"fmt"
	"log"
	"sync"
)

const (
	inBufSize  = 65536 // tune to expected bursts
	outBufSize = 65536 // tune to expected bursts
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	logger.Init("app.log", true)

	server := web.CreateServer()
	inChannel := make(chan []byte, inBufSize)

	ds := datasource.NewDataSource(20777)

	gob.Register(domain.Button{})
	gob.Register(domain.Session{})
	gob.Register(domain.Participants{})
	gob.Register(domain.CarSetups{})
	gob.Register(domain.CarDamages{})
	gob.Register(domain.SessionHistory{})
	gob.Register(domain.LapDatas{})
	gob.Register(domain.MotionData{})
	gob.Register(domain.Telemetry{})
	gob.Register(domain.StartLight{})
	gob.Register(domain.Penalty{})

	go handleDatasource(ds, inChannel)
	go handleWebsocket(listener.NewListener(server), server, inChannel)

	log.Println("App ready (" + util.IpAddress().String() + ")")

	wg.Wait()
}

func handleWebsocket(l listener.Listener, s *web.Server, dataChan chan []byte) {
	fmt.Println("start webserver")
	go s.ListenAndServe(":8080")

	for {
		data := <-dataChan
		envelope := versions.MessageEnvelope{}
		DecodeFromBytes(data, &envelope)

		switch envelope.Type {
		case domain.SESSION:
			l.OnSession(envelope.Data.(domain.Session))
			break
		case domain.BUTTON_STATUS:
			l.OnButton(envelope.Data.(domain.Button))
			break
		case domain.PARTICIPANTS:
			l.OnParticipants(envelope.Data.(domain.Participants))
			break
		case domain.CAR_SETUPS:
			l.OnCarSetupData(envelope.Data.(domain.CarSetups))
			break
		case domain.LAP:
			l.OnLapData(envelope.Data.(domain.LapDatas))
			break
		case domain.MOTION:
			l.OnMotionData(envelope.Data.(domain.MotionData))
			break
		case domain.TELEMETRY:
			l.OnTelemetryData(envelope.Data.(domain.Telemetry))
			break
		case domain.CAR_DAMAGE:
			l.OnCarDamage(envelope.Data.(domain.CarDamages))
			break
		case domain.START_LIGHTS:
			l.OnStartLight(envelope.Data.(domain.StartLight))
			break
		case domain.SESSION_HISTORY:
			l.OnSessionHistory(envelope.Data.(domain.SessionHistory))
			break
		default:
			fmt.Println(fmt.Sprintf("Unhandled >> %v", envelope.Type))
		}
	}
}

func handleDatasource(ds datasource.UDPTelemetry, dataChan chan []byte) {
	fmt.Println("start udp")

	for {
		buff := ds.ReadPacket()
		p := telemetry.MapData(buff)
		if p != nil && p.Data != nil {
			fmt.Println(p.Type)
			go func() {
				defer func() {
				}()
				dataChan <- EncodeToBytes(p)
			}()
			fmt.Println("channel sent")
		}
	}
}

func EncodeToBytes(p *versions.MessageEnvelope) []byte {
	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	return buf.Bytes()
}

func DecodeFromBytes(data []byte, result interface{}) {
	buf := bytes.NewBuffer(data)
	dec := gob.NewDecoder(buf)
	dec.Decode(result)
}
