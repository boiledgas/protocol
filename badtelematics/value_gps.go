package telematics

// gpsData flags
const (
	GPSSTRUCT_FLAGS_LATLNG     byte = 0x01
	GPSSTRUCT_FLAGS_ALTITUDE        = 0x02
	GPSSTRUCT_FLAGS_SPEED           = 0x04
	GPSSTRUCT_FLAGS_COURSE          = 0x08
	GPSSTRUCT_FLAGS_SATELLITES      = 0x10
)

type GpsStruct struct {
	flags byte

	latitude  float64
	longitude float64
	altitude  int16
	speed     byte
	course    byte
	sat       byte
}

type GpsValue interface {
	GetFlag() byte

	GetLatLng() (float64, float64, bool)
	SetLatLng(float64, float64)

	GetAltitude() (int16, bool)
	SetAltitude(int16)

	GetSpeed() (byte, bool)
	SetSpeed(byte)

	GetCourse() (byte, bool)
	SetCourse(byte)

	GetSat() (byte, bool)
	SetSat(byte)
}

func (c *GpsStruct) GetFlag() byte {
	return c.flags
}

func (c *GpsStruct) GetLatLng() (lat float64, lng float64, ok bool) {
	lat = c.latitude
	lng = c.longitude
	ok = c.flags&GPSSTRUCT_FLAGS_LATLNG > 0
	return
}

func (c *GpsStruct) SetLatLng(lat float64, lng float64) {
	c.latitude = lat
	c.longitude = lng
	c.flags = c.flags | GPSSTRUCT_FLAGS_LATLNG
}

func (c *GpsStruct) GetAltitude() (val int16, ok bool) {
	val = c.altitude
	ok = c.flags&GPSSTRUCT_FLAGS_ALTITUDE > 0
	return
}

func (c *GpsStruct) SetAltitude(val int16) {
	c.altitude = val
	c.flags = c.flags | GPSSTRUCT_FLAGS_ALTITUDE
}

func (c *GpsStruct) GetSpeed() (val byte, ok bool) {
	val = c.speed
	ok = c.flags&GPSSTRUCT_FLAGS_SPEED > 0
	return
}

func (c *GpsStruct) SetSpeed(val byte) {
	c.speed = val
	c.flags = c.flags | GPSSTRUCT_FLAGS_SPEED
}

func (c *GpsStruct) GetCourse() (val byte, ok bool) {
	val = c.course
	ok = c.flags&GPSSTRUCT_FLAGS_COURSE > 0
	return
}

func (c *GpsStruct) SetCourse(val byte) {
	c.course = val
	c.flags = c.flags | GPSSTRUCT_FLAGS_COURSE
}

func (c *GpsStruct) GetSat() (val byte, ok bool) {
	val = c.sat
	ok = c.flags&GPSSTRUCT_FLAGS_SATELLITES > 0
	return
}

func (c *GpsStruct) SetSat(val byte) {
	c.sat = val
	c.flags = c.flags | GPSSTRUCT_FLAGS_SATELLITES
}
