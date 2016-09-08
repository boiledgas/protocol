package telematics

import (
	"encoding/binary"
	"io"
	"github.com/boiledgas/protocol/telematics/value"
	"github.com/boiledgas/protocol/utils"
	"strconv"
)

type TelematicsWriter struct {
	Writer        io.Writer
	Checksum      utils.Checksum
	Configuration *Configuration
}

func NewWriter(w io.Writer) *TelematicsWriter {
	checksum := utils.Checksum{Table: utils.CRC8[:]}
	writer := io.MultiWriter(w, &checksum)
	return &TelematicsWriter{Writer: writer, Checksum: checksum}
}

func (w *TelematicsWriter) WriteBool(v bool) {
	val := byte(0)
	if v {
		val = 1
	}
	binary.Write(w.Writer, binary.LittleEndian, val)
}

func (w *TelematicsWriter) WriteInt24(v int32) {
	buf := [3]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
	binary.Write(w.Writer, binary.LittleEndian, buf)
}

func (w *TelematicsWriter) WriteUInt24(v uint32) {
	buf := [3]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
	binary.Write(w.Writer, binary.LittleEndian, buf)
}

func (w *TelematicsWriter) WriteBytes(buf []byte) {
	if err := binary.Write(w.Writer, binary.LittleEndian, byte(len(buf))); err != nil {
		panic(err)
	}
	if err := binary.Write(w.Writer, binary.LittleEndian, buf); err != nil {
		panic(err)
	}
}

func (w *TelematicsWriter) WriteString(s string) {
	c := byte(len(s))
	binary.Write(w.Writer, binary.LittleEndian, c)
	binary.Write(w.Writer, binary.LittleEndian, []byte(s))
}

func (w *TelematicsWriter) WriteCommon(v *value.Common) {
	binary.Write(w.Writer, binary.LittleEndian, v.Flags8)
	if v.Has(value.COMMON_FLAG_STATE) {
		w.WriteBool(v.State)
	}
	if v.Has(value.COMMON_FLAG_PERCENTAGE) {
		binary.Write(w.Writer, binary.LittleEndian, v.Percentage)
	}
	if v.Has(value.COMMON_FLAG_VALUE) {
		binary.Write(w.Writer, binary.LittleEndian, int32(v.Value))
	}
	if v.Has(value.COMMON_FLAG_METER) {
		binary.Write(w.Writer, binary.LittleEndian, uint32(v.Meter))
	}
}

func (w *TelematicsWriter) WriteNameValue(v *value.NameValue, dataType value.DataType) {
	if dataType == value.NotSet {
		panic("type not set")
	}

	w.WriteString(v.Name)
	w.WriteData(v.Value, dataType)
}

func (w *TelematicsWriter) WriteNameList(list []value.NameValue, dataType value.DataType) {
	binary.Write(w.Writer, binary.LittleEndian, byte(len(list)))
	for _, v := range list {
		w.WriteNameValue(&v, dataType)
	}
}

func (w *TelematicsWriter) WriteGps(v *value.Gps) {
	binary.Write(w.Writer, binary.LittleEndian, v.Flags8)
	if lat, lng, ok := v.Latitude, v.Longitude, v.Has(value.GPS_FLAG_LATLNG); ok {
		latVal := int32(lat * 10000000)
		lngVal := int32(lng * 10000000)
		binary.Write(w.Writer, binary.LittleEndian, latVal)
		binary.Write(w.Writer, binary.LittleEndian, lngVal)
	}
	if alt, ok := v.Altitude, v.Has(value.GPS_FLAG_ALTITUDE); ok {
		binary.Write(w.Writer, binary.LittleEndian, alt)
	}
	if speed, ok := v.Speed, v.Has(value.GPS_FLAG_SPEED); ok {
		binary.Write(w.Writer, binary.LittleEndian, speed)
	}
	if course, ok := v.Course, v.Has(value.GPS_FLAG_COURSE); ok {
		val := course / 2
		binary.Write(w.Writer, binary.LittleEndian, val)
	}
	if sat, ok := v.Sat, v.Has(value.GPS_FLAG_SATELLITES); ok {
		binary.Write(w.Writer, binary.LittleEndian, sat)
	}
}

func (w *TelematicsWriter) WriteGsm(v *value.Gsm) {
	mm := v.MCC + v.MNC
	mcc_mnc, err := strconv.ParseUint(mm, 10, 32)
	if err != nil {
		mcc_mnc = 0
	}
	w.WriteUInt24(uint32(mcc_mnc))
	binary.Write(w.Writer, binary.LittleEndian, v.LAC)
	binary.Write(w.Writer, binary.LittleEndian, v.CID)
	binary.Write(w.Writer, binary.LittleEndian, v.Signal)
}

func (w *TelematicsWriter) WriteAcceleration(v *value.Acceleration) {
	binary.Write(w.Writer, binary.LittleEndian, v.Flags8)
	if x, ok := v.AxisX, v.Has(value.ACCELERATION_FLAG_X); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(x*1000.0))
	}
	if y, ok := v.AxisY, v.Has(value.ACCELERATION_FLAG_Y); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(y*1000.0))
	}
	if z, ok := v.AxisZ, v.Has(value.ACCELERATION_FLAG_Z); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(z*1000.0))
	}
	binary.Write(w.Writer, binary.LittleEndian, v.Duration)
}

func (w *TelematicsWriter) WriteRgb(v *value.Rgb) {
	binary.Write(w.Writer, binary.LittleEndian, v.R)
	binary.Write(w.Writer, binary.LittleEndian, v.G)
	binary.Write(w.Writer, binary.LittleEndian, v.B)
}
