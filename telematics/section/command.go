package section

// command flags
const (
	COMMAND_FLAGS_NAME        byte = 0x01
	COMMAND_FLAGS_DESCRIPTION byte = 0x02
)

type Command struct {
	moduleId    byte
	id          byte
	name        string
	description string
}
