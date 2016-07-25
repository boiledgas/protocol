package telematics

import "bytes"

// identification flags
const (
	IDENTIFICATION_FLAGS_CODE       byte = 0x01
	IDENTIFICATION_FLAGS_CODETEXT   byte = 0x02
	IDENTIFICATION_FLAGS_DEVICETYPE byte = 0x04
	IDENTIFICATION_FLAGS_FIRMWARE   byte = 0x08
	IDENTIFICATION_FLAGS_HARDWARE   byte = 0x10
	IDENTIFICATION_FLAGS_DEVICEHASH byte = 0x20
)

type Identification interface {
	Section

	GetCode() (code uint32, ok bool)
	SetCode(code uint32)

	GetCodeText() (code string, ok bool)
	SetCodeText(code string)

	GetDeviceType() (deviceType DeviceType, ok bool)
	SetDeviceType(deviceType DeviceType)

	GetFirmware() (firmware int16, ok bool)
	SetFirmware(firmware int16)

	GetHardware() (hardware int16, ok bool)
	SetHardware(hardware int16)

	GetHash() (hash byte, ok bool)
	SetHash(hash byte)
}

type identificationSection struct {
	baseSection

	code       uint32
	codeText   string
	deviceType DeviceType
	firmware   int16
	hardware   int16
	hash       byte
}

func NewIdentification() Identification {
	s := identificationSection{}
	return &s
}

func (s *identificationSection) GetCode() (code uint32, ok bool) {
	code = s.code
	ok = s.hasFlag(IDENTIFICATION_FLAGS_CODE)
	return
}

func (s *identificationSection) SetCode(code uint32) {
	s.setFlag(IDENTIFICATION_FLAGS_CODE, code > 0)
	s.code = code
}

func (s *identificationSection) GetCodeText() (code string, ok bool) {
	code = s.codeText
	ok = s.hasFlag(IDENTIFICATION_FLAGS_CODETEXT)
	return
}

func (s *identificationSection) SetCodeText(code string) {
	s.setFlag(IDENTIFICATION_FLAGS_CODETEXT, len(code) > 0)
	s.codeText = code
}

func (s *identificationSection) GetDeviceType() (deviceType DeviceType, ok bool) {
	deviceType = s.deviceType
	ok = s.hasFlag(IDENTIFICATION_FLAGS_DEVICETYPE)
	return
}

func (s *identificationSection) SetDeviceType(deviceType DeviceType) {
	i := bytes.IndexByte(deviceTypes, byte(deviceType))
	s.setFlag(IDENTIFICATION_FLAGS_DEVICETYPE, i > -1)
	s.deviceType = deviceType
}

func (s *identificationSection) GetFirmware() (firmware int16, ok bool) {
	firmware = s.firmware
	ok = s.hasFlag(IDENTIFICATION_FLAGS_FIRMWARE)
	return
}

func (s *identificationSection) SetFirmware(firmware int16) {
	s.setFlag(IDENTIFICATION_FLAGS_FIRMWARE, firmware > -1)
	s.firmware = firmware
}

func (s *identificationSection) GetHardware() (hardware int16, ok bool) {
	hardware = s.hardware
	ok = s.hasFlag(IDENTIFICATION_FLAGS_HARDWARE)
	return
}

func (s *identificationSection) SetHardware(hardware int16) {
	s.setFlag(IDENTIFICATION_FLAGS_HARDWARE, hardware > -1)
	s.hardware = hardware
}

func (s *identificationSection) GetHash() (hash byte, ok bool) {
	hash = s.hash
	ok = s.hasFlag(IDENTIFICATION_FLAGS_DEVICEHASH)
	return
}

func (s *identificationSection) SetHash(hash byte) {
	s.setFlag(IDENTIFICATION_FLAGS_DEVICEHASH, true)
	s.hash = hash
}
