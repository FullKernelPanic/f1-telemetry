package main

import (
	"log"
	"sync"

	"github.com/FullKernelPanic/f1-telemetry/datasource"
	"github.com/FullKernelPanic/f1-telemetry/listener"
	"github.com/FullKernelPanic/f1-telemetry/logger"
	"github.com/FullKernelPanic/f1-telemetry/web"
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
