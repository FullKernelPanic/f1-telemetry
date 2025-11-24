package telemetry

import (
	"f1telemetry/src/telemetry/mapper"
	"f1telemetry/src/telemetry/versions"
	"f1telemetry/src/telemetry/versions/f1_2021"
	"f1telemetry/src/util"
)

type BaseHeader struct {
	Format uint16
}
type BasePacket struct {
	BaseHeader
}

func MapData(buf []byte) *versions.MessageEnvelope {
	bp := BasePacket{}
	err := util.MapBytesToPacket(buf, &bp)

	if err != nil {
		return nil
	}

	switch bp.Format {
	case 2021:
		p := f1_2021.MapBytes(buf)
		if p != nil {
			return mapper.Map2021(p)
		}
	}
	return nil
}
