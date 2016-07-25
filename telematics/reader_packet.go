package telematics

import (
	"encoding/binary"
	"fmt"
)

const (
	PACKET_TYPE_REQUEST  byte = 0xAA
	PACKET_TYPE_RESPONSE byte = 0xCC
)

func (r *TelematicsReader) ReadPacket() interface{} {
	var pt byte
	if err := binary.Read(r.reader, binary.BigEndian, &pt); err != nil {
		panic(err)
	}

	switch pt {
	case PACKET_TYPE_REQUEST:
		return r.readRequest()
	case PACKET_TYPE_RESPONSE:
		return r.readResponse()
	default:
		panic(fmt.Sprintf("packet type %x not supported", pt))
	}
}

func (r *TelematicsReader) readRequest() Request {
	req := requestStruct{readonly: true}
	var v byte
	binary.Read(r.reader, binary.LittleEndian, &v)
	if v != 2 {
		panic("version not supported")
	}

	if err := binary.Read(r.reader, binary.LittleEndian, &req.sequence); err != nil {
		panic("sequence")
	}
	if err := binary.Read(r.reader, binary.LittleEndian, &req.timestamp); err != nil {
		panic("timestamp")
	}
	r.readSections(&req)

	var crc byte
	if err := binary.Read(r.reader, binary.LittleEndian, &crc); err != nil {
		panic("crc")
	}

	delta := r.checksum.Compute()
	if delta != 0 {
		panic(fmt.Sprintf("checksum %X", crc))
	}

	return &req
}

func (r *TelematicsReader) readResponse() *Response {
	p := Response{}
	if err := binary.Read(r.reader, binary.LittleEndian, &p.Sequence); err != nil {
		panic("sequence")
	}
	if err := binary.Read(r.reader, binary.LittleEndian, &p.Flags); err != nil {
		panic("flags")
	}
	if err := binary.Read(r.reader, binary.LittleEndian, &p.Crc); err != nil {
		panic("checksum")
	}

	return &p
}

func (r *TelematicsReader) readSections(req *requestStruct) {
	exit := false
	for !exit {
		var t SectionType
		binary.Read(r.reader, binary.LittleEndian, &t)
		switch t {
		case SECTION_ENDOFPAYLOAD:
			exit = true
		case SECTION_IDENTIFICATION:
			{
				// проверка на уникальность секции
				if id, ok := req.Identification(); !ok {
					r.readIdentification(id.(*identificationSection))
				} else {
					panic("identification exists")
				}
			}
		case SECTION_AUTHENTICATION:
			{
				// проверка на уникальность секции
				if auth, ok := req.Authentication(); !ok {
					r.readAuthentication(auth.(*authenticationSection))
				} else {
					panic("authentication exists")
				}
			}
		case SECTION_SUPPORTED:
			{
				// проверка на уникальность секции
				if sup, ok := req.Supported(); !ok {
					r.readSupported(sup.(*supportedSection))
				} else {
					panic("supported exists")
				}
			}

		case SECTION_MODULE:
			{
				m := &moduleSection{}
				r.readModule(m)
				req.pushConf(m)
			}
		case SECTION_MODULE_PROPERTY:
			{
				mp := &modulePropertySection{}
				r.readModuleProperty(mp)
				req.pushConf(mp)
			}
		case SECTION_COMMAND:
			{
				c := &commandSection{}
				r.readCommand(c)
				req.pushConf(c)
			}
		case SECTION_COMMAND_ARGUMENT:
			{
				ca := &commandArgumentStruct{}
				r.readCommandArgument(ca)
				req.pushConf(ca)
			}

		case SECTION_MODULE_PROPERTY_DISABLED:
			{
				pd := &ModulePropertyDisable{DisabledProperties: make(map[byte]byte)}
				r.readModulePropertyDisable(pd)
				req.section(pd)
			}
		case SECTION_MODULE_PROPERTY_VALUE:
			{
				pv := &modulePropertyValueSection{}
				r.readModulePropertyValue(pv)
				req.section(pv)
			}
		case SECTION_COMMAND_EXECUTE:
			{
				ce := &commandExecuteStruct{}
				r.readCommandExecute(ce)
				req.section(ce)
			}
		default:
			panic("section not found")
		}
	}
}

func (r *TelematicsReader) readIdentification(s *identificationSection) {
	binary.Read(r.reader, binary.LittleEndian, &s.flags)
	codeText, code := false, false
	for _, flag := range s.getFlags() {
		switch flag {
		case IDENTIFICATION_FLAGS_CODE:
			binary.Read(r.reader, binary.LittleEndian, &s.code)
			code = true
		case IDENTIFICATION_FLAGS_CODETEXT:
			s.codeText = r.readString()
			codeText = true
		case IDENTIFICATION_FLAGS_DEVICETYPE:
			binary.Read(r.reader, binary.LittleEndian, &s.deviceType)
		case IDENTIFICATION_FLAGS_FIRMWARE:
			binary.Read(r.reader, binary.LittleEndian, &s.firmware)
		case IDENTIFICATION_FLAGS_HARDWARE:
			binary.Read(r.reader, binary.LittleEndian, &s.hardware)
		case IDENTIFICATION_FLAGS_DEVICEHASH:
			binary.Read(r.reader, binary.LittleEndian, &s.hash)
		default:
			panic("flag not supported")
		}
	}

	if code && codeText {
		panic("model can't contain both Code and CodeText field")
	}

	if !code && !codeText {
		panic("model must contain Code or CodeText field")
	}
}

func (r *TelematicsReader) readAuthentication(s *authenticationSection) {
	binary.Read(r.reader, binary.LittleEndian, &s.flags)
	for _, flag := range s.getFlags() {
		switch flag {
		case AUTHENTICATION_FLAGS_IDENTIFIER:
			s.Identifier = r.readString()
		case AUTHENTICATION_FLAGS_SECRET:
			s.Secret = r.readBytes()
		default:
			panic("flag not supported")
		}
	}
}

func (r *TelematicsReader) readModule(m *moduleSection) {
	binary.Read(r.reader, binary.LittleEndian, &m.flags)
	binary.Read(r.reader, binary.LittleEndian, &m.Id)
	for _, flag := range m.getFlags() {
		switch flag {
		case MODULE_FLAGS_NAME:
			m.Name = r.readString()
		case MODULE_FLAGS_DESCRIPTION:
			m.Description = r.readString()
		default:
			panic(fmt.Sprintf("flag not supported %v from %v", flag, m.flags))
		}
	}
}

func (r *TelematicsReader) readModuleProperty(mp *modulePropertySection) {
	binary.Read(r.reader, binary.LittleEndian, &mp.flags)
	binary.Read(r.reader, binary.LittleEndian, &mp.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &mp.Id)
	binary.Read(r.reader, binary.LittleEndian, &mp.Type)
	for _, flag := range mp.getFlags() {
		switch flag {
		case MODULE_PROPERTY_FLAGS_MIN:
			binary.Read(r.reader, binary.LittleEndian, &mp.Min)
		case MODULE_PROPERTY_FLAGS_MAX:
			binary.Read(r.reader, binary.LittleEndian, &mp.Max)
		case MODULE_PROPERTY_FLAGS_LIST:
			mp.List = r.readNameValues(mp.Type)
		case MODULE_PROPERTY_FLAGS_ACCESS:
			binary.Read(r.reader, binary.LittleEndian, &mp.Access)
		case MODULE_PROPERTY_FLAGS_NAME:
			mp.Name = r.readString()
		case MODULE_PROPERTY_FLAGS_DESCRIPTION:
			mp.Desc = r.readString()
		}
	}

	if _, e := r.conf.properties[mp.ModuleId]; !e {
		r.conf.properties[mp.ModuleId] = make(map[byte]ModuleProperty)
	}

	if _, e := r.conf.properties[mp.ModuleId][mp.Id]; !e {
		r.conf.properties[mp.ModuleId][mp.Id] = mp
	} else {
		panic("module property exists")
	}
}

func (r *TelematicsReader) readModulePropertyValue(s *modulePropertyValueSection) {
	if len(r.conf.properties) == 0 {
		panic("no properties")
	}

	s.values = make(map[byte]interface{})
	binary.Read(r.reader, binary.LittleEndian, &s.moduleId)
	var c, i byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i = 0; i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)

		var p ModuleProperty
		if mpmap, e := r.conf.properties[s.moduleId]; e {
			if mp, e := mpmap[id]; e {
				p = mp
			}
		}

		if p == nil {
			panic("module property not found")
		}

		s.values[id] = r.readData(p.GetType())
	}
}

func (r *TelematicsReader) readModulePropertyDisable(s *ModulePropertyDisable) {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)
		v := r.readData(Byte)
		s.DisabledProperties[id] = v.(byte)
	}
}

func (r *TelematicsReader) readCommand(c *commandSection) {
	binary.Read(r.reader, binary.LittleEndian, &c.flags)
	binary.Read(r.reader, binary.LittleEndian, &c.moduleId)
	binary.Read(r.reader, binary.LittleEndian, &c.id)
	for _, flag := range c.getFlags() {
		switch flag {
		case COMMAND_FLAGS_NAME:
			c.name = r.readString()
		case COMMAND_FLAGS_DESCRIPTION:
			c.description = r.readString()
		default:
			panic("flag not supported")
		}
	}

	if _, e := r.conf.commands[c.moduleId]; !e {
		r.conf.commands[c.moduleId] = make(map[byte]Command)
	}
	if _, e := r.conf.commands[c.moduleId][c.id]; !e {
		r.conf.commands[c.moduleId][c.id] = c
	} else {
		panic("module command exists")
	}
}

func (r *TelematicsReader) readCommandArgument(ca *commandArgumentStruct) {
	binary.Read(r.reader, binary.LittleEndian, &ca.flags)
	binary.Read(r.reader, binary.LittleEndian, &ca.moduleId)
	binary.Read(r.reader, binary.LittleEndian, &ca.commandId)
	binary.Read(r.reader, binary.LittleEndian, &ca.id)
	binary.Read(r.reader, binary.LittleEndian, &ca.dataType)

	for _, flag := range ca.getFlags() {
		switch flag {
		case COMMAND_ARGUMENT_FLAGS_MIN:
			ca.min = r.readData(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_MAX:
			ca.max = r.readData(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_LIST:
			ca.list = r.readNameValues(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_REQUIRED:
			binary.Read(r.reader, binary.LittleEndian, &ca.required)
		case COMMAND_ARGUMENT_FLAGS_NAME:
			ca.name = r.readString()
		case COMMAND_ARGUMENT_FLAGS_DESCRIPTION:
			ca.desc = r.readString()
		}
	}

	if _, e := r.conf.arguments[ca.moduleId]; !e {
		r.conf.arguments[ca.moduleId] = make(map[byte]map[byte]CommandArgument)
	}
	if _, e := r.conf.arguments[ca.moduleId][ca.commandId]; !e {
		r.conf.arguments[ca.moduleId][ca.commandId] = make(map[byte]CommandArgument)
	}
	if _, e := r.conf.arguments[ca.moduleId][ca.commandId][ca.id]; !e {
		r.conf.arguments[ca.moduleId][ca.commandId][ca.id] = ca
	} else {
		panic("command argument exists")
	}
}

func (r *TelematicsReader) readCommandExecute(ce *commandExecuteStruct) {
	if len(r.conf.commands) == 0 {
		panic("no commands")
	}

	ce.Arguments = make(map[byte]interface{})

	binary.Read(r.reader, binary.LittleEndian, &ce.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &ce.CommandId)
	if _, e := r.conf.commands[ce.ModuleId]; !e {
		panic("no module commands")
	} else if _, e := r.conf.commands[ce.ModuleId][ce.CommandId]; !e {
		panic("no command in module")
	}

	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)

		var arg CommandArgument
		if _, e := r.conf.arguments[ce.ModuleId]; e {
			if _, e := r.conf.arguments[ce.ModuleId][ce.CommandId]; e {
				if carg, e := r.conf.arguments[ce.ModuleId][ce.CommandId][id]; e {
					arg = carg
				}
			}
		}
		if arg == nil {
			panic("unable to read commandArgument")
		}

		ce.Arguments[id] = r.readData(arg.GetType())
	}
}

func (r *TelematicsReader) readSupported(s *supportedSection) {
	codes := r.readBytes()
	s.Set(codes)
}
