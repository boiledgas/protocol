package telematics

import (
	"bytes"
	"protocol/utils"
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

	code       uint32
	codeText   string
	deviceType DeviceType
	firmware   int16
	hardware   int16
	hash       byte
}

func (s *Identification) GetCode() (code uint32, ok bool) {
	code = s.code
	ok = s.Has(IDENTIFICATION_FLAGS_CODE)
	return
}

func (s *Identification) SetCode(code uint32) {
	s.Set(IDENTIFICATION_FLAGS_CODE, code > 0)
	s.code = code
}

func (s *Identification) GetCodeText() (code string, ok bool) {
	code = s.codeText
	ok = s.Has(IDENTIFICATION_FLAGS_CODETEXT)
	return
}

func (s *Identification) SetCodeText(code string) {
	s.Set(IDENTIFICATION_FLAGS_CODETEXT, len(code) > 0)
	s.codeText = code
}

func (s *Identification) GetDeviceType() (deviceType DeviceType, ok bool) {
	deviceType = s.deviceType
	ok = s.Has(IDENTIFICATION_FLAGS_DEVICETYPE)
	return
}

func (s *Identification) SetDeviceType(deviceType DeviceType) {
	i := bytes.IndexByte(deviceTypes, byte(deviceType))
	s.Set(IDENTIFICATION_FLAGS_DEVICETYPE, i > -1)
	s.deviceType = deviceType
}

func (s *Identification) GetFirmware() (firmware int16, ok bool) {
	firmware = s.firmware
	ok = s.Has(IDENTIFICATION_FLAGS_FIRMWARE)
	return
}

func (s *Identification) SetFirmware(firmware int16) {
	s.Set(IDENTIFICATION_FLAGS_FIRMWARE, firmware > -1)
	s.firmware = firmware
}

func (s *Identification) GetHardware() (hardware int16, ok bool) {
	hardware = s.hardware
	ok = s.Has(IDENTIFICATION_FLAGS_HARDWARE)
	return
}

func (s *Identification) SetHardware(hardware int16) {
	s.Set(IDENTIFICATION_FLAGS_HARDWARE, hardware > -1)
	s.hardware = hardware
}

func (s *Identification) GetHash() (hash byte, ok bool) {
	hash = s.hash
	ok = s.Has(IDENTIFICATION_FLAGS_DEVICEHASH)
	return
}

func (s *Identification) SetHash(hash byte) {
	s.Set(IDENTIFICATION_FLAGS_DEVICEHASH, true)
	s.hash = hash
}
