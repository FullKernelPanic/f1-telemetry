package telemetry

import (
	"f1telemetry/src/telemetry/common"
	"f1telemetry/src/telemetry/f1_2021"
	"f1telemetry/src/telemetry/f1_2024"
	"f1telemetry/src/util"
)

type BaseHeader struct {
	Format uint16
}
type BasePacket struct {
	BaseHeader
}

func MapData(buf []byte) *common.MessageEnvelope {
	bp := BasePacket{}
	err := util.MapBytesToPacket(buf, &bp)

	if err != nil {
		return nil
	}

	switch bp.Format {
	case 2021:
		p := f1_2021.MapBytes(buf)
		if p != nil {
			return f1_2021.Map2021(p)
		}
	case 2024:
		p := f1_2024.MapBytes(buf)
		if p != nil {
			return f1_2024.Map2024(p)
		}
	}
	return nil
}
