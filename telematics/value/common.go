package value

import "protocol/utils"

const (
	COMMON_FLAG_STATE      byte = 0x01
	COMMON_FLAG_PERCENTAGE byte = 0x02
	COMMON_FLAG_VALUE      byte = 0x04
	COMMON_FLAG_METER      byte = 0x08
)

type Common struct {
	utils.Flags8
	State      bool
	Percentage byte
	Value      float64
	Meter      float64
}
