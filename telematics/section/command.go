package section

import (
	"protocol/utils"
	"fmt"
)

// command flags
const (
	COMMAND_FLAGS_NAME        byte = 0x01
	COMMAND_FLAGS_DESCRIPTION byte = 0x02
)

type Command struct {
	utils.Flags8
	ModuleId    byte
	Id          byte
	Name        string
	Description string
}

func (c Command) String() string {
	return fmt.Sprintf("{ModuleId:%v; Id:%v; Name:%v}", c.ModuleId, c.Id, c.Name)
}
