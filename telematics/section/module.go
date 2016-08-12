package section

import "protocol/utils"

// module flags
const (
	MODULE_FLAGS_NAME        byte = 0x01
	MODULE_FLAGS_DESCRIPTION byte = 0x02
)

type Module struct {
	utils.Flags8
	Id          byte
	Name        string
	Description string
}
