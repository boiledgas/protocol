package value

import (
	"github.com/boiledgas/protocol/utils"
	"bytes"
	"fmt"
)

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

func (v Gps) String() string {
	var buf bytes.Buffer
	if v.Has(GPS_FLAG_LATLNG) {
		buf.WriteString(fmt.Sprintf("Lat:%v; Lng:%v; ", v.Latitude, v.Longitude))
	}
	if v.Has(GPS_FLAG_ALTITUDE) {
		buf.WriteString(fmt.Sprintf("Altitude:%v; ", v.Altitude))
	}
	if v.Has(GPS_FLAG_SPEED) {
		buf.WriteString(fmt.Sprintf("Speed:%v; ", v.Speed))
	}
	if v.Has(GPS_FLAG_COURSE) {
		buf.WriteString(fmt.Sprintf("Course:%v; ", v.Course))
	}
	if v.Has(GPS_FLAG_SATELLITES) {
		buf.WriteString(fmt.Sprintf("Sat:%v; ", v.Sat))
	}
	return buf.String()
}
