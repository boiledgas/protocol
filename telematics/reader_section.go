package telematics

import (
	"encoding/binary"
	"fmt"
	"protocol/telematics/section"
)

const (
	PACKET_TYPE_REQUEST  byte = 0xAA
	PACKET_TYPE_RESPONSE byte = 0xCC
)

func (r *TelematicsReader) ReadPacket() interface{} {
	var pt byte
	if err := binary.Read(r.Reader, binary.BigEndian, &pt); err != nil {
		panic(err)
	}

	switch pt {
	case PACKET_TYPE_REQUEST:
		request := Request{}
		return r.ReadRequest(&request)
	case PACKET_TYPE_RESPONSE:
		response := Response{}
		return r.ReadResponse(&response)
	default:
		panic(fmt.Sprintf("packet type %x not supported", pt))
	}
}

func (r *TelematicsReader) ReadRequest(req *Request) {
	var v byte
	binary.Read(r.Reader, binary.LittleEndian, &v)
	if v != 2 {
		panic("version not supported")
	}

	if err := binary.Read(r.Reader, binary.LittleEndian, &req.Sequence); err != nil {
		panic("sequence")
	}
	if err := binary.Read(r.Reader, binary.LittleEndian, &req.Timestamp); err != nil {
		panic("timestamp")
	}

	var t SectionType
sections:
	for {
		binary.Read(r.Reader, binary.LittleEndian, &t)
		switch t {
		case SECTION_ENDOFPAYLOAD:
			break sections
		case SECTION_IDENTIFICATION:
			{
				if req.Has(SECTION_IDENTIFICATION) {
					panic("identification exists")
				}
				r.ReadIdentification(&req.Id)
			}
		case SECTION_AUTHENTICATION:
			{
				if req.Has(SECTION_AUTHENTICATION) {
					panic("authentication exists")
				}
				r.ReadAuthentication(&req.Auth)
			}
		case SECTION_SUPPORTED:
			{
				if req.Has(SECTION_SUPPORTED) {
					panic("supported exists")
				}
				r.ReadSupported(&req.Sup)
			}

		case SECTION_MODULE:
			{
				m := &Module{}
				r.ReadModule(m)
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

	var crc byte
	if err := binary.Read(r.Reader, binary.LittleEndian, &crc); err != nil {
		panic("crc")
	}

	delta := r.checksum.Compute()
	if delta != 0 {
		panic(fmt.Sprintf("checksum %X", crc))
	}
}

func (r *TelematicsReader) ReadResponse(response *Response) {
	if err := binary.Read(r.Reader, binary.LittleEndian, &response.Sequence); err != nil {
		panic("sequence")
	}
	if err := binary.Read(r.Reader, binary.LittleEndian, &response.Flags); err != nil {
		panic("flags")
	}
	if err := binary.Read(r.Reader, binary.LittleEndian, &response.Crc); err != nil {
		panic("checksum")
	}
}

func (r *TelematicsReader) ReadIdentification(s *section.Identification) {
	binary.Read(r.Reader, binary.LittleEndian, &s.Flags8)
	codeText, code := false, false
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		switch flag {
		case section.IDENTIFICATION_FLAGS_CODE:
			binary.Read(r.Reader, binary.LittleEndian, &s.code)
			code = true
		case section.IDENTIFICATION_FLAGS_CODETEXT:
			s.codeText = r.ReadString()
			codeText = true
		case section.IDENTIFICATION_FLAGS_DEVICETYPE:
			binary.Read(r.Reader, binary.LittleEndian, &s.deviceType)
		case section.IDENTIFICATION_FLAGS_FIRMWARE:
			binary.Read(r.Reader, binary.LittleEndian, &s.firmware)
		case section.IDENTIFICATION_FLAGS_HARDWARE:
			binary.Read(r.Reader, binary.LittleEndian, &s.hardware)
		case section.IDENTIFICATION_FLAGS_DEVICEHASH:
			binary.Read(r.Reader, binary.LittleEndian, &s.hash)
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

func (r *TelematicsReader) ReadAuthentication(s *Authentication) {
	binary.Read(r.Reader, binary.LittleEndian, &s.Flags8)
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		switch flag {
		case AUTHENTICATION_FLAGS_IDENTIFIER:
			s.Identifier = r.ReadString()
		case AUTHENTICATION_FLAGS_SECRET:
			s.Secret = r.ReadBytes()
		default:
			panic("flag not supported")
		}
	}
}

func (r *TelematicsReader) ReadModule(s *section.Module) {
	binary.Read(r.Reader, binary.LittleEndian, &s.Flags8)
	binary.Read(r.Reader, binary.LittleEndian, &s.Id)
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		switch flag {
		case section.MODULE_FLAGS_NAME:
			s.Name = r.ReadString()
		case section.MODULE_FLAGS_DESCRIPTION:
			s.Description = r.ReadString()
		default:
			panic(fmt.Sprintf("flag not supported %v from %v", flag, flags))
		}
	}
}

func (r *TelematicsReader) readModuleProperty(mp *section.ModuleProperty) {
	binary.Read(r.Reader, binary.LittleEndian, &mp.flags)
	binary.Read(r.Reader, binary.LittleEndian, &mp.ModuleId)
	binary.Read(r.Reader, binary.LittleEndian, &mp.Id)
	binary.Read(r.Reader, binary.LittleEndian, &mp.Type)
	for _, flag := range mp.getFlags() {
		switch flag {
		case MODULE_PROPERTY_FLAGS_MIN:
			binary.Read(r.Reader, binary.LittleEndian, &mp.Min)
		case MODULE_PROPERTY_FLAGS_MAX:
			binary.Read(r.Reader, binary.LittleEndian, &mp.Max)
		case MODULE_PROPERTY_FLAGS_LIST:
			mp.List = r.ReadNameValues(mp.Type)
		case MODULE_PROPERTY_FLAGS_ACCESS:
			binary.Read(r.Reader, binary.LittleEndian, &mp.Access)
		case MODULE_PROPERTY_FLAGS_NAME:
			mp.Name = r.ReadString()
		case MODULE_PROPERTY_FLAGS_DESCRIPTION:
			mp.Desc = r.ReadString()
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
	binary.Read(r.Reader, binary.LittleEndian, &s.moduleId)
	var c, i byte
	binary.Read(r.Reader, binary.LittleEndian, &c)
	var id byte
	for i = 0; i < c; i++ {
		binary.Read(r.Reader, binary.LittleEndian, &id)

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
	binary.Read(r.Reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.Reader, binary.LittleEndian, &id)
		v := r.readData(Byte)
		s.DisabledProperties[id] = v.(byte)
	}
}

func (r *TelematicsReader) readCommand(c *commandSection) {
	binary.Read(r.Reader, binary.LittleEndian, &c.flags)
	binary.Read(r.Reader, binary.LittleEndian, &c.moduleId)
	binary.Read(r.Reader, binary.LittleEndian, &c.id)
	for _, flag := range c.getFlags() {
		switch flag {
		case COMMAND_FLAGS_NAME:
			c.name = r.ReadString()
		case COMMAND_FLAGS_DESCRIPTION:
			c.description = r.ReadString()
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
	binary.Read(r.Reader, binary.LittleEndian, &ca.flags)
	binary.Read(r.Reader, binary.LittleEndian, &ca.moduleId)
	binary.Read(r.Reader, binary.LittleEndian, &ca.commandId)
	binary.Read(r.Reader, binary.LittleEndian, &ca.id)
	binary.Read(r.Reader, binary.LittleEndian, &ca.dataType)

	for _, flag := range ca.getFlags() {
		switch flag {
		case COMMAND_ARGUMENT_FLAGS_MIN:
			ca.min = r.readData(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_MAX:
			ca.max = r.readData(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_LIST:
			ca.list = r.ReadNameValues(ca.dataType)
		case COMMAND_ARGUMENT_FLAGS_REQUIRED:
			binary.Read(r.Reader, binary.LittleEndian, &ca.required)
		case COMMAND_ARGUMENT_FLAGS_NAME:
			ca.name = r.ReadString()
		case COMMAND_ARGUMENT_FLAGS_DESCRIPTION:
			ca.desc = r.ReadString()
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

	binary.Read(r.Reader, binary.LittleEndian, &ce.ModuleId)
	binary.Read(r.Reader, binary.LittleEndian, &ce.CommandId)
	if _, e := r.conf.commands[ce.ModuleId]; !e {
		panic("no module commands")
	} else if _, e := r.conf.commands[ce.ModuleId][ce.CommandId]; !e {
		panic("no command in module")
	}

	var c byte
	binary.Read(r.Reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.Reader, binary.LittleEndian, &id)

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

func (r *TelematicsReader) ReadSupported(s *Supported) {
	codes := r.ReadBytes()
	s.Set(codes)
}
