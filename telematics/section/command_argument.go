package section

import (
	"fmt"
	"protocol/telematics/value"
	"protocol/utils"
)

// command argument flags
const (
	COMMAND_ARGUMENT_FLAGS_MIN         byte = 0x01
	COMMAND_ARGUMENT_FLAGS_MAX         byte = 0x02
	COMMAND_ARGUMENT_FLAGS_LIST        byte = 0x04
	COMMAND_ARGUMENT_FLAGS_REQUIRED    byte = 0x08
	COMMAND_ARGUMENT_FLAGS_NAME        byte = 0x10
	COMMAND_ARGUMENT_FLAGS_DESCRIPTION byte = 0x20
)

type CommandArgument struct {
	utils.Flags8
	ModuleId  byte
	CommandId byte
	Id        byte
	Type      value.DataType
	Min       interface{}
	Max       interface{}
	List      []value.NameValue
	Required  byte
	Name      string
	Desc      string
}

func (ca CommandArgument) String() string {
	return fmt.Sprintf("{Id:%v; ModuleId:%v; CommandId:%v; Type:%v}", ca.Id, ca.ModuleId, ca.CommandId, ca.Type)
}
