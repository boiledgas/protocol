package section

import "fmt"

type CommandExecute struct {
	ModuleId  byte
	CommandId byte
	Arguments map[byte]interface{}
}

func (s CommandExecute) String() string {
	return fmt.Sprintf("%v", s)
}
