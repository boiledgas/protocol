package telematics

// command flags
const (
	COMMAND_FLAGS_NAME        byte = 0x01
	COMMAND_FLAGS_DESCRIPTION byte = 0x02
)

type commandSection struct {
	baseSection

	moduleId    byte
	id          byte
	name        string
	description string
}

type Command interface {
	Section

	GetModuleId() byte

	GetId() byte
	SetId(byte)

	GetName() (string, bool)
	SetName(string)

	GetDesc() (string, bool)
	SetDesc(string)
}

func NewCommand(id byte, m Module) Command {
	c := commandSection{moduleId: m.GetId(), id: id}
	return &c
}

func (c *commandSection) GetModuleId() byte {
	return c.moduleId
}

func (c *commandSection) GetId() byte {
	return c.id
}

func (c *commandSection) SetId(id byte) {
	c.id = id
}

func (c *commandSection) GetName() (name string, ok bool) {
	name = c.name
	ok = c.hasFlag(COMMAND_FLAGS_NAME)
	return
}

func (c *commandSection) SetName(name string) {
	c.name = name
	c.setFlag(COMMAND_FLAGS_NAME, true)
}

func (c *commandSection) GetDesc() (desc string, ok bool) {
	desc = c.description
	ok = c.hasFlag(COMMAND_FLAGS_DESCRIPTION)
	return
}

func (c *commandSection) SetDesc(desc string) {
	c.description = desc
	c.setFlag(COMMAND_FLAGS_DESCRIPTION, true)
}
