package util

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"log"
	"net"
	"os/exec"
	"runtime"
)

func IpAddress() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	return conn.LocalAddr().(*net.UDPAddr).IP
}

func Openbrowser(url string) {
	var err error

	switch runtime.GOOS {
	case "linux":
		err = exec.Command("xdg-open", url).Start()
	case "windows":
		err = exec.Command("rundll32", "url.dll,FileProtocolHandler", url).Start()
	case "darwin":
		err = exec.Command("open", url).Start()
	default:
		err = fmt.Errorf("unsupported platform")
	}
	if err != nil {
		log.Fatal(err)
	}

}

func MapBytesToPacket(data []byte, packet interface{}) error {
	reader := bytes.NewReader(data)
	if err := binary.Read(reader, binary.LittleEndian, packet); err != nil {
		return err
	}
	return nil
}
