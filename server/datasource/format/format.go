package format

import (
	"log"

	"github.com/FullKernelPanic/f1-telemetry/datasource/format/f12021"
	"github.com/FullKernelPanic/f1-telemetry/listener"
)

func MapData(buf []byte, gateway listener.PacketGateway) {
	ok := f12021.MapBytes(buf, gateway)

	if !ok {
		log.Println("Unsupported packet!", buf)
	}
}
