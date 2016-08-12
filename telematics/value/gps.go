package value

import "protocol/utils"

// gpsData flags
const (
	GPS_FLAG_LATLNG     byte = 0x01
	GPS_FLAG_ALTITUDE   byte = 0x02
	GPS_FLAG_SPEED      byte = 0x04
	GPS_FLAG_COURSE     byte = 0x08
	GPS_FLAG_SATELLITES byte = 0x10
)

type Gps struct {
	utils.Flags8
	Latitude  float64
	Longitude float64
	Altitude  int16
	Speed     byte
	Course    byte
	Sat       byte
}
