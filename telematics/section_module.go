package telematics

// module flags
const (
	MODULE_FLAGS_NAME        byte = 0x01
	MODULE_FLAGS_DESCRIPTION byte = 0x02
)

type moduleSection struct {
	baseSection

	Id          byte
	Name        string
	Description string
}

type Module interface {
	GetId() byte
	GetName() (string, bool)
	SetName(name string)
	GetDesc() (string, bool)
	SetDesc(desc string)
}

func NewModule(id byte) Module {
	s := moduleSection{Id: id}
	return &s
}

func (s *moduleSection) GetId() byte {
	return s.Id
}

func (s *moduleSection) SetId(id byte) {
	s.Id = id
}

func (s *moduleSection) GetName() (name string, ok bool) {
	name = s.Name
	ok = s.hasFlag(MODULE_FLAGS_NAME)
	return
}

func (s *moduleSection) SetName(name string) {
	s.Name = name
	s.setFlag(MODULE_FLAGS_NAME, true)
}

func (s *moduleSection) GetDesc() (desc string, ok bool) {
	desc = s.Description
	ok = s.hasFlag(MODULE_FLAGS_DESCRIPTION)
	return
}

func (s *moduleSection) SetDesc(desc string) {
	s.Description = desc
	s.setFlag(MODULE_FLAGS_DESCRIPTION, true)
}
