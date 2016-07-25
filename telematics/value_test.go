package telematics

import (
	"bytes"
	"testing"
	"time"
)

func Test_Bool(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Bool
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Sbyte(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := SByte
	val := int8(-7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int8) != val {
		t.Fail()
	}
}

func Test_Byte(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Byte
	val := byte(7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(byte) != val {
		t.Fail()
	}
}

func Test_Short(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Short
	val := int16(-777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int16) != val {
		t.Fail()
	}
}

func Test_UShort(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := UShort
	val := uint16(777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint16) != val {
		t.Fail()
	}
}

func Test_Int24(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Int24
	val := int32(-77777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != val {
		t.Fail()
	}
}

func Test_UInt24(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := UInt24
	val := uint32(777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Int(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Int
	val := int32(-7777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != val {
		t.Fail()
	}
}

func Test_UInt(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := UInt
	val := uint32(77777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Long(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Long
	val := int64(-777777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int64) != val {
		t.Fail()
	}
}

func Test_ULong(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := ULong
	val := uint64(7777777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint64) != val {
		t.Fail()
	}
}

func Test_Float(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Float
	val := float32(-7777777.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Double(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Double
	val := float64(7777777.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_String(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := String
	val := "777"
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(string) != val {
		t.Fail()
	}
}

func Test_Binary(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Binary
	val := []byte{0x7, 0x7, 0x7}
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if !bytes.Equal(v.([]byte), val) {
		t.Fail()
	}
}

func Test_Identify(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Identify
	val := []byte{0x7, 0x7, 0x7}
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if !bytes.Equal(v.([]byte), val) {
		t.Fail()
	}
}

func Test_OpenClose(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := OpenClose
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_OnOff(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := OnOff
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_YesNo(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := YesNo
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_IOPin(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := IOPin
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Tamper(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Tamper
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Break(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Break
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Ignition(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Ignition
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Movement(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Movement
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Alarm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Alarm
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Panic(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Panic
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Smoke(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Smoke
	val := true
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(bool) != val {
		t.Fail()
	}
}

func Test_Frequency(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Frequency
	val := uint32(777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != val {
		t.Fail()
	}
}

func Test_Analog(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Analog
	val := float64(777.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_Timestamp(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Timestamp
	val := time.Now()
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(time.Time).Unix() != val.Unix() {
		t.Fail()
	}
}

func Test_Timespan(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Timespan
	val := time.Second * 777
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(time.Duration) != val {
		t.Fail()
	}
}

func Test_Temperature(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Temperature
	val := float32(777.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Humidity(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Humidity
	val := float32(777.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Pressure(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Pressure
	val := float32(777.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != 700 {
		t.Fail()
	}
}

func Test_Weight(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Weight
	val := float32(7.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Loudness(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Loudness
	val := byte(77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(byte) != val {
		t.Fail()
	}
}

func Test_Angle(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Angle
	val := float32(77.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Speed(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Speed
	val := float32(77.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float32) != val {
		t.Fail()
	}
}

func Test_Mileage(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Mileage
	val := float64(77777.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != val {
		t.Fail()
	}
}

func Test_Rpm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Rpm
	val := int32(7777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(int32) != 7770 {
		t.Fail()
	}
}

func Test_EngineHours(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := EngineHours
	val := uint32(7777777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(uint32) != 7777777 {
		t.Fail()
	}
}

func Test_Distance(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Distance
	val := float64(7777.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(float64) != 7777.777 {
		t.Fail()
	}
}

func Test_Common(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Common
	val := CommonStruct{}
	val.SetState(true)
	val.SetValue(7)
	val.SetPercentage(77)
	val.SetMeter(777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Voltage(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Voltage
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.77)
	val.SetMeter(77.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Battery(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Battery
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.77)
	val.SetMeter(77.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Power(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Power
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.77)
	val.SetMeter(77.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Liquid(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Liquid
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.777)
	val.SetMeter(77.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Water(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Water
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.777)
	val.SetMeter(77.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Fuel(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Fuel
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.777)
	val.SetMeter(77.7)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Gas(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Gas
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.777)
	val.SetMeter(77.777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Illuminance(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Illuminance
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_Radiation(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Radiation
	val := CommonStruct{}
	val.SetState(true)
	val.SetPercentage(77)
	val.SetValue(7.7)
	val.SetMeter(77.77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(CommonStruct) != val {
		t.Fail()
	}
}

func Test_IOPort(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := IOPort
	val := IoPortStruct{Flags: 255, State: 255}
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(IoPortStruct) != val {
		t.Fail()
	}
}

func Test_Gps(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := GPS
	val := GpsStruct{}
	val.SetLatLng(77.77, 77.77)
	val.SetAltitude(77)
	val.SetSpeed(77)
	val.SetCourse(180)
	val.SetSat(77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(GpsStruct) != val {
		t.Fail()
	}
}

func Test_Gsm(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := GSM
	val := GsmStruct{}
	val.SetCID(77)
	val.SetLAC(77)
	val.SetMCC("777")
	val.SetMNC("777")
	val.SetSignal(77)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(GsmStruct) != val {
		t.Fail()
	}
}

func Test_Acceleration(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Acceleration
	val := AccelerationStruct{}
	val.SetAxisX(7.77)
	val.SetAxisY(7.77)
	val.SetAxisZ(7.77)
	val.SetDuration(777)
	w.writeData(val, dataType)
	v := r.readData(dataType)
	if v.(AccelerationStruct) != val {
		t.Fail()
	}
}

func Test_Rgb(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{reader: &buf}
	w := TelematicsWriter{writer: &buf}
	dataType := Rgb
	val := RgbStruct{r: 7, g: 77, b: 0x7}
	w.writeData(val, dataType)
	v := r.readData(dataType)
	rgb := v.(RgbStruct)
	if rgb != val {
		t.Fail()
	}
}
