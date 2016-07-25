package telematics

type commandExecuteStruct struct {
	baseSection

	ModuleId  byte
	CommandId byte
	Arguments map[byte]interface{}
}

type CommandExecute interface {
	Section

	GetModuleId() byte
	GetCommandId() byte

	GetArgument(byte) (interface{}, bool)
	SetArgument(byte, interface{})
}

func NewCommandExecute(c Command) CommandExecute {
	v := commandExecuteStruct{ModuleId: c.GetModuleId(), CommandId: c.GetId(), Arguments: make(map[byte]interface{})}
	return &v
}

func (s *commandExecuteStruct) GetModuleId() byte {
	return s.ModuleId
}

func (s *commandExecuteStruct) GetCommandId() byte {
	return s.CommandId
}

func (s *commandExecuteStruct) GetArgument(id byte) (v interface{}, ok bool) {
	v, ok = s.Arguments[id]
	return
}

func (s *commandExecuteStruct) SetArgument(id byte, v interface{}) {
	s.Arguments[id] = v
}
