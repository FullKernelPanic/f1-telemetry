package datasource

import (
	"net"
)

type Connection struct {
	connection *net.UDPConn
}

func (c *Connection) Read() ([]byte, error) {
	buf := make([]byte, 1024+1024/2)

	_, _, err := c.connection.ReadFromUDP(buf)

	if err != nil {
		return nil, err
	}

	return buf, nil
}

func CreateConnection(ip string, port string) (*Connection, error) {
	protocol := "udp"

	addr, err := net.ResolveUDPAddr(protocol, port)

	if err != nil {
		return nil, err
	}

	conn, err := net.ListenUDP(protocol, addr)

	if err != nil {
		return nil, err
	}

	dataSourceConnection := Connection{conn}

	return &dataSourceConnection, nil
}
