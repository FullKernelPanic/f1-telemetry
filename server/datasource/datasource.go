package datasource

import (
	"log"

	"github.com/FullKernelPanic/f1-telemetry/datasource/format"
	"github.com/FullKernelPanic/f1-telemetry/listener"
)

func NewDataSource(gateway listener.PacketGateway) DataSource {
	conn, _ := CreateConnection("127.0.0.1", ":20777")

	result := DataSource{conn, gateway}

	return result
}

type DataSource struct {
	connection *Connection
	gateway    listener.PacketGateway
}

func (d *DataSource) ReadPacket() {
	buf, err := d.connection.Read()

	if err != nil {
		log.Println(err)
		return
	}

	format.MapData(buf, d.gateway)
}
