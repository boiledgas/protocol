package telematics

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

func (s *Module) GetId() byte {
	return s.Id
}

func (s *Module) SetId(id byte) {
	s.Id = id
}

func (s *Module) GetName() (name string, ok bool) {
	name = s.Name
	ok = s.Has(MODULE_FLAGS_NAME)
	return
}

func (s *Module) SetName(name string) {
	s.Name = name
	s.Set(MODULE_FLAGS_NAME, true)
}

func (s *Module) GetDesc() (desc string, ok bool) {
	desc = s.Description
	ok = s.Has(MODULE_FLAGS_DESCRIPTION)
	return
}

func (s *Module) SetDesc(desc string) {
	s.Description = desc
	s.Set(MODULE_FLAGS_DESCRIPTION, true)
}
