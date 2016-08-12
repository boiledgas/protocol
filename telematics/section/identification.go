package section

import (
	"protocol/utils"
)

type DeviceType byte

var deviceTypes []byte = []byte{
	byte(DEVICETYPE_NOTSPECIFIED),
	byte(DEVICETYPE_APPLICATION),
	byte(DEVICETYPE_PERSONAL),
	byte(DEVICETYPE_STATIONARY),
	byte(DEVICETYPE_CAR),
	byte(DEVICETYPE_CAROBD),
	byte(DEVICETYPE_CARSOCKET),
	byte(DEVICETYPE_CARBEACON),
}

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
	Flag       utils.Flags8
	Code       uint32
	CodeText   string
	DeviceType DeviceType
	Firmware   int16
	Hardware   int16
	Hash       byte
}
