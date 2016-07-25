package telematics

// module property flags
const (
	MODULE_PROPERTY_FLAGS_MIN         byte = 0x01
	MODULE_PROPERTY_FLAGS_MAX              = 0x02
	MODULE_PROPERTY_FLAGS_LIST             = 0x04
	MODULE_PROPERTY_FLAGS_ACCESS           = 0x08
	MODULE_PROPERTY_FLAGS_NAME             = 0x10
	MODULE_PROPERTY_FLAGS_DESCRIPTION      = 0x20
)

type PropertyAccess byte

// property access
const (
	PROPERTYACCESS_NOTSET   PropertyAccess = 0x00
	PROPERTYACCESS_READ     PropertyAccess = 0x01
	PROPERTYACCESS_WRITE    PropertyAccess = 0x02
	PROPERTYACCESS_CONFIG   PropertyAccess = 0x04
	PROPERTYACCESS_DISABLED PropertyAccess = 0x08
)

type modulePropertySection struct {
	baseSection

	ModuleId byte
	Id       byte
	Type     DataType
	Min      interface{}
	Max      interface{}
	List     []NameValue
	Access   PropertyAccess
	Name     string
	Desc     string
}

type ModuleProperty interface {
	Section

	GetId() byte
	SetId(id byte)

	GetModuleId() byte

	GetType() DataType
	SetType(dataType DataType)

	GetMin() (interface{}, bool)
	SetMin(min interface{})

	GetMax() (interface{}, bool)
	SetMax(min interface{})

	GetList() ([]NameValue, bool)
	SetList(list []NameValue)

	GetAccess() (PropertyAccess, bool)
	SetAccess(access PropertyAccess)

	GetName() (string, bool)
	SetName(min string)

	GetDesc() (string, bool)
	SetDesc(desc string)
}

func NewModuleProperty(id byte, m Module) ModuleProperty {
	v := modulePropertySection{ModuleId: m.GetId(), Id: id}
	return &v
}

func (s *modulePropertySection) GetId() byte {
	return s.Id
}

func (s *modulePropertySection) SetId(id byte) {
	s.Id = id
}

func (s *modulePropertySection) GetModuleId() byte {
	return s.ModuleId
}

func (s *modulePropertySection) SetModuleId(id byte) {
	s.ModuleId = id
}

func (s *modulePropertySection) GetType() DataType {
	return s.Type
}

func (s *modulePropertySection) SetType(dataType DataType) {
	s.Type = dataType
}

func (s *modulePropertySection) GetMin() (min interface{}, ok bool) {
	min = s.Min
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_MIN)
	return
}

func (s *modulePropertySection) SetMin(min interface{}) {
	s.Min = min
	s.setFlag(MODULE_PROPERTY_FLAGS_MIN, true)
}

func (s *modulePropertySection) GetMax() (max interface{}, ok bool) {
	max = s.Max
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_MAX)
	return
}

func (s *modulePropertySection) SetMax(max interface{}) {
	s.Max = max
	s.setFlag(MODULE_PROPERTY_FLAGS_MAX, true)
}

func (s *modulePropertySection) GetList() (list []NameValue, ok bool) {
	list = s.List
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_LIST)
	return
}

func (s *modulePropertySection) SetList(list []NameValue) {
	s.List = list
	s.setFlag(MODULE_PROPERTY_FLAGS_LIST, true)
}

func (s *modulePropertySection) GetAccess() (access PropertyAccess, ok bool) {
	access = s.Access
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_ACCESS)
	return
}

func (s *modulePropertySection) SetAccess(access PropertyAccess) {
	s.Access = access
	s.setFlag(MODULE_PROPERTY_FLAGS_ACCESS, access != PROPERTYACCESS_NOTSET)
}

func (s *modulePropertySection) GetName() (name string, ok bool) {
	name = s.Name
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_NAME)
	return
}

func (s *modulePropertySection) SetName(name string) {
	s.Name = name
	s.setFlag(MODULE_PROPERTY_FLAGS_NAME, true)
}

func (s *modulePropertySection) GetDesc() (desc string, ok bool) {
	desc = s.Desc
	ok = s.hasFlag(MODULE_PROPERTY_FLAGS_DESCRIPTION)
	return
}

func (s *modulePropertySection) SetDesc(desc string) {
	s.Desc = desc
	s.setFlag(MODULE_PROPERTY_FLAGS_DESCRIPTION, true)
}
