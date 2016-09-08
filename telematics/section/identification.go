package section

import (
	"bytes"
	"fmt"
	"github.com/boiledgas/protocol/utils"
)

type DeviceType byte

// device type
const (
	DEVICETYPE_NOTSPECIFIED DeviceType = 0x00
	DEVICETYPE_APPLICATION  DeviceType = 0x01
	DEVICETYPE_PERSONAL     DeviceType = 0x02
	DEVICETYPE_STATIONARY   DeviceType = 0x03
	DEVICETYPE_CAR          DeviceType = 0x04
	DEVICETYPE_CAROBD       DeviceType = 0x05
	DEVICETYPE_CARSOCKET    DeviceType = 0x06
	DEVICETYPE_CARBEACON    DeviceType = 0x07
)

// identification flags
const (
	IDENTIFICATION_FLAGS_CODE       byte = 0x01
	IDENTIFICATION_FLAGS_CODETEXT   byte = 0x02
	IDENTIFICATION_FLAGS_DEVICETYPE byte = 0x04
	IDENTIFICATION_FLAGS_FIRMWARE   byte = 0x08
	IDENTIFICATION_FLAGS_HARDWARE   byte = 0x10
	IDENTIFICATION_FLAGS_DEVICEHASH byte = 0x20
)

type Identification struct {
	utils.Flags8
	Code     uint32
	CodeText string
	Type     DeviceType
	Firmware int16
	Hardware int16
	Hash     byte
}

func (s Identification) String() string {
	var buf bytes.Buffer
	buf.WriteString("Identification {")
	var flags [8]byte
	s.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case IDENTIFICATION_FLAGS_CODE:
			buf.WriteString(fmt.Sprintf("Code: %v; ", s.Code))
		case IDENTIFICATION_FLAGS_CODETEXT:
			buf.WriteString(fmt.Sprintf("CodeText: %v; ", s.CodeText))
		case IDENTIFICATION_FLAGS_DEVICETYPE:
			buf.WriteString(fmt.Sprintf("Type: %v; ", s.Type))
		case IDENTIFICATION_FLAGS_FIRMWARE:
			buf.WriteString(fmt.Sprintf("Firmware: %v; ", s.Firmware))
		case IDENTIFICATION_FLAGS_HARDWARE:
			buf.WriteString(fmt.Sprintf("Hardware: %v; ", s.Hardware))
		case IDENTIFICATION_FLAGS_DEVICEHASH:
			buf.WriteString(fmt.Sprintf("Hash: %v; ", s.Hash))
		}
	}
	buf.WriteString("}")
	return buf.String()
}
