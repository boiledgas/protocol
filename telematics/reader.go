package telematics

import (
	"encoding/binary"
	"fmt"
	"io"
	"protocol/utils"
)

type TelematicsReader struct {
	reader   io.Reader
	conf     *conf
	checksum *Checksum
}

func NewReader(r io.Reader, c *conf) *TelematicsReader {
	if c == nil {
		c = NewConfiguration()
	}

	checksum := NewChecksum()
	reader := io.TeeReader(r, checksum)
	return &TelematicsReader{reader: reader, conf: c, checksum: checksum}
}

func (r *TelematicsReader) skip(c byte) {
	b := make([]byte, c)
	binary.Read(r.reader, binary.LittleEndian, b)
}

// общие методы чтения
func (r *TelematicsReader) readBoolean(v *bool) {
	var flag byte
	binary.Read(r.reader, binary.LittleEndian, &flag)
	*v = flag == 1
}

func (r *TelematicsReader) readInt24(v *int32) {
	var buf [3]byte
	binary.Read(r.reader, binary.LittleEndian, &buf)
	x := int(buf[0]) | (int(buf[1]) << 8) | (int(buf[2]) << 16)
	if x&0x800000 > 0 {
		x |= 0xff000000
	} else {
		x &= 0xffffff
	}

	*v = int32(x)
}

func (r *TelematicsReader) readUInt24(v *uint32) {
	var buf [3]byte
	binary.Read(r.reader, binary.LittleEndian, &buf)
	x := uint(buf[0]) | (uint(buf[1]) << 8) | (uint(buf[2]) << 16)
	if x&0x800000 > 0 {
		x |= 0xff000000
	} else {
		x &= 0xffffff
	}

	*v = uint32(x)
}

func (r *TelematicsReader) readBytes() []byte {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	buf := make([]byte, c)
	binary.Read(r.reader, binary.LittleEndian, buf)
	return buf
}

func (r *TelematicsReader) readString() string {
	return string(r.readBytes())
}

// специализированные методы чтения
func (r *TelematicsReader) readCommonValue(v *CommonStruct) {
	var f byte
	binary.Read(r.reader, binary.LittleEndian, &f)
	flags := utils.GetFlags8(f)
	for _, flag := range flags {
		switch flag {
		case COMMON_VALUE_FLAGS_STATE:
			v.state_set = true
			r.readBoolean(&v.state)
		case COMMON_VALUE_FLAGS_PERCENTAGE:
			v.percentage_set = true
			binary.Read(r.reader, binary.LittleEndian, &v.percentage)
		case COMMON_VALUE_FLAGS_VALUE:
			v.value_set = true
			var val int32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.value = float64(val)
		case COMMON_VALUE_FLAGS_METER:
			v.meter_set = true
			var val uint32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.meter = float64(val)
		default:
			panic("not implemented")
		}
	}
}

func (r *TelematicsReader) readGps(v *GpsStruct) {
	var f byte
	binary.Read(r.reader, binary.LittleEndian, &f)
	flags := utils.GetFlags8(f)
	for _, flag := range flags {
		switch flag {
		case GPSSTRUCT_FLAGS_LATLNG:
			v.flags = v.flags | GPSSTRUCT_FLAGS_LATLNG
			var val int32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.latitude = float64(val) / 10000000.0
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.longitude = float64(val) / 10000000.0
		case GPSSTRUCT_FLAGS_ALTITUDE:
			v.flags = v.flags | GPSSTRUCT_FLAGS_ALTITUDE
			binary.Read(r.reader, binary.LittleEndian, &v.altitude)
		case GPSSTRUCT_FLAGS_SPEED:
			v.flags = v.flags | GPSSTRUCT_FLAGS_SPEED
			binary.Read(r.reader, binary.LittleEndian, &v.speed)
		case GPSSTRUCT_FLAGS_COURSE:
			v.flags = v.flags | GPSSTRUCT_FLAGS_COURSE
			binary.Read(r.reader, binary.LittleEndian, &v.course)
			v.course = v.course * 2
		case GPSSTRUCT_FLAGS_SATELLITES:
			v.flags = v.flags | GPSSTRUCT_FLAGS_SATELLITES
			binary.Read(r.reader, binary.LittleEndian, &v.sat)
		default:
			panic("gpsData flag not supported")
		}
	}
}

func (r *TelematicsReader) readGsm(v *GsmStruct) {
	var mcc_mnc int32
	r.readInt24(&mcc_mnc)
	mcc_mnc_str := fmt.Sprintf("%d", mcc_mnc)
	if len(mcc_mnc_str) < 3 {
		v.mcc = mcc_mnc_str
		v.mnc = "0"
		if len(v.mcc) == 0 {
			v.mcc = "0"
		}
	} else {
		v.mcc = mcc_mnc_str[:3]
		v.mnc = mcc_mnc_str[3:]
	}

	binary.Read(r.reader, binary.LittleEndian, &v.lac)
	binary.Read(r.reader, binary.LittleEndian, &v.cid)
	binary.Read(r.reader, binary.LittleEndian, &v.signal)
}

func (r *TelematicsReader) readAcceleration(v *AccelerationStruct) {
	var f byte
	binary.Read(r.reader, binary.LittleEndian, &f)
	flags := utils.GetFlags8(f)
	var axis int16
	var mult float32 = 1000.0
	for _, flag := range flags {
		switch flag {
		case ACCELERATION_FLAGS_X:
			v.flags = v.flags | ACCELERATION_FLAGS_X
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.axisX = float32(axis) / mult
		case ACCELERATION_FLAGS_Y:
			v.flags = v.flags | ACCELERATION_FLAGS_Y
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.axisY = float32(axis) / mult
		case ACCELERATION_FLAGS_Z:
			v.flags = v.flags | ACCELERATION_FLAGS_Z
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.axisZ = float32(axis) / mult
		case ACCELERATION_FLAGS_DURATION:
			v.flags = v.flags | ACCELERATION_FLAGS_DURATION
			binary.Read(r.reader, binary.LittleEndian, &v.duration)
		default:
			panic(fmt.Sprintf("acceleration flag %X not supported", flag))
		}
	}
}

func (r *TelematicsReader) readRgb(v *RgbStruct) {
	binary.Read(r.reader, binary.LittleEndian, &v.r)
	binary.Read(r.reader, binary.LittleEndian, &v.g)
	binary.Read(r.reader, binary.LittleEndian, &v.b)
}

func (r *TelematicsReader) readNameValue(dataType DataType) NameValue {
	if dataType == NotSet {
		panic("type for read must be set")
	}

	v := NameValue{}
	v.Name = r.readString()
	v.Value = r.readData(dataType)
	return v
}

func (r *TelematicsReader) readNameValues(dataType DataType) []NameValue {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	result := make([]NameValue, c)
	for i := byte(0); i < c; i++ {
		v := r.readNameValue(dataType)
		result = append(result, v)
	}

	return result
}
