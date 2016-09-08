package telematics

import (
	"bytes"
	"testing"
	"time"
	"github.com/boiledgas/protocol/telematics/value"
)

func Test_Bool(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Bool
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Sbyte(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.SByte
	val := int8(-7)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int8) != val {
		t.Fail()
	}
}

func Test_Byte(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Byte
	val := byte(7)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(byte) != val {
		t.Fail()
	}
}

func Test_Short(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Short
	val := int16(-777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int16) != val {
		t.Fail()
	}
}

func Test_UShort(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.UShort
	val := uint16(777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint16) != val {
		t.Fail()
	}
}

func Test_Int24(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Int24
	val := int32(-77777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != val {
		t.Fail()
	}
}

func Test_UInt24(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.UInt24
	val := uint32(777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Int(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Int
	val := int32(-7777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != val {
		t.Fail()
	}
}

func Test_UInt(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.UInt
	val := uint32(77777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Long(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Long
	val := int64(-777777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int64) != val {
		t.Fail()
	}
}

func Test_ULong(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.ULong
	val := uint64(7777777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint64) != val {
		t.Fail()
	}
}

func Test_Float(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Float
	val := float32(-7777777.77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Double(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Double
	val := float64(7777777.777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_String(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.String
	val := "777"
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(string) != val {
		t.Fail()
	}
}

func Test_Binary(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Binary
	val := []byte{0x7, 0x7, 0x7}
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if !bytes.Equal(v.([]byte), val) {
		t.Fail()
	}
}

func Test_Identify(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Identify
	val := []byte{0x7, 0x7, 0x7}
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if !bytes.Equal(v.([]byte), val) {
		t.Fail()
	}
}

func Test_OpenClose(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.OpenClose
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_OnOff(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.OnOff
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_YesNo(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.YesNo
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_IOPin(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.IOPin
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Tamper(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Tamper
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Break(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Break
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Ignition(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Ignition
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Movement(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Movement
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Alarm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Alarm
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Panic(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Panic
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Smoke(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Smoke
	val := true
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Frequency(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Frequency
	val := uint32(777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Analog(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Analog
	val := float64(777.777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_Timestamp(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Timestamp
	val := time.Now()
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(time.Time).Unix() != val.Unix() {
		t.Fail()
	}
}

func Test_Timespan(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Timespan
	val := time.Second * 777
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(time.Duration) != val {
		t.Fail()
	}
}

func Test_Temperature(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Temperature
	val := float32(777.7)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Humidity(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Humidity
	val := float32(777.7)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Pressure(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Pressure
	val := float32(777.77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != 700 {
		t.Fail()
	}
}

func Test_Weight(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Weight
	val := float32(7.777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Loudness(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Loudness
	val := byte(77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(byte) != val {
		t.Fail()
	}
}

func Test_Angle(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Angle
	val := float32(77.77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Speed(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Speed
	val := float32(77.7)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Mileage(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Mileage
	val := float64(77777.777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_Rpm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Rpm
	val := int32(7777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != 7770 {
		t.Fail()
	}
}

func Test_EngineHours(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.EngineHours
	val := uint32(7777777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != 7777777 {
		t.Fail()
	}
}

func Test_Distance(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Distance
	val := float64(7777.777)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != 7777.777 {
		t.Fail()
	}
}

func Test_Common(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.COMMON
	val := value.Common{}
	val.State = true
	val.Value = 7
	val.Percentage = 77
	val.Meter = 777
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Voltage(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Voltage
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.77
	val.Meter = 77.77
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Battery(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Battery
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.77
	val.Meter = 77.77
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Power(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Power
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.77
	val.Meter = 77.77
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Liquid(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Liquid
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.777
	val.Meter = 77.7
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Water(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Water
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.777
	val.Meter = 77.7
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Fuel(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Fuel
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.777
	val.Meter = 77.7
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Gas(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Gas
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.777
	val.Meter = 77.777
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Illuminance(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Illuminance
	val := value.Common{}
	val.State = true
	val.Percentage = 77
	val.Value = 7.77
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_Radiation(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.Radiation
	val := value.Common{}
	val.State = true
	val.Set(value.COMMON_FLAG_STATE, true)
	val.Percentage = 77
	val.Set(value.COMMON_FLAG_PERCENTAGE, true)
	val.Value = 7.7
	val.Set(value.COMMON_FLAG_VALUE, true)
	val.Meter = 77.77
	val.Set(value.COMMON_FLAG_METER, true)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Common) != val {
		t.Fail()
	}
}

func Test_IOPort(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.IOPort
	val := value.IoPort{Flags: 255, State: 255}
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.IoPort) != val {
		t.Fail()
	}
}

func Test_Gps(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.GPS
	val := value.Gps{}
	val.Latitude, val.Longitude = 77.77, 77.77
	val.Altitude = (77)
	val.Speed = (77)
	val.Course = (180)
	val.Sat = (77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Gps) != val {
		t.Fail()
	}
}

func Test_Gsm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.GSM
	val := value.Gsm{}
	val.CID = (77)
	val.LAC = (77)
	val.MCC = ("777")
	val.MNC = ("777")
	val.Signal = (77)
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Gsm) != val {
		t.Fail()
	}
}

func Test_Acceleration(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.ACCELERATION
	val := value.Acceleration{}
	val.AxisX = 7.77
	val.AxisY = 7.77
	val.AxisZ = 7.77
	val.Duration = 777
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	if v.(value.Acceleration) != val {
		t.Fail()
	}
}

func Test_Rgb(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{Writer: &buf}
	dataType := value.RGB
	val := value.Rgb{R: 7, G: 77, B: 0x7}
	w.WriteData(val, dataType)
	v := r.readData(dataType)
	rgb := v.(value.Rgb)
	if rgb != val {
		t.Fail()
	}
}
