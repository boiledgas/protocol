package value

type DataType byte

const (
	NotSet       DataType = 0x00
	Bool         DataType = 0x01
	SByte        DataType = 0x02
	Byte         DataType = 0x03
	Short        DataType = 0x04
	UShort       DataType = 0x05
	Int24        DataType = 0x06
	UInt24       DataType = 0x07
	Int          DataType = 0x08
	UInt         DataType = 0x09
	Long         DataType = 0x0A
	ULong        DataType = 0x0B
	Float        DataType = 0x0C
	Double       DataType = 0x0D
	Array        DataType = 0x0E
	String       DataType = 0x0F
	Binary       DataType = 0x10
	Id           DataType = 0x11
	Name         DataType = 0x12
	COMMON       DataType = 0x13
	OpenClose    DataType = 0x14
	OnOff        DataType = 0x15
	YesNo        DataType = 0x16
	IOPin        DataType = 0x17
	Tamper       DataType = 0x18
	Break        DataType = 0x19
	Ignition     DataType = 0x1A
	Movement     DataType = 0x1B
	Alarm        DataType = 0x1C
	Panic        DataType = 0x1D
	Smoke        DataType = 0x1E
	Frequency    DataType = 0x1F
	Analog       DataType = 0x20
	Timestamp    DataType = 0x21
	Timespan     DataType = 0x22
	Temperature  DataType = 0x23
	Humidity     DataType = 0x24
	Pressure     DataType = 0x25
	Weight       DataType = 0x26
	Loudness     DataType = 0x27
	Angle        DataType = 0x28
	Speed        DataType = 0x29
	Mileage      DataType = 0x2A
	Rpm          DataType = 0x2B
	EngineHours  DataType = 0x2C
	Distance     DataType = 0x2D
	Identify     DataType = 0x2E
	Voltage      DataType = 0x2F
	Battery      DataType = 0x30
	Power        DataType = 0x31
	Liquid       DataType = 0x32
	Water        DataType = 0x33
	Fuel         DataType = 0x34
	Gas          DataType = 0x35
	IOPort       DataType = 0x36
	GPS          DataType = 0x37
	GSM          DataType = 0x38
	ACCELERATION DataType = 0x39
	DataSampling DataType = 0x3A
	Sound        DataType = 0x3B
	Accident     DataType = 0x3C
	TextMessage  DataType = 0x3D
	Illuminance  DataType = 0x3E
	Radiation    DataType = 0x3F
	RGB          DataType = 0x41
)

func (t DataType) String() string {
	switch t {
	case NotSet:
		return "NotSet"
	case Bool:
		return "Bool"
	case SByte:
		return "SByte"
	case Byte:
		return "Byte"
	case Short:
		return "Short"
	case UShort:
		return "UShort"
	case Int24:
		return "Int24"
	case UInt24:
		return "UInt24"
	case Int:
		return "Int"
	case UInt:
		return "UInt"
	case Long:
		return "Long"
	case ULong:
		return "ULong"
	case Float:
		return "Float"
	case Double:
		return "Double"
	case Array:
		return "Array"
	case String:
		return "String"
	case Binary:
		return "Binary"
	case Id:
		return "Id"
	case Name:
		return "Name"
	case COMMON:
		return "COMMON"
	case OpenClose:
		return "OpenClose"
	case OnOff:
		return "OnOff"
	case YesNo:
		return "YesNo"
	case IOPin:
		return "IOPin"
	case Tamper:
		return "Tamper"
	case Break:
		return "Break"
	case Ignition:
		return "Ignition"
	case Movement:
		return "Movement"
	case Alarm:
		return "Alarm"
	case Panic:
		return "Panic"
	case Smoke:
		return "Smoke"
	case Frequency:
		return "Frequency"
	case Analog:
		return "Analog"
	case Timestamp:
		return "Timestamp"
	case Timespan:
		return "Timespan"
	case Temperature:
		return "Temperature"
	case Humidity:
		return "Humidity"
	case Pressure:
		return "Pressure"
	case Weight:
		return "Weight"
	case Loudness:
		return "Loudness"
	case Angle:
		return "Angle"
	case Speed:
		return "Speed"
	case Mileage:
		return "Mileage"
	case Rpm:
		return "Rpm"
	case EngineHours:
		return "EngineHours"
	case Distance:
		return "Distance"
	case Identify:
		return "Identify"
	case Voltage:
		return "Voltage"
	case Battery:
		return "Battery"
	case Power:
		return "Power"
	case Liquid:
		return "Liquid"
	case Water:
		return "Water"
	case Fuel:
		return "Fuel"
	case Gas:
		return "Gas"
	case IOPort:
		return "IOPort"
	case GPS:
		return "GPS"
	case GSM:
		return "GSM"
	case ACCELERATION:
		return "ACCELERATION"
	case DataSampling:
		return "DataSampling"
	case Sound:
		return "Sound"
	case Accident:
		return "Accident"
	case TextMessage:
		return "TextMessage"
	case Illuminance:
		return "Illuminance"
	case Radiation:
		return "Radiation"
	case RGB:
		return "RGB"
	}
	return ""
}
