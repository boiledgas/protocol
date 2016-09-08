package telematics

import (
	"encoding/binary"
	"fmt"
	"io"
	"github.com/boiledgas/protocol/telematics/value"
	"github.com/boiledgas/protocol/utils"
)

type TelematicsReader struct {
	Configuration *Configuration
	checksum      utils.Checksum
	reader        io.Reader
	buffer        [255]byte // skip buffer
}

func NewReader(r io.Reader) *TelematicsReader {
	checksum := utils.Checksum{Table: utils.CRC8[:]}
	reader := TelematicsReader{checksum: checksum, reader: io.TeeReader(r, &checksum)}
	return &reader
}

func (r *TelematicsReader) skip(c byte) {
	b := r.buffer[0:c]
	binary.Read(r.reader, binary.LittleEndian, b)
}

// общие методы чтения
func (r *TelematicsReader) ReadBoolean(v *bool) {
	var flag byte
	binary.Read(r.reader, binary.LittleEndian, &flag)
	*v = flag == 1
}

func (r *TelematicsReader) ReadInt24(v *int32) {
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

func (r *TelematicsReader) ReadUInt24(v *uint32) {
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

func (r *TelematicsReader) ReadByte(v *byte) (err error) {
	err = binary.Read(r.reader, binary.BigEndian, v)
	return
}

//bad method
func (r *TelematicsReader) ReadBytes() []byte {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	buf := make([]byte, int(c))
	binary.Read(r.reader, binary.LittleEndian, buf)
	return buf
}

//bad method
func (r *TelematicsReader) ReadString() string {
	return string(r.ReadBytes())
}

// специализированные методы чтения
func (r *TelematicsReader) ReadCommon(v *value.Common) {
	binary.Read(r.reader, binary.LittleEndian, &v.Flags8)
	var flags [8]byte
	v.Load(&flags)
	for _, flag := range flags {
		switch flag {
		case value.COMMON_FLAG_STATE:
			r.ReadBoolean(&v.State)
			v.Set(value.COMMON_FLAG_STATE, true)
		case value.COMMON_FLAG_PERCENTAGE:
			binary.Read(r.reader, binary.LittleEndian, &v.Percentage)
			v.Set(value.COMMON_FLAG_PERCENTAGE, true)
		case value.COMMON_FLAG_VALUE:
			var val int32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.Value = float64(val)
			v.Set(value.COMMON_FLAG_VALUE, true)
		case value.COMMON_FLAG_METER:
			var val uint32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.Meter = float64(val)
			v.Set(value.COMMON_FLAG_METER, true)
		default:
			panic("not implemented")
		}
	}
}

func (r *TelematicsReader) ReadGps(v *value.Gps) {
	binary.Read(r.reader, binary.LittleEndian, &v.Flags8)
	var flags [8]byte
	v.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case value.GPS_FLAG_LATLNG:
			var val int32
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.Latitude = float64(val) / 10000000.0
			binary.Read(r.reader, binary.LittleEndian, &val)
			v.Longitude = float64(val) / 10000000.0
			v.Set(value.GPS_FLAG_LATLNG, true)
		case value.GPS_FLAG_ALTITUDE:
			v.Set(value.GPS_FLAG_ALTITUDE, true)
			binary.Read(r.reader, binary.LittleEndian, &v.Altitude)
		case value.GPS_FLAG_SPEED:
			v.Set(value.GPS_FLAG_SPEED, true)
			binary.Read(r.reader, binary.LittleEndian, &v.Speed)
		case value.GPS_FLAG_COURSE:
			v.Set(value.GPS_FLAG_COURSE, true)
			binary.Read(r.reader, binary.LittleEndian, &v.Course)
			v.Course = v.Course * 2
		case value.GPS_FLAG_SATELLITES:
			v.Set(value.GPS_FLAG_SATELLITES, true)
			binary.Read(r.reader, binary.LittleEndian, &v.Sat)
		default:
			panic("gpsData flag not supported")
		}
	}
}

func (r *TelematicsReader) ReadGsm(v *value.Gsm) {
	var mcc_mnc int32
	r.ReadInt24(&mcc_mnc)
	mcc_mnc_str := fmt.Sprintf("%d", mcc_mnc)
	if len(mcc_mnc_str) < 3 {
		v.MCC = mcc_mnc_str
		v.MNC = "0"
		if len(v.MCC) == 0 {
			v.MCC = "0"
		}
	} else {
		v.MCC = mcc_mnc_str[:3]
		v.MNC = mcc_mnc_str[3:]
	}

	binary.Read(r.reader, binary.LittleEndian, &v.LAC)
	binary.Read(r.reader, binary.LittleEndian, &v.CID)
	binary.Read(r.reader, binary.LittleEndian, &v.Signal)
}

func (r *TelematicsReader) ReadAcceleration(v *value.Acceleration) {
	binary.Read(r.reader, binary.LittleEndian, &v.Flags8)
	var flags [8]byte
	v.Load(&flags)
	var axis int16
	var mult float32 = 1000.0
	for _, flag := range flags {
		switch flag {
		case value.ACCELERATION_FLAG_X:
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.AxisX = float32(axis) / mult
		case value.ACCELERATION_FLAG_Y:
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.AxisY = float32(axis) / mult
		case value.ACCELERATION_FLAG_Z:
			binary.Read(r.reader, binary.LittleEndian, &axis)
			v.AxisZ = float32(axis) / mult
		case value.ACCELERATION_FLAG_DURATION:
			binary.Read(r.reader, binary.LittleEndian, &v.Duration)
		default:
			panic(fmt.Sprintf("acceleration flag %X not supported", flag))
		}
	}
}

func (r *TelematicsReader) ReadRgb(v *value.Rgb) {
	binary.Read(r.reader, binary.LittleEndian, &v.R)
	binary.Read(r.reader, binary.LittleEndian, &v.G)
	binary.Read(r.reader, binary.LittleEndian, &v.B)
}

func (r *TelematicsReader) ReadNameValue(dataType value.DataType, v *value.NameValue) {
	if dataType == value.NotSet {
		panic("type for read must be set")
	}

	v.Name = r.ReadString()
	v.Value = r.readData(dataType)
}

func (r *TelematicsReader) ReadNameValues(dataType value.DataType) []value.NameValue {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	result := make([]value.NameValue, int(c))
	for i := byte(0); i < c; i++ {
		var v value.NameValue
		r.ReadNameValue(dataType, &v)
		result = append(result, v)
	}

	return result
}
