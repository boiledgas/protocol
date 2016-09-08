package telematics

import (
	"github.com/boiledgas/protocol/telematics/section"
	"bytes"
)

// Device configuration
type Configuration struct {
	Hash       byte
	Modules    []section.Module
	Properties []section.ModuleProperty
	Commands   []section.Command
	Arguments  []section.CommandArgument
}

func (c *Configuration) GetProperty(moduleId byte, propertyId byte, property *section.ModuleProperty) (ok bool) {
	ok = false
	for _, p := range c.Properties {
		if p.ModuleId == moduleId && p.Id == propertyId {
			*property = p
			ok = true
			break
		}
	}
	return
}

func (c *Configuration) GetArgument(moduleId byte, commandId byte, argumentId byte, argument *section.CommandArgument) (ok bool) {
	ok = false
	for _, arg := range c.Arguments {
		if arg.ModuleId == moduleId && arg.CommandId == commandId && arg.Id == argumentId {
			*argument = arg
			ok = true
			break
		}
	}
	return
}

func (c *Configuration) Validate() bool {
	var modules map[byte]*section.Module
	var module section.Module
	for i := 0; i < len(c.Modules); i++ {
		module = c.Modules[i]
		if _, ok := modules[module.Id]; ok {
			return false
		}
		modules[module.Id] = &c.Modules[i]
		for j := i; j < len(c.Modules); j++ {
			if module.Name == c.Modules[j].Name {
				return false
			}
		}
	}
	var properties map[byte]map[byte]*section.ModuleProperty
	var property section.ModuleProperty
	for i := 0; i < len(c.Properties); i++ {
		property = c.Properties[i]
		if _, ok := properties[property.ModuleId]; !ok {
			properties[property.ModuleId] = make(map[byte]*section.ModuleProperty)
		}
		if _, ok := properties[property.ModuleId][property.Id]; ok {
			return false
		}
		properties[property.ModuleId][property.Id] = &c.Properties[i]
	}

	var commands map[byte]map[byte]*section.Command
	var command section.Command
	for i := 0; i < len(c.Commands); i++ {
		command = c.Commands[i]
		if _, ok := commands[command.ModuleId]; !ok {
			commands[command.ModuleId] = make(map[byte]*section.Command)
		}
		if _, ok := commands[command.ModuleId][command.Id]; ok {
			return false
		}
		commands[command.ModuleId][command.Id] = &c.Commands[i]
	}

	var arguments map[byte]map[byte]map[byte]*section.CommandArgument
	var argument section.CommandArgument
	for i := 0; i < len(c.Arguments); i++ {
		argument = c.Arguments[i]
		if _, ok := arguments[argument.ModuleId]; !ok {
			arguments[argument.ModuleId] = make(map[byte]map[byte]*section.CommandArgument)
		}
		if _, ok := arguments[argument.ModuleId][argument.CommandId]; !ok {
			arguments[argument.ModuleId][argument.CommandId] = make(map[byte]*section.CommandArgument)
		}
		if _, ok := arguments[argument.ModuleId][argument.CommandId][argument.Id]; ok {
			return false
		}
		arguments[argument.ModuleId][argument.CommandId][argument.Id] = &c.Arguments[i]
	}
	return true
}

func (s Configuration) String() string {
	var buf bytes.Buffer
	buf.WriteString("Configuration: {")
	buf.WriteString("\nModules:[")
	for _, m := range s.Modules {
		buf.WriteString(m.String())
	}
	buf.WriteString("]")
	buf.WriteString("\nProperties:[")
	for _, p := range s.Properties {
		buf.WriteString(p.String())
	}
	buf.WriteString("]")
	buf.WriteString("\nCommands:[")
	for _, c := range s.Commands {
		buf.WriteString(c.String())
	}
	buf.WriteString("]")
	buf.WriteString("\nArguments:[")
	for _, a := range s.Arguments {
		buf.WriteString(a.String())
	}
	buf.WriteString("]")
	return buf.String()
}
