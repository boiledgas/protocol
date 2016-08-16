package section

import "fmt"

type ModulePropertyDisable struct {
	DisabledProperties map[byte]byte
}

func (s ModulePropertyDisable) String() string {
	return fmt.Sprintf("%v", s)
}
