package telematics

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

// argument required
const (
	ARGUMENTREQUIRED_NOTSET      byte = 0x00
	ARGUMENTREQUIRED_NOTREQUIRED byte = 0x01
)

type Response struct {
	Flags    byte
	Sequence byte
	Crc      byte
}

type ModulePropertyDisable struct {
	DisabledProperties map[byte]byte
}
