package telematics

import (
	"encoding/binary"
	"protocol/telematics/value"
	"protocol/utils"
	"time"
)

func (r *TelematicsReader) readData(t value.DataType) interface{} {
	switch t {
	case value.Bool:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.SByte:
		var v int8
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Byte:
		var v byte
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Short:
		var v int16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.UShort:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Int24:
		var v int32
		r.ReadInt24(&v)
		return v
	case value.UInt24:
		var v uint32
		r.ReadUInt24(&v)
		return v
	case value.Int:
		var v int32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.UInt:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Long:
		var v int64
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.ULong:
		var v uint64
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Float:
		var v float32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Double:
		var v float64
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.String:
		return r.ReadString()
	case value.Binary:
		return r.ReadBytes()
	case value.Identify:
		return r.ReadBytes()
	case value.OpenClose:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.OnOff:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.YesNo:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.IOPin:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Tamper:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Break:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Ignition:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Movement:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Alarm:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Panic:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Smoke:
		var v bool
		r.ReadBoolean(&v)
		return v
	case value.Frequency:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Analog:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		res := float64(v) / 1000.0
		return res
	case value.Timestamp:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return time.Unix(int64(v), 0)
	case value.Timespan:
		var v int32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return time.Duration(v) * time.Second
	case value.Temperature:
		var v int16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) / 10.0
	case value.Humidity:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) / 10.0
	case value.Pressure:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) * 100.0
	case value.Weight:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) / 1000.0
	case value.Loudness:
		var v byte
		binary.Read(r.reader, binary.LittleEndian, &v)
		return v
	case value.Angle:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) / 100.0
	case value.Speed:
		var v uint16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float32(v) / 10.0
	case value.Mileage:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float64(v) / 1000.0
	case value.Rpm:
		var v int16
		binary.Read(r.reader, binary.LittleEndian, &v)
		return int32(v * 10.0)
	case value.EngineHours:
		var v uint32
		r.ReadUInt24(&v)
		return v
	case value.Distance:
		var v uint32
		binary.Read(r.reader, binary.LittleEndian, &v)
		return float64(v) / 1000.0
	case value.COMMON:
		v := value.Common{}
		r.ReadCommon(&v)
		return v
	case value.Voltage:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 1000.0
		return v
	case value.Battery:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 1000.0
		return v
	case value.Power:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 1000.0
		return v
	case value.Liquid:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 10.0
		return v
	case value.Water:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 10.0
		return v
	case value.Fuel:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 10.0
		return v
	case value.Gas:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 1000.0
		v.Meter = v.Meter / 1000.0
		return v
	case value.Illuminance:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 100.0
		return v
	case value.Radiation:
		v := value.Common{}
		r.ReadCommon(&v)
		v.Value = v.Value / 10.0
		v.Meter = v.Meter / 100.0
		return v
	case value.IOPort:
		v := value.IoPort{}
		binary.Read(r.reader, binary.LittleEndian, &v.Flags)
		binary.Read(r.reader, binary.LittleEndian, &v.State)
		return v
	case value.GPS:
		v := value.Gps{}
		r.ReadGps(&v)
		return v
	case value.GSM:
		v := value.Gsm{}
		r.ReadGsm(&v)
		return v
	case value.ACCELERATION:
		v := value.Acceleration{}
		r.ReadAcceleration(&v)
		return v
	case value.RGB:
		v := value.Rgb{}
		r.ReadRgb(&v)
		return v
	default:
	}
	panic("not implemented dataType")
}

func (w *TelematicsWriter) WriteData(v interface{}, t value.DataType) {
	switch t {
	case value.Bool:
		w.WriteBool(v.(bool))
	case value.SByte:
		binary.Write(w.Writer, binary.LittleEndian, v.(int8))
	case value.Byte:
		binary.Write(w.Writer, binary.LittleEndian, v.(byte))
	case value.Short:
		binary.Write(w.Writer, binary.LittleEndian, v.(int16))
	case value.UShort:
		binary.Write(w.Writer, binary.LittleEndian, v.(uint16))
	case value.Int24:
		w.WriteInt24(v.(int32))
	case value.UInt24:
		w.WriteUInt24(v.(uint32))
	case value.Int:
		binary.Write(w.Writer, binary.LittleEndian, v.(int32))
	case value.UInt:
		binary.Write(w.Writer, binary.LittleEndian, v.(uint32))
	case value.Long:
		binary.Write(w.Writer, binary.LittleEndian, v.(int64))
	case value.ULong:
		binary.Write(w.Writer, binary.LittleEndian, v.(uint64))
	case value.Float:
		binary.Write(w.Writer, binary.LittleEndian, v.(float32))
	case value.Double:
		binary.Write(w.Writer, binary.LittleEndian, v.(float64))
	case value.String:
		w.WriteString(v.(string))
	case value.Binary:
		w.WriteBytes(v.([]byte))
	case value.Identify:
		w.WriteBytes(v.([]byte))
	case value.OpenClose:
		w.WriteBool(v.(bool))
	case value.OnOff:
		w.WriteBool(v.(bool))
	case value.YesNo:
		w.WriteBool(v.(bool))
	case value.IOPin:
		w.WriteBool(v.(bool))
	case value.Tamper:
		w.WriteBool(v.(bool))
	case value.Break:
		w.WriteBool(v.(bool))
	case value.Ignition:
		w.WriteBool(v.(bool))
	case value.Movement:
		w.WriteBool(v.(bool))
	case value.Alarm:
		w.WriteBool(v.(bool))
	case value.Panic:
		w.WriteBool(v.(bool))
	case value.Smoke:
		w.WriteBool(v.(bool))
	case value.Frequency:
		binary.Write(w.Writer, binary.LittleEndian, v.(uint32))
	case value.Analog:
		analog := uint32(v.(float64) * 1000)
		binary.Write(w.Writer, binary.LittleEndian, analog)
	case value.Timestamp:
		timestamp := uint32(v.(time.Time).Unix())
		binary.Write(w.Writer, binary.LittleEndian, timestamp)
	case value.Timespan:
		timespan := uint32(v.(time.Duration).Seconds())
		binary.Write(w.Writer, binary.LittleEndian, timespan)
	case value.Temperature:
		temperature := int16(v.(float32) * 10.0)
		binary.Write(w.Writer, binary.LittleEndian, temperature)
	case value.Humidity:
		humidity := uint16(v.(float32) * 10.0)
		binary.Write(w.Writer, binary.LittleEndian, humidity)
	case value.Pressure:
		pressure := uint16(v.(float32) / 100.0)
		binary.Write(w.Writer, binary.LittleEndian, pressure)
	case value.Weight:
		weight := uint16(v.(float32) * 1000.0)
		binary.Write(w.Writer, binary.LittleEndian, weight)
	case value.Loudness:
		binary.Write(w.Writer, binary.LittleEndian, v.(byte))
	case value.Angle:
		//77.77 - bad value must be 7777 but it 7776!
		angle := uint16(utils.Round(float64(v.(float32) * 100.0)))
		binary.Write(w.Writer, binary.LittleEndian, angle)
	case value.Speed:
		speed := uint16(v.(float32) * 10.0)
		binary.Write(w.Writer, binary.LittleEndian, speed)
	case value.Mileage:
		mileage := uint32(v.(float64) * 1000.0)
		binary.Write(w.Writer, binary.LittleEndian, mileage)
	case value.Rpm:
		rpm := v.(int32) / 10.0
		binary.Write(w.Writer, binary.LittleEndian, rpm)
	case value.EngineHours:
		w.WriteUInt24(v.(uint32))
	case value.Distance:
		distance := uint32(v.(float64) * 1000.0)
		binary.Write(w.Writer, binary.LittleEndian, distance)
	case value.COMMON:
		value := v.(value.Common)
		w.WriteCommon(&value)
	case value.Voltage:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 1000.0
		}
		w.WriteCommon(&val)
	case value.Battery:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 1000.0
		}
		w.WriteCommon(&val)
	case value.Power:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 1000.0
		}
		w.WriteCommon(&val)
	case value.Liquid:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 10.0
		}
		w.WriteCommon(&val)
	case value.Water:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 10.0
		}
		w.WriteCommon(&val)
	case value.Fuel:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 10.0
		}
		w.WriteCommon(&val)
	case value.Gas:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 1000.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 1000.0
		}
		w.WriteCommon(&val)
	case value.Illuminance:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 100.0
		}
		w.WriteCommon(&val)
	case value.Radiation:
		val := v.(value.Common)
		if voltage_value, ok := val.Value, val.Has(value.COMMON_FLAG_VALUE); ok {
			val.Value = voltage_value * 10.0
		}
		if value_meter, ok := val.Meter, val.Has(value.COMMON_FLAG_METER); ok {
			val.Value = value_meter * 100.0
		}
		w.WriteCommon(&val)
	case value.IOPort:
		value := v.(value.IoPort)
		binary.Write(w.Writer, binary.LittleEndian, value.Flags)
		binary.Write(w.Writer, binary.LittleEndian, value.State)
	case value.GPS:
		gps := v.(value.Gps)
		w.WriteGps(&gps)
	case value.GSM:
		gsm := v.(value.Gsm)
		w.WriteGsm(&gsm)
	case value.ACCELERATION:
		acc := v.(value.Acceleration)
		w.WriteAcceleration(&acc)
	case value.RGB:
		rgb := v.(value.Rgb)
		w.WriteRgb(&rgb)
	default:
		panic("not implemented dataType")
	}
}
