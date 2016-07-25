package telematics

const (
	COMMON_VALUE_FLAGS_STATE      byte = 0x01
	COMMON_VALUE_FLAGS_PERCENTAGE byte = 0x02
	COMMON_VALUE_FLAGS_VALUE      byte = 0x04
	COMMON_VALUE_FLAGS_METER      byte = 0x08
)

type CommonStruct struct {
	state_set      bool
	percentage_set bool
	value_set      bool
	meter_set      bool

	state      bool
	percentage byte
	value      float64
	meter      float64
}

type CommonValue interface {
	GetFlag() byte

	GetState() (bool, bool)
	SetState(bool)

	GetPercentage() (byte, bool)
	SetPercentage(byte)

	GetValue() (float64, bool)
	SetValue(float64)

	GetMeter() (float64, bool)
	SetMeter(float64)
}

func (c *CommonStruct) GetFlag() byte {
	flag := byte(0)
	if c.state_set {
		flag = flag | COMMON_VALUE_FLAGS_STATE
	}
	if c.percentage_set {
		flag = flag | COMMON_VALUE_FLAGS_PERCENTAGE
	}
	if c.value_set {
		flag = flag | COMMON_VALUE_FLAGS_VALUE
	}
	if c.meter_set {
		flag = flag | COMMON_VALUE_FLAGS_METER
	}

	return flag
}

func (c *CommonStruct) GetState() (state bool, ok bool) {
	state = c.state
	ok = c.state_set
	return
}

func (c *CommonStruct) SetState(state bool) {
	c.state = state
	c.state_set = true
}

func (c *CommonStruct) GetPercentage() (perc byte, ok bool) {
	perc = c.percentage
	ok = c.percentage_set
	return
}

func (c *CommonStruct) SetPercentage(perc byte) {
	c.percentage = perc
	c.percentage_set = true
}

func (c *CommonStruct) GetValue() (val float64, ok bool) {
	val = c.value
	ok = c.value_set
	return
}

func (c *CommonStruct) SetValue(val float64) {
	c.value = val
	c.value_set = true
}

func (c *CommonStruct) GetMeter() (meter float64, ok bool) {
	meter = c.meter
	ok = c.meter_set
	return
}

func (c *CommonStruct) SetMeter(meter float64) {
	c.meter = meter
	c.meter_set = true
}
