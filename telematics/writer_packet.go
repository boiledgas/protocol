package telematics

import (
	"encoding/binary"
	"fmt"
	"protocol/telematics/section"
	"protocol/telematics/value"
)

func (w *TelematicsWriter) WriteResponse(p *Response) {
	binary.Write(w.Writer, binary.LittleEndian, PACKET_TYPE_RESPONSE)
	binary.Write(w.Writer, binary.LittleEndian, p.Sequence)
	binary.Write(w.Writer, binary.LittleEndian, p.Flags)
	p.Crc = w.Checksum.Compute()
	binary.Write(w.Writer, binary.LittleEndian, p.Crc)
}

func (w *TelematicsWriter) WriteRequest(p *Request) {
	binary.Write(w.Writer, binary.LittleEndian, PACKET_TYPE_REQUEST)
	binary.Write(w.Writer, binary.LittleEndian, byte(0x02))
	binary.Write(w.Writer, binary.LittleEndian, p.Sequence)
	binary.Write(w.Writer, binary.LittleEndian, p.Timestamp)
	var flags [16]uint16
	p.Load(&flags)
	for _, flag := range flags {
		t := section.ToSectionType(flag)
		switch t {
		case section.SECTION_IDENTIFICATION:
			binary.Write(w.Writer, binary.LittleEndian, byte(t))
			w.WriteIdentification(&p.Id)
		case section.SECTION_AUTHENTICATION:
			binary.Write(w.Writer, binary.LittleEndian, byte(t))
			w.WriteAuthentication(&p.Auth)
		case section.SECTION_SUPPORTED:
			binary.Write(w.Writer, binary.LittleEndian, byte(t))
			w.writeSupported(&p.Sup)

		case section.SECTION_MODULE:
			for _, m := range p.Conf.Modules {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteModule(&m)
			}
		case section.SECTION_MODULE_PROPERTY:
			for _, p := range p.Conf.Properties {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteModuleProperty(&p)
			}
		case section.SECTION_COMMAND:
			for _, c := range p.Conf.Commands {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteCommand(&c)
			}
		case section.SECTION_COMMAND_ARGUMENT:
			for _, ca := range p.Conf.Arguments {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteCommandArgument(&ca)
			}

		case section.SECTION_MODULE_PROPERTY_VALUE:
			for _, pv := range p.Values {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteModulePropertyValue(&pv)
			}
		case section.SECTION_MODULE_PROPERTY_DISABLED:
			for _, d := range p.Disabled {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteModulePropertyDisable(&d)
			}
		case section.SECTION_COMMAND_EXECUTE:
			for _, e := range p.Executes {
				binary.Write(w.Writer, binary.LittleEndian, byte(t))
				w.WriteCommandExecute(&e)
			}
		default:
			panic(fmt.Sprintf("section type %v not defined", t))
		}
	}

	binary.Write(w.Writer, binary.LittleEndian, section.SECTION_ENDOFPAYLOAD)
	crc := w.Checksum.Compute()
	binary.Write(w.Writer, binary.LittleEndian, crc)
}

func (w *TelematicsWriter) WriteIdentification(s *section.Identification) {
	binary.Write(w.Writer, binary.LittleEndian, s.Flags8)
	if s.Has(section.IDENTIFICATION_FLAGS_CODE) {
		binary.Write(w.Writer, binary.LittleEndian, s.Code)
	}
	if s.Has(section.IDENTIFICATION_FLAGS_CODETEXT) {
		w.WriteString(s.CodeText)
	}
	if s.Has(section.IDENTIFICATION_FLAGS_DEVICETYPE) {
		binary.Write(w.Writer, binary.LittleEndian, s.Type)
	}
	if s.Has(section.IDENTIFICATION_FLAGS_FIRMWARE) {
		binary.Write(w.Writer, binary.LittleEndian, s.Firmware)
	}
	if s.Has(section.IDENTIFICATION_FLAGS_HARDWARE) {
		binary.Write(w.Writer, binary.LittleEndian, s.Hardware)
	}
	if s.Has(section.IDENTIFICATION_FLAGS_DEVICEHASH) {
		binary.Write(w.Writer, binary.LittleEndian, s.Hash)
	}
}

func (w *TelematicsWriter) WriteAuthentication(s *section.Authentication) {
	binary.Write(w.Writer, binary.LittleEndian, byte(s.Flags8))

	if s.Has(section.AUTHENTICATION_FLAGS_IDENTIFIER) {
		w.WriteString(s.Identifier)
	}
	if s.Has(section.AUTHENTICATION_FLAGS_SECRET) {
		w.WriteBytes(s.Secret)
	}
}

func (w *TelematicsWriter) WriteModule(s *section.Module) {
	binary.Write(w.Writer, binary.LittleEndian, s.Flags8)
	binary.Write(w.Writer, binary.LittleEndian, s.Id)
	if s.Has(section.MODULE_FLAGS_NAME) {
		w.WriteString(s.Name)
	}
	if s.Has(section.MODULE_FLAGS_DESCRIPTION) {
		w.WriteString(s.Description)
	}
}

func (w *TelematicsWriter) WriteModuleProperty(s *section.ModuleProperty) {
	binary.Write(w.Writer, binary.LittleEndian, s.Flags8)
	binary.Write(w.Writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.Writer, binary.LittleEndian, s.Id)
	binary.Write(w.Writer, binary.LittleEndian, s.Type)

	if s.Has(section.MODULE_PROPERTY_FLAGS_MIN) {
		w.WriteData(s.Min, s.Type)
	}
	if s.Has(section.MODULE_PROPERTY_FLAGS_MAX) {
		w.WriteData(s.Max, s.Type)
	}
	if s.Has(section.MODULE_PROPERTY_FLAGS_LIST) {
		w.WriteNameList(s.List, s.Type)
	}
	if s.Has(section.MODULE_PROPERTY_FLAGS_ACCESS) {
		binary.Write(w.Writer, binary.LittleEndian, s.Access)
	}
	if s.Has(section.MODULE_PROPERTY_FLAGS_NAME) {
		w.WriteString(s.Name)
	}
	if s.Has(section.MODULE_PROPERTY_FLAGS_DESCRIPTION) {
		w.WriteString(s.Desc)
	}
}

func (w *TelematicsWriter) WriteModulePropertyValue(s *section.ModulePropertyValue) {
	binary.Write(w.Writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.Writer, binary.LittleEndian, byte(len(s.Values)))

	var p section.ModuleProperty
	for id, v := range s.Values {
		if !w.Configuration.GetProperty(s.ModuleId, id, &p) {
			panic("property not found")
		}
		if p.Type == value.NotSet {
			panic("type not set")
		}

		binary.Write(w.Writer, binary.LittleEndian, id)
		w.WriteData(v, p.Type)
	}
}

func (w *TelematicsWriter) WriteModulePropertyDisable(s *section.ModulePropertyDisable) {
	binary.Write(w.Writer, binary.LittleEndian, byte(len(s.DisabledProperties)))
	for id, v := range s.DisabledProperties {
		binary.Write(w.Writer, binary.LittleEndian, id)
		w.WriteData(v, value.Byte)
	}
}

func (w *TelematicsWriter) WriteCommand(s *section.Command) {
	binary.Write(w.Writer, binary.LittleEndian, s.Flags8)
	binary.Write(w.Writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.Writer, binary.LittleEndian, s.Id)
	if s.Has(section.COMMAND_FLAGS_NAME) {
		w.WriteString(s.Name)
	}
	if s.Has(section.COMMAND_FLAGS_DESCRIPTION) {
		w.WriteString(s.Description)
	}
}

func (w *TelematicsWriter) WriteCommandArgument(s *section.CommandArgument) {
	binary.Write(w.Writer, binary.LittleEndian, s.Flags8)
	binary.Write(w.Writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.Writer, binary.LittleEndian, s.CommandId)
	binary.Write(w.Writer, binary.LittleEndian, s.Id)
	binary.Write(w.Writer, binary.LittleEndian, s.Type)
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_MIN) {
		w.WriteData(s.Min, s.Type)
	}
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_MAX) {
		w.WriteData(s.Max, s.Type)
	}
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_LIST) {
		w.WriteNameList(s.List, s.Type)
	}
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_REQUIRED) {
		binary.Write(w.Writer, binary.LittleEndian, s.Required)
	}
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_NAME) {
		w.WriteString(s.Name)
	}
	if s.Has(section.COMMAND_ARGUMENT_FLAGS_DESCRIPTION) {
		w.WriteString(s.Desc)
	}
}

func (w *TelematicsWriter) WriteCommandExecute(s *section.CommandExecute) {
	binary.Write(w.Writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.Writer, binary.LittleEndian, s.CommandId)
	binary.Write(w.Writer, binary.LittleEndian, byte(len(s.Arguments)))

	var arg section.CommandArgument
	for id, v := range s.Arguments {
		if w.Configuration.GetArgument(s.ModuleId, s.CommandId, id, &arg) {
			binary.Write(w.Writer, binary.LittleEndian, id)
			w.WriteData(v, arg.Type)
		} else {
			panic("argument not found")
		}
	}
}

func (w *TelematicsWriter) writeSupported(s *section.Supported) {
	var flags [16]uint16
	s.Load(&flags)
	var bytes []byte
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		bytes = append(bytes, byte(section.ToSectionType(flag)))
	}
	w.WriteBytes(bytes)
}
