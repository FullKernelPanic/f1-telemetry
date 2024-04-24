package format

import (
	"log"

	"f1telemetry/src/datasource/format/f12021"
	"f1telemetry/src/listener"
)

func MapData(buf []byte, gateway listener.PacketGateway) {
	ok := f12021.MapBytes(buf, gateway)

	if !ok {
		log.Println("Unsupported packet!", buf)
	}
}
