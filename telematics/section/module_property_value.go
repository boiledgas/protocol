package section

import (
	"fmt"
	"bytes"
)

type ModulePropertyValue struct {
	ModuleId byte
	Values   map[byte]interface{}
}

func (s ModulePropertyValue) String() string {
	var buf bytes.Buffer
	buf.WriteString(fmt.Sprintf("{ModuleId:%v; Values:[", s.ModuleId))
	for id, value := range s.Values {
		buf.WriteString(fmt.Sprintf("{%v:%v},", id, value))
	}
	buf.WriteString("]}")
	return buf.String()
}
