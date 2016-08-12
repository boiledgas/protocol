package telematics

const (
	ACCELERATION_FLAGS_X        byte = 0x01
	ACCELERATION_FLAGS_Y             = 0x02
	ACCELERATION_FLAGS_Z             = 0x04
	ACCELERATION_FLAGS_DURATION      = 0x08
)

type AccelerationStruct struct {
	flags byte

	axisX    float32
	axisY    float32
	axisZ    float32
	duration uint16
}

type AccelerationValue interface {
	GetFlags() byte

	GetAxisX() (float32, bool)
	SetAxisX(float32)

	GetAxisY() (float32, bool)
	SetAxisY(float32)

	GetAxisZ() (float32, bool)
	SetAxisZ(float32)

	GetDuration() (uint16, bool)
	SetDuration(uint16)
}

func (v *AccelerationStruct) GetFlags() byte {
	return v.flags
}

func (c *AccelerationStruct) GetAxisX() (val float32, ok bool) {
	val = c.axisX
	ok = c.flags&ACCELERATION_FLAGS_X > 0
	return
}

func (c *AccelerationStruct) SetAxisX(val float32) {
	c.axisX = val
	c.flags = c.flags | ACCELERATION_FLAGS_X
}

func (c *AccelerationStruct) GetAxisY() (val float32, ok bool) {
	val = c.axisY
	ok = c.flags&ACCELERATION_FLAGS_Y > 0
	return
}

func (c *AccelerationStruct) SetAxisY(val float32) {
	c.axisY = val
	c.flags = c.flags | ACCELERATION_FLAGS_Y
}

func (c *AccelerationStruct) GetAxisZ() (val float32, ok bool) {
	val = c.axisZ
	ok = c.flags&ACCELERATION_FLAGS_Z > 0
	return
}

func (c *AccelerationStruct) SetAxisZ(val float32) {
	c.axisZ = val
	c.flags = c.flags | ACCELERATION_FLAGS_Z
}

func (c *AccelerationStruct) GetDuration() (val uint16, ok bool) {
	val = c.duration
	ok = c.flags&ACCELERATION_FLAGS_DURATION > 0
	return
}

func (c *AccelerationStruct) SetDuration(val uint16) {
	c.duration = val
	c.flags = c.flags | ACCELERATION_FLAGS_DURATION
}
