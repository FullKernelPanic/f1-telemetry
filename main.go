package main

import (
	"bytes"
	"encoding/gob"
	"f1telemetry/src/util"
	"fmt"
	"log"
	"sync"

	"f1telemetry/src/datasource"
	"f1telemetry/src/listener"
	"f1telemetry/src/logger"
	"f1telemetry/src/web"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)

	logger.Init("app.log", true)

	server := web.CreateServer()

	l := listener.NewListener(server)

	ds := datasource.NewDataSource(l)

	go handleWebsocket(server)
	go handleDatasource(ds)

	log.Println("App ready (" + util.IpAddress().String() + ")")

	wg.Wait()
}

func EncodeToBytes(p interface{}) []byte {

	buf := bytes.Buffer{}
	enc := gob.NewEncoder(&buf)
	err := enc.Encode(p)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("uncompressed size (bytes): ", len(buf.Bytes()))
	return buf.Bytes()
}

func handleWebsocket(s *web.Server) {
	fmt.Println("start webserver")
	s.ListenAndServe(":8080")
}

func handleDatasource(ds datasource.DataSource) {
	fmt.Println("start udp")

	for {
		ds.ReadPacket()
	}
}
