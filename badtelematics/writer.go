package telematics

import (
	"encoding/binary"
	"io"
	"strconv"
)

type TelematicsWriter struct {
	Writer   io.Writer
	conf     *Configuration
	checksum *Checksum
}

func NewWriter(w io.Writer, c *Configuration) *TelematicsWriter {
	if c == nil {
		c = NewConfiguration()
	}

	checksum := NewChecksum()
	writer := io.MultiWriter(w, checksum)
	return &TelematicsWriter{Writer: writer, conf: c, checksum: checksum}
}

func (w *TelematicsWriter) writeBool(v bool) {
	val := byte(0)
	if v {
		val = 1
	}
	binary.Write(w.Writer, binary.LittleEndian, val)
}

func (w *TelematicsWriter) writeInt24(v int32) {
	buf := [3]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
	binary.Write(w.Writer, binary.LittleEndian, buf)
}

func (w *TelematicsWriter) writeUInt24(v uint32) {
	buf := [3]byte{
		byte(v),
		byte(v >> 8),
		byte(v >> 16),
	}
	binary.Write(w.Writer, binary.LittleEndian, buf)
}

func (w *TelematicsWriter) writeBytes(buf []byte) {
	if err := binary.Write(w.Writer, binary.LittleEndian, byte(len(buf))); err != nil {
		panic(err)
	}
	if err := binary.Write(w.Writer, binary.LittleEndian, buf); err != nil {
		panic(err)
	}
}

func (w *TelematicsWriter) writeString(s string) {
	c := byte(len(s))
	binary.Write(w.Writer, binary.LittleEndian, c)
	binary.Write(w.Writer, binary.LittleEndian, []byte(s))
}

func (w *TelematicsWriter) writeCommonValue(v *CommonStruct) {
	flag := v.GetFlag()
	binary.Write(w.Writer, binary.LittleEndian, flag)
	if state, ok := v.GetState(); ok {
		w.writeBool(state)
	}
	if perc, ok := v.GetPercentage(); ok {
		binary.Write(w.Writer, binary.LittleEndian, perc)
	}
	if val, ok := v.GetValue(); ok {
		binary.Write(w.Writer, binary.LittleEndian, int32(val))
	}
	if meter, ok := v.GetMeter(); ok {
		binary.Write(w.Writer, binary.LittleEndian, uint32(meter))
	}
}

func (w *TelematicsWriter) writeNameValue(v *NameValue, dataType DataType) {
	if dataType == NotSet {
		panic("type not set")
	}

	w.writeString(v.Name)
	w.writeData(v.Value, dataType)
}

func (w *TelematicsWriter) writeNameList(list []NameValue, dataType DataType) {
	binary.Write(w.Writer, binary.LittleEndian, byte(len(list)))
	for _, v := range list {
		w.writeNameValue(&v, dataType)
	}
}

func (w *TelematicsWriter) writeGps(v *GpsStruct) {
	binary.Write(w.Writer, binary.LittleEndian, v.flags)
	if lat, lng, ok := v.GetLatLng(); ok {
		latVal := int32(lat * 10000000)
		lngVal := int32(lng * 10000000)
		binary.Write(w.Writer, binary.LittleEndian, latVal)
		binary.Write(w.Writer, binary.LittleEndian, lngVal)
	}
	if alt, ok := v.GetAltitude(); ok {
		binary.Write(w.Writer, binary.LittleEndian, alt)
	}
	if speed, ok := v.GetSpeed(); ok {
		binary.Write(w.Writer, binary.LittleEndian, speed)
	}
	if course, ok := v.GetCourse(); ok {
		val := course / 2
		binary.Write(w.Writer, binary.LittleEndian, val)
	}
	if sat, ok := v.GetSat(); ok {
		binary.Write(w.Writer, binary.LittleEndian, sat)
	}
}

func (w *TelematicsWriter) writeGsm(v *GsmStruct) {
	mm := v.mcc + v.mnc
	mcc_mnc, err := strconv.ParseUint(mm, 10, 32)
	if err != nil {
		mcc_mnc = 0
	}
	w.writeUInt24(uint32(mcc_mnc))
	binary.Write(w.Writer, binary.LittleEndian, v.lac)
	binary.Write(w.Writer, binary.LittleEndian, v.cid)
	binary.Write(w.Writer, binary.LittleEndian, v.signal)
}

func (w *TelematicsWriter) writeAcceleration(v *AccelerationStruct) {
	binary.Write(w.Writer, binary.LittleEndian, v.flags)
	if x, ok := v.GetAxisX(); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(x*1000.0))
	}
	if y, ok := v.GetAxisY(); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(y*1000.0))
	}
	if z, ok := v.GetAxisZ(); ok {
		binary.Write(w.Writer, binary.LittleEndian, int16(z*1000.0))
	}
	binary.Write(w.Writer, binary.LittleEndian, v.duration)
}

func (w *TelematicsWriter) writeRgb(v *RgbStruct) {
	binary.Write(w.Writer, binary.LittleEndian, v.r)
	binary.Write(w.Writer, binary.LittleEndian, v.g)
	binary.Write(w.Writer, binary.LittleEndian, v.b)
}
