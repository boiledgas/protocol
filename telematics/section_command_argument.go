package telematics

// command argument flags
const (
	COMMAND_ARGUMENT_FLAGS_MIN         byte = 0x01
	COMMAND_ARGUMENT_FLAGS_MAX              = 0x02
	COMMAND_ARGUMENT_FLAGS_LIST             = 0x04
	COMMAND_ARGUMENT_FLAGS_REQUIRED         = 0x08
	COMMAND_ARGUMENT_FLAGS_NAME             = 0x10
	COMMAND_ARGUMENT_FLAGS_DESCRIPTION      = 0x20
)

type commandArgumentStruct struct {
	baseSection

	moduleId  byte
	commandId byte
	id        byte
	dataType  DataType
	min       interface{}
	max       interface{}
	list      []NameValue
	required  byte
	name      string
	desc      string
}

type CommandArgument interface {
	Section

	GetModuleId() byte
	GetCommandId() byte

	GetId() byte
	SetId(byte)

	GetType() DataType
	SetType(DataType)

	GetMin() (interface{}, bool)
	SetMin(min interface{})

	GetMax() (interface{}, bool)
	SetMax(min interface{})

	GetList() ([]NameValue, bool)
	SetList(list []NameValue)

	GetRequired() (byte, bool)
	SetRequired(byte)

	GetName() (string, bool)
	SetName(min string)

	GetDesc() (string, bool)
	SetDesc(desc string)
}

func NewCommandArgument(id byte, c Command) CommandArgument {
	ca := commandArgumentStruct{commandId: c.GetId(), moduleId: c.GetModuleId(), id: id}
	return &ca
}

func (s *commandArgumentStruct) GetId() byte {
	return s.id
}

func (s *commandArgumentStruct) SetId(id byte) {
	s.id = id
}

func (s *commandArgumentStruct) GetModuleId() byte {
	return s.moduleId
}

func (s *commandArgumentStruct) SetModuleId(id byte) {
	s.moduleId = id
}

func (s *commandArgumentStruct) GetCommandId() byte {
	return s.commandId
}

func (s *commandArgumentStruct) SetCommandId(id byte) {
	s.commandId = id
}

func (s *commandArgumentStruct) GetType() DataType {
	return s.dataType
}

func (s *commandArgumentStruct) SetType(dataType DataType) {
	s.dataType = dataType
}

func (s *commandArgumentStruct) GetMin() (min interface{}, ok bool) {
	min = s.min
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_MIN)
	return
}

func (s *commandArgumentStruct) SetMin(min interface{}) {
	s.min = min
	s.setFlag(COMMAND_ARGUMENT_FLAGS_MIN, true)
}

func (s *commandArgumentStruct) GetMax() (max interface{}, ok bool) {
	max = s.max
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_MAX)
	return
}

func (s *commandArgumentStruct) SetMax(max interface{}) {
	s.max = max
	s.setFlag(COMMAND_ARGUMENT_FLAGS_MAX, true)
}

func (s *commandArgumentStruct) GetList() (list []NameValue, ok bool) {
	list = s.list
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_LIST)
	return
}

func (s *commandArgumentStruct) SetList(list []NameValue) {
	s.list = list
	s.setFlag(COMMAND_ARGUMENT_FLAGS_LIST, true)
}

func (s *commandArgumentStruct) GetRequired() (required byte, ok bool) {
	required = s.required
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_REQUIRED)
	return
}

func (s *commandArgumentStruct) SetRequired(required byte) {
	s.required = required
	s.setFlag(COMMAND_ARGUMENT_FLAGS_REQUIRED, true)
}

func (s *commandArgumentStruct) GetName() (name string, ok bool) {
	name = s.name
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_NAME)
	return
}

func (s *commandArgumentStruct) SetName(name string) {
	s.name = name
	s.setFlag(COMMAND_ARGUMENT_FLAGS_NAME, true)
}

func (s *commandArgumentStruct) GetDesc() (desc string, ok bool) {
	desc = s.desc
	ok = s.hasFlag(COMMAND_ARGUMENT_FLAGS_DESCRIPTION)
	return
}

func (s *commandArgumentStruct) SetDesc(desc string) {
	s.desc = desc
	s.setFlag(COMMAND_ARGUMENT_FLAGS_DESCRIPTION, true)
}
