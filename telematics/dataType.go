package telematics

import (
	"encoding/binary"
	"protocol/utils"
	"time"
)

func (r *TelematicsReader) readData(t DataType) interface{} {
	switch t {
	case Bool:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case SByte:
		{
			var v int8
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Byte:
		{
			var v byte
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Short:
		{
			var v int16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case UShort:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Int24:
		{
			var v int32
			r.readInt24(&v)
			return v
		}
	case UInt24:
		{
			var v uint32
			r.readUInt24(&v)
			return v
		}
	case Int:
		{
			var v int32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case UInt:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Long:
		{
			var v int64
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case ULong:
		{
			var v uint64
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Float:
		{
			var v float32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Double:
		{
			var v float64
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case String:
		{
			return r.readString()
		}
	case Binary:
		{
			return r.readBytes()
		}
	case Identify:
		{
			return r.readBytes()
		}
	case OpenClose:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case OnOff:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case YesNo:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case IOPin:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Tamper:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Break:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Ignition:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Movement:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Alarm:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Panic:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Smoke:
		{
			var v bool
			r.readBoolean(&v)
			return v
		}
	case Frequency:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Analog:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			res := float64(v) / 1000.0
			return res
		}
	case Timestamp:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return time.Unix(int64(v), 0)
		}
	case Timespan:
		{
			var v int32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return time.Duration(v) * time.Second
		}
	case Temperature:
		{
			var v int16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) / 10.0
		}
	case Humidity:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) / 10.0
		}
	case Pressure:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) * 100.0
		}
	case Weight:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) / 1000.0
		}
	case Loudness:
		{
			var v byte
			binary.Read(r.reader, binary.LittleEndian, &v)
			return v
		}
	case Angle:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) / 100.0
		}
	case Speed:
		{
			var v uint16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float32(v) / 10.0
		}
	case Mileage:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float64(v) / 1000.0
		}
	case Rpm:
		{
			var v int16
			binary.Read(r.reader, binary.LittleEndian, &v)
			return int32(v * 10.0)
		}
	case EngineHours:
		{
			var v uint32
			r.readUInt24(&v)
			return v
		}
	case Distance:
		{
			var v uint32
			binary.Read(r.reader, binary.LittleEndian, &v)
			return float64(v) / 1000.0
		}
	case Common:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			return v
		}
	case Voltage:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 1000.0
			return v
		}
	case Battery:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 1000.0
			return v
		}
	case Power:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 1000.0
			return v
		}
	case Liquid:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 10.0
			return v
		}
	case Water:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 10.0
			return v
		}
	case Fuel:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 10.0
			return v
		}
	case Gas:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 1000.0
			v.meter = v.meter / 1000.0
			return v
		}
	case Illuminance:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 100.0
			return v
		}
	case Radiation:
		{
			v := CommonStruct{}
			r.readCommonValue(&v)
			v.value = v.value / 10.0
			v.meter = v.meter / 100.0
			return v
		}
	case IOPort:
		{
			v := IoPortStruct{}
			binary.Read(r.reader, binary.LittleEndian, &v.Flags)
			binary.Read(r.reader, binary.LittleEndian, &v.State)
			return v
		}
	case GPS:
		{
			v := GpsStruct{}
			r.readGps(&v)
			return v
		}
	case GSM:
		{
			v := GsmStruct{}
			r.readGsm(&v)
			return v
		}
	case Acceleration:
		{
			v := AccelerationStruct{}
			r.readAcceleration(&v)
			return v
		}
	case Rgb:
		{
			v := RgbStruct{}
			r.readRgb(&v)
			return v
		}
	default:
	}

	panic("not implemented dataType")
}

func (w *TelematicsWriter) writeData(v interface{}, t DataType) {
	switch t {
	case Bool:
		{
			w.writeBool(v.(bool))
		}
	case SByte:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(int8))
		}
	case Byte:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(byte))
		}
	case Short:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(int16))
		}
	case UShort:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(uint16))
		}
	case Int24:
		{
			w.writeInt24(v.(int32))
		}
	case UInt24:
		{
			w.writeUInt24(v.(uint32))
		}
	case Int:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(int32))
		}
	case UInt:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(uint32))
		}
	case Long:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(int64))
		}
	case ULong:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(uint64))
		}
	case Float:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(float32))
		}
	case Double:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(float64))
		}
	case String:
		{
			w.writeString(v.(string))
		}
	case Binary:
		{
			w.writeBytes(v.([]byte))
		}
	case Identify:
		{
			w.writeBytes(v.([]byte))
		}
	case OpenClose:
		{
			w.writeBool(v.(bool))
		}
	case OnOff:
		{
			w.writeBool(v.(bool))
		}
	case YesNo:
		{
			w.writeBool(v.(bool))
		}
	case IOPin:
		{
			w.writeBool(v.(bool))
		}
	case Tamper:
		{
			w.writeBool(v.(bool))
		}
	case Break:
		{
			w.writeBool(v.(bool))
		}
	case Ignition:
		{
			w.writeBool(v.(bool))
		}
	case Movement:
		{
			w.writeBool(v.(bool))
		}
	case Alarm:
		{
			w.writeBool(v.(bool))
		}
	case Panic:
		{
			w.writeBool(v.(bool))
		}
	case Smoke:
		{
			w.writeBool(v.(bool))
		}
	case Frequency:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(uint32))
		}
	case Analog:
		{
			analog := uint32(v.(float64) * 1000)
			binary.Write(w.writer, binary.LittleEndian, analog)
		}
	case Timestamp:
		{
			timestamp := uint32(v.(time.Time).Unix())
			binary.Write(w.writer, binary.LittleEndian, timestamp)
		}
	case Timespan:
		{
			timespan := uint32(v.(time.Duration).Seconds())
			binary.Write(w.writer, binary.LittleEndian, timespan)
		}
	case Temperature:
		{
			temperature := int16(v.(float32) * 10.0)
			binary.Write(w.writer, binary.LittleEndian, temperature)
		}
	case Humidity:
		{
			humidity := uint16(v.(float32) * 10.0)
			binary.Write(w.writer, binary.LittleEndian, humidity)
		}
	case Pressure:
		{
			pressure := uint16(v.(float32) / 100.0)
			binary.Write(w.writer, binary.LittleEndian, pressure)
		}
	case Weight:
		{
			weight := uint16(v.(float32) * 1000.0)
			binary.Write(w.writer, binary.LittleEndian, weight)
		}
	case Loudness:
		{
			binary.Write(w.writer, binary.LittleEndian, v.(byte))
		}
	case Angle:
		{
			//77.77 - bad value must be 7777 but it 7776!
			angle := uint16(utils.Round(float64(v.(float32) * 100.0)))
			binary.Write(w.writer, binary.LittleEndian, angle)
		}
	case Speed:
		{
			speed := uint16(v.(float32) * 10.0)
			binary.Write(w.writer, binary.LittleEndian, speed)
		}
	case Mileage:
		{
			mileage := uint32(v.(float64) * 1000.0)
			binary.Write(w.writer, binary.LittleEndian, mileage)
		}
	case Rpm:
		{
			rpm := v.(int32) / 10.0
			binary.Write(w.writer, binary.LittleEndian, rpm)
		}
	case EngineHours:
		{
			w.writeUInt24(v.(uint32))
		}
	case Distance:
		{
			distance := uint32(v.(float64) * 1000.0)
			binary.Write(w.writer, binary.LittleEndian, distance)
		}
	case Common:
		{
			value := v.(CommonStruct)
			w.writeCommonValue(&value)
		}
	case Voltage:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 1000.0)
			}
			w.writeCommonValue(&value)
		}
	case Battery:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 1000.0)
			}
			w.writeCommonValue(&value)
		}
	case Power:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 1000.0)
			}
			w.writeCommonValue(&value)
		}
	case Liquid:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 10.0)
			}
			w.writeCommonValue(&value)
		}
	case Water:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 10.0)
			}
			w.writeCommonValue(&value)
		}
	case Fuel:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 10.0)
			}
			w.writeCommonValue(&value)
		}
	case Gas:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 1000.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 1000.0)
			}
			w.writeCommonValue(&value)
		}
	case Illuminance:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 100.0)
			}
			w.writeCommonValue(&value)
		}
	case Radiation:
		{
			value := v.(CommonStruct)
			if value_value, ok := value.GetValue(); ok {
				value.SetValue(value_value * 10.0)
			}
			if value_meter, ok := value.GetMeter(); ok {
				value.SetMeter(value_meter * 100.0)
			}
			w.writeCommonValue(&value)
		}
	case IOPort:
		{
			value := v.(IoPortStruct)
			binary.Write(w.writer, binary.LittleEndian, value.Flags)
			binary.Write(w.writer, binary.LittleEndian, value.State)
		}
	case GPS:
		{
			gps := v.(GpsStruct)
			w.writeGps(&gps)
		}
	case GSM:
		{
			gsm := v.(GsmStruct)
			w.writeGsm(&gsm)
		}
	case Acceleration:
		{
			acc := v.(AccelerationStruct)
			w.writeAcceleration(&acc)
		}
	case Rgb:
		{
			rgb := v.(RgbStruct)
			w.writeRgb(&rgb)
		}
	default:
		panic("not implemented dataType")
	}
}
