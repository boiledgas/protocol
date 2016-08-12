package section

type CommandExecute struct {
	ModuleId  byte
	CommandId byte
	Arguments map[byte]interface{}
}
