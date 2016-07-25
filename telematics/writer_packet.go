package telematics

import (
	"container/list"
	"encoding/binary"
	"fmt"
)

func (w *TelematicsWriter) WritePacket(packet interface{}) {
	switch p := packet.(type) {
	case *requestStruct:
		w.writeRequest(p)
	case Response:
		w.writeResponse(p)
	default:
		panic("packet type not defined")
	}

	crc := w.checksum.Compute()
	binary.Write(w.writer, binary.LittleEndian, crc)
}

func (w *TelematicsWriter) writeResponse(p Response) {
	binary.Write(w.writer, binary.LittleEndian, PACKET_TYPE_RESPONSE)
	binary.Write(w.writer, binary.LittleEndian, p.Sequence)
	binary.Write(w.writer, binary.LittleEndian, p.Flags)
}

func (w *TelematicsWriter) writeRequest(p *requestStruct) {
	binary.Write(w.writer, binary.LittleEndian, PACKET_TYPE_REQUEST)
	binary.Write(w.writer, binary.LittleEndian, byte(0x02))
	binary.Write(w.writer, binary.LittleEndian, p.sequence)
	binary.Write(w.writer, binary.LittleEndian, p.timestamp)
	w.writeSections(p.sections)
}

func (w *TelematicsWriter) writeSections(ss [10]interface{}) {
	for i, section := range ss {
		if section == nil {
			continue
		}
		if t, ok := sectionType(byte(i)); ok {
			switch t {
			case SECTION_IDENTIFICATION:
				binary.Write(w.writer, binary.LittleEndian, byte(t))
				w.writeIdentification(section.(*identificationSection))
			case SECTION_AUTHENTICATION:
				binary.Write(w.writer, binary.LittleEndian, byte(t))
				w.writeAuthentication(section.(*authenticationSection))
			case SECTION_SUPPORTED:
				binary.Write(w.writer, binary.LittleEndian, byte(t))
				w.writeSupported(section.(*supportedSection))

			case SECTION_MODULE:
				modules := section.(*list.List)
				for e := modules.Front(); e != nil; e = e.Next() {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeModule(e.Value.(*moduleSection))
				}
			case SECTION_MODULE_PROPERTY:
				properties := section.(*list.List)
				for e := properties.Front(); e != nil; e = e.Next() {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeModuleProperty(e.Value.(*modulePropertySection))
				}
			case SECTION_COMMAND:
				commands := section.(*list.List)
				for e := commands.Front(); e != nil; e = e.Next() {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeCommand(e.Value.(*commandSection))
				}
			case SECTION_COMMAND_ARGUMENT:
				args := section.(*list.List)
				for e := args.Front(); e != nil; e = e.Next() {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeCommandArgument(e.Value.(*commandArgumentStruct))
				}

			case SECTION_MODULE_PROPERTY_VALUE:
				for _, pv := range section.([]ModulePropertyValue) {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeModulePropertyValue(pv.(*modulePropertyValueSection))
				}
			case SECTION_MODULE_PROPERTY_DISABLED:
				for _, pd := range section.([]*ModulePropertyDisable) {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeModulePropertyDisable(pd)
				}
			case SECTION_COMMAND_EXECUTE:
				for _, ce := range section.([]CommandExecute) {
					binary.Write(w.writer, binary.LittleEndian, byte(t))
					w.writeCommandExecute(ce.(*commandExecuteStruct))
				}
			default:
				panic(fmt.Sprintf("section type %v not defined", t))
			}
		} else {
			panic("section type is not defined")
		}
	}

	binary.Write(w.writer, binary.LittleEndian, SECTION_ENDOFPAYLOAD)
}

func (w *TelematicsWriter) writeIdentification(s *identificationSection) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)
	if code, ok := s.GetCode(); ok {
		binary.Write(w.writer, binary.LittleEndian, code)
	}
	if code, ok := s.GetCodeText(); ok {
		w.writeString(code)
	}
	if deviceType, ok := s.GetDeviceType(); ok {
		binary.Write(w.writer, binary.LittleEndian, deviceType)
	}
	if firmware, ok := s.GetFirmware(); ok {
		binary.Write(w.writer, binary.LittleEndian, firmware)
	}
	if hardware, ok := s.GetHardware(); ok {
		binary.Write(w.writer, binary.LittleEndian, hardware)
	}
	if hash, ok := s.GetHash(); ok {
		binary.Write(w.writer, binary.LittleEndian, hash)
	}
}

func (w *TelematicsWriter) writeAuthentication(s *authenticationSection) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)

	if id, ok := s.GetIdentifier(); ok {
		w.writeString(id)
	}
	if secret, ok := s.GetSecret(); ok {
		w.writeBytes(secret)
	}
}

func (w *TelematicsWriter) writeModule(s *moduleSection) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)
	binary.Write(w.writer, binary.LittleEndian, s.Id)
	if name, ok := s.GetName(); ok {
		w.writeString(name)
	}
	if desc, ok := s.GetDesc(); ok {
		w.writeString(desc)
	}
}

func (w *TelematicsWriter) writeModuleProperty(s *modulePropertySection) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)
	binary.Write(w.writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.writer, binary.LittleEndian, s.Id)
	binary.Write(w.writer, binary.LittleEndian, s.Type)

	if min, ok := s.GetMin(); ok {
		w.writeData(min, s.Type)
	}
	if max, ok := s.GetMax(); ok {
		w.writeData(max, s.Type)
	}
	if list, ok := s.GetList(); ok {
		w.writeNameList(list, s.Type)
	}
	if access, ok := s.GetAccess(); ok {
		binary.Write(w.writer, binary.LittleEndian, access)
	}
	if name, ok := s.GetName(); ok {
		w.writeString(name)
	}
	if desc, ok := s.GetDesc(); ok {
		w.writeString(desc)
	}
}

func (w *TelematicsWriter) writeModulePropertyValue(s *modulePropertyValueSection) {
	binary.Write(w.writer, binary.LittleEndian, s.moduleId)
	binary.Write(w.writer, binary.LittleEndian, byte(len(s.values)))

	var dataType DataType
	for id, v := range s.values {
		p := w.conf.properties[s.GetModuleId()][id]
		dataType = p.GetType()
		if dataType == NotSet {
			panic("type not set")
		}

		binary.Write(w.writer, binary.LittleEndian, id)
		w.writeData(v, dataType)
	}
}

func (w *TelematicsWriter) writeModulePropertyDisable(s *ModulePropertyDisable) {
	binary.Write(w.writer, binary.LittleEndian, byte(len(s.DisabledProperties)))
	for id, v := range s.DisabledProperties {
		binary.Write(w.writer, binary.LittleEndian, id)
		w.writeData(v, Byte)
	}
}

func (w *TelematicsWriter) writeCommand(s *commandSection) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)
	binary.Write(w.writer, binary.LittleEndian, s.moduleId)
	binary.Write(w.writer, binary.LittleEndian, s.id)
	if name, ok := s.GetName(); ok {
		w.writeString(name)
	}
	if desc, ok := s.GetDesc(); ok {
		w.writeString(desc)
	}
}

func (w *TelematicsWriter) writeCommandArgument(s *commandArgumentStruct) {
	binary.Write(w.writer, binary.LittleEndian, s.flags)
	binary.Write(w.writer, binary.LittleEndian, s.moduleId)
	binary.Write(w.writer, binary.LittleEndian, s.commandId)
	binary.Write(w.writer, binary.LittleEndian, s.id)
	binary.Write(w.writer, binary.LittleEndian, s.dataType)
	if min, ok := s.GetMin(); ok {
		w.writeData(min, s.dataType)
	}
	if max, ok := s.GetMax(); ok {
		w.writeData(max, s.dataType)
	}
	if list, ok := s.GetList(); ok {
		w.writeNameList(list, s.dataType)
	}
	if required, ok := s.GetRequired(); ok {
		binary.Write(w.writer, binary.LittleEndian, required)
	}
	if name, ok := s.GetName(); ok {
		w.writeString(name)
	}
	if desc, ok := s.GetDesc(); ok {
		w.writeString(desc)
	}
}

func (w *TelematicsWriter) writeCommandExecute(s *commandExecuteStruct) {
	if _, ok := w.conf.commands[s.ModuleId]; !ok {
		panic(fmt.Sprintf("command with id %v not found", s.CommandId))
	}

	var args map[byte]CommandArgument
	if margs, ok := w.conf.arguments[s.ModuleId]; ok {
		if cargs, ok := margs[s.CommandId]; ok {
			args = cargs
		}
	}

	if args == nil {
		panic("arguments for command not found")
	}

	binary.Write(w.writer, binary.LittleEndian, s.getFlags())
	binary.Write(w.writer, binary.LittleEndian, s.ModuleId)
	binary.Write(w.writer, binary.LittleEndian, s.CommandId)
	binary.Write(w.writer, binary.LittleEndian, byte(len(s.Arguments)))

	for id, v := range s.Arguments {
		if arg, ok := args[id]; ok {
			binary.Write(w.writer, binary.LittleEndian, id)
			w.writeData(v, arg.GetType())
		} else {
			panic("argument not found")
		}
	}
}

func (w *TelematicsWriter) writeSupported(s *supportedSection) {
	w.writeBytes(s.Get())
}
