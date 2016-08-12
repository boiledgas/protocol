package value

import "protocol/utils"

const (
	ACCELERATION_FLAG_X        byte = 0x01
	ACCELERATION_FLAG_Y        byte = 0x02
	ACCELERATION_FLAG_Z        byte = 0x04
	ACCELERATION_FLAG_DURATION byte = 0x08
)

type Acceleration struct {
	utils.Flags8
	AxisX    float32
	AxisY    float32
	AxisZ    float32
	Duration uint16
}
