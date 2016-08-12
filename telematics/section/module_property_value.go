package section

import "fmt"

type modulePropertyValueSection struct {
	baseSection

	moduleId byte
	values   map[byte]interface{}
}

type ModulePropertyValue interface {
	Section

	GetModuleId() byte

	GetValue(id byte) (interface{}, bool)
	SetValue(id byte, v interface{})
}

func NewModulePropertyValue(m Module) ModulePropertyValue {
	v := modulePropertyValueSection{moduleId: m.GetId(), values: make(map[byte]interface{})}
	return &v
}

func (s *modulePropertyValueSection) GetModuleId() byte {
	return s.moduleId
}

func (s *modulePropertyValueSection) SetModuleId(id byte) {
	s.moduleId = id
}

func (s *modulePropertyValueSection) GetValue(id byte) (v interface{}, ok bool) {
	v, ok = s.values[id]
	return
}

func (s *modulePropertyValueSection) SetValue(id byte, v interface{}) {
	if _, ok := s.values[id]; ok {
		panic(fmt.Sprintf("id exists %d", id))
	}

	s.values[id] = v
}
