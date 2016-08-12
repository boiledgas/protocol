package section

import "protocol/telematics/value"

// command argument flags
const (
	COMMAND_ARGUMENT_FLAGS_MIN         byte = 0x01
	COMMAND_ARGUMENT_FLAGS_MAX              = 0x02
	COMMAND_ARGUMENT_FLAGS_LIST             = 0x04
	COMMAND_ARGUMENT_FLAGS_REQUIRED         = 0x08
	COMMAND_ARGUMENT_FLAGS_NAME             = 0x10
	COMMAND_ARGUMENT_FLAGS_DESCRIPTION      = 0x20
)

type CommandArgument struct {
	moduleId  byte
	commandId byte
	id        byte
	dataType  value.DataType
	min       interface{}
	max       interface{}
	list      []value.NameValue
	required  byte
	name      string
	desc      string
}
