package main

import (
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

	log.Println("App ready")

	wg.Wait()
}

func handleWebsocket(s *web.Server) {
	s.ListenAndServe(":8080")
}

func handleDatasource(ds datasource.DataSource) {
	for {
		ds.ReadPacket()
	}
}
