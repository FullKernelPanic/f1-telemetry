package datasource

import (
	"fmt"
	"log"
)

func NewDataSource(port int) UDPTelemetry {
	conn, _ := CreateConnection(fmt.Sprintf(":%v", port))

	result := UDPTelemetry{conn}

	return result
}

type UDPTelemetry struct {
	connection *Connection
}

func (d *UDPTelemetry) ReadPacket() []byte {
	buf, err := d.connection.Read()

	if err != nil {
		log.Println(err)
		return make([]byte, 0)
	}

	return buf
}
