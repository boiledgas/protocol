package telematics

import (
	"encoding/binary"
	"fmt"
	"protocol/telematics/section"
	"protocol/telematics/value"
	"receiver/errors"
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
		request := Request{}
		r.ReadRequest(&request)
		return request
	case PACKET_TYPE_RESPONSE:
		response := Response{}
		r.ReadResponse(&response)
		return response
	default:
		panic(fmt.Sprintf("packet type %x not supported", pt))
	}
}

func (r *TelematicsReader) ReadRequest(req *Request) (err error) {
	var v byte
	binary.Read(r.reader, binary.LittleEndian, &v)
	if v != 2 {
		panic("version not supported")
	}

	if err := binary.Read(r.reader, binary.LittleEndian, &req.Sequence); err != nil {
		panic("sequence")
	}
	if err := binary.Read(r.reader, binary.LittleEndian, &req.Timestamp); err != nil {
		panic("timestamp")
	}

	var t section.Type
sections:
	for {
		binary.Read(r.reader, binary.LittleEndian, &t)
		switch t {
		case section.SECTION_ENDOFPAYLOAD:
			break sections
		case section.SECTION_IDENTIFICATION:
			if req.Has(section.SECTION_IDENTIFICATION.Flag()) {
				err = errors.New("identification exists")
			}
			r.ReadIdentification(&req.Id)
		case section.SECTION_AUTHENTICATION:
			if req.Has(section.SECTION_AUTHENTICATION.Flag()) {
				err = errors.New("authentication exists")
			}
			r.ReadAuthentication(&req.Auth)
		case section.SECTION_SUPPORTED:
			if req.Has(section.SECTION_SUPPORTED.Flag()) {
				err = errors.New("supported exists")
			}
			r.ReadSupported(&req.Sup)
		case section.SECTION_MODULE:
			m := section.Module{}
			r.ReadModule(&m)
			req.Conf.Modules = append(req.Conf.Modules, m)
		case section.SECTION_MODULE_PROPERTY:
			mp := section.ModuleProperty{}
			r.ReadModuleProperty(&mp)
			req.Conf.Properties = append(req.Conf.Properties, mp)
		case section.SECTION_COMMAND:
			c := section.Command{}
			r.ReadCommand(&c)
			req.Conf.Commands = append(req.Conf.Commands, c)
		case section.SECTION_COMMAND_ARGUMENT:
			ca := section.CommandArgument{}
			r.ReadCommandArgument(&ca)
			req.Conf.Arguments = append(req.Conf.Arguments, ca)
		case section.SECTION_MODULE_PROPERTY_DISABLED:
			pd := section.ModulePropertyDisable{DisabledProperties: make(map[byte]byte)}
			r.ReadModulePropertyDisable(&pd)
			req.Disabled = append(req.Disabled, pd)
		case section.SECTION_MODULE_PROPERTY_VALUE:
			pv := section.ModulePropertyValue{}
			if err = r.ReadModulePropertyValue(&pv); err != nil {
				return
			}
			req.Values = append(req.Values, pv)
		case section.SECTION_COMMAND_EXECUTE:
			ce := section.CommandExecute{}
			r.ReadCommandExecute(&ce)
			req.Executes = append(req.Executes, ce)
		default:
			err = errors.New("section not found")
			return
		}
		req.Set(t.Flag(), true)
	}

	var crc byte
	if err = binary.Read(r.reader, binary.LittleEndian, &crc); err != nil {
		return
	}

	delta := r.checksum.Compute()
	if delta != 0 {
		err = errors.New("crc not valid")
	}
	return
}

func (r *TelematicsReader) ReadResponse(response *Response) (err error) {
	if err = binary.Read(r.reader, binary.LittleEndian, &response.Sequence); err != nil {
		panic("sequence")
	}
	if err = binary.Read(r.reader, binary.LittleEndian, &response.Flags); err != nil {
		panic("flags")
	}
	if err = binary.Read(r.reader, binary.LittleEndian, &response.Crc); err != nil {
		panic("checksum")
	}
	return
}

func (r *TelematicsReader) ReadIdentification(s *section.Identification) {
	binary.Read(r.reader, binary.LittleEndian, &s.Flags8)
	codeText, code := false, false
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case section.IDENTIFICATION_FLAGS_CODE:
			binary.Read(r.reader, binary.LittleEndian, &s.Code)
			code = true
		case section.IDENTIFICATION_FLAGS_CODETEXT:
			s.CodeText = r.ReadString()
			codeText = true
		case section.IDENTIFICATION_FLAGS_DEVICETYPE:
			binary.Read(r.reader, binary.LittleEndian, &s.Type)
		case section.IDENTIFICATION_FLAGS_FIRMWARE:
			binary.Read(r.reader, binary.LittleEndian, &s.Firmware)
		case section.IDENTIFICATION_FLAGS_HARDWARE:
			binary.Read(r.reader, binary.LittleEndian, &s.Hardware)
		case section.IDENTIFICATION_FLAGS_DEVICEHASH:
			binary.Read(r.reader, binary.LittleEndian, &s.Hash)
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

func (r *TelematicsReader) ReadAuthentication(s *section.Authentication) {
	binary.Read(r.reader, binary.LittleEndian, &s.Flags8)
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case section.AUTHENTICATION_FLAGS_IDENTIFIER:
			s.Identifier = r.ReadString()
		case section.AUTHENTICATION_FLAGS_SECRET:
			s.Secret = r.ReadBytes()
		default:
			panic("flag not supported")
		}
	}
}

func (r *TelematicsReader) ReadModule(s *section.Module) {
	binary.Read(r.reader, binary.LittleEndian, &s.Flags8)
	binary.Read(r.reader, binary.LittleEndian, &s.Id)
	var flags [8]uint8
	s.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
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

func (r *TelematicsReader) ReadModuleProperty(mp *section.ModuleProperty) {
	binary.Read(r.reader, binary.LittleEndian, &mp.Flags8)
	binary.Read(r.reader, binary.LittleEndian, &mp.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &mp.Id)
	binary.Read(r.reader, binary.LittleEndian, &mp.Type)
	var flags [8]byte
	mp.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case section.MODULE_PROPERTY_FLAGS_MIN:
			binary.Read(r.reader, binary.LittleEndian, &mp.Min)
		case section.MODULE_PROPERTY_FLAGS_MAX:
			binary.Read(r.reader, binary.LittleEndian, &mp.Max)
		case section.MODULE_PROPERTY_FLAGS_LIST:
			mp.List = r.ReadNameValues(mp.Type)
		case section.MODULE_PROPERTY_FLAGS_ACCESS:
			binary.Read(r.reader, binary.LittleEndian, &mp.Access)
		case section.MODULE_PROPERTY_FLAGS_NAME:
			mp.Name = r.ReadString()
		case section.MODULE_PROPERTY_FLAGS_DESCRIPTION:
			mp.Desc = r.ReadString()
		}
	}
}

func (r *TelematicsReader) ReadModulePropertyValue(s *section.ModulePropertyValue) (err error) {
	s.Values = make(map[byte]interface{})
	binary.Read(r.reader, binary.LittleEndian, &s.ModuleId)
	var c, i byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i = 0; i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)

		var p section.ModuleProperty
		if !r.Configuration.GetProperty(s.ModuleId, id, &p) {
			err = errors.New(fmt.Sprintf("property not found: %v %v", s.ModuleId, id))
			return
		}

		s.Values[id] = r.readData(p.Type)
	}
	return
}

func (r *TelematicsReader) ReadModulePropertyDisable(s *section.ModulePropertyDisable) {
	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)
		v := r.readData(value.Byte)
		s.DisabledProperties[id] = v.(byte)
	}
}

func (r *TelematicsReader) ReadCommand(c *section.Command) {
	binary.Read(r.reader, binary.LittleEndian, &c.Flags8)
	binary.Read(r.reader, binary.LittleEndian, &c.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &c.Id)
	var flags [8]byte
	c.Load(&flags)
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		switch flag {
		case section.COMMAND_FLAGS_NAME:
			c.Name = r.ReadString()
		case section.COMMAND_FLAGS_DESCRIPTION:
			c.Description = r.ReadString()
		default:
			panic("flag not supported")
		}
	}
}

func (r *TelematicsReader) ReadCommandArgument(ca *section.CommandArgument) {
	binary.Read(r.reader, binary.LittleEndian, &ca.Flags8)
	binary.Read(r.reader, binary.LittleEndian, &ca.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &ca.CommandId)
	binary.Read(r.reader, binary.LittleEndian, &ca.Id)
	binary.Read(r.reader, binary.LittleEndian, &ca.Type)

	var flags [8]byte
	ca.Load(&flags)
	for _, flag := range flags {
		switch flag {
		case section.COMMAND_ARGUMENT_FLAGS_MIN:
			ca.Min = r.readData(ca.Type)
		case section.COMMAND_ARGUMENT_FLAGS_MAX:
			ca.Max = r.readData(ca.Type)
		case section.COMMAND_ARGUMENT_FLAGS_LIST:
			ca.List = r.ReadNameValues(ca.Type)
		case section.COMMAND_ARGUMENT_FLAGS_REQUIRED:
			binary.Read(r.reader, binary.LittleEndian, &ca.Required)
		case section.COMMAND_ARGUMENT_FLAGS_NAME:
			ca.Name = r.ReadString()
		case section.COMMAND_ARGUMENT_FLAGS_DESCRIPTION:
			ca.Desc = r.ReadString()
		}
	}
}

func (r *TelematicsReader) ReadCommandExecute(ce *section.CommandExecute) {
	binary.Read(r.reader, binary.LittleEndian, &ce.ModuleId)
	binary.Read(r.reader, binary.LittleEndian, &ce.CommandId)

	var c byte
	binary.Read(r.reader, binary.LittleEndian, &c)
	var id byte
	for i := byte(0); i < c; i++ {
		binary.Read(r.reader, binary.LittleEndian, &id)

		var arg section.CommandArgument
		if !r.Configuration.GetArgument(ce.ModuleId, ce.CommandId, id, &arg) {
			panic("unable to read commandArgument")
		}

		ce.Arguments[id] = r.readData(arg.Type)
	}
}

func (r *TelematicsReader) ReadSupported(s *section.Supported) {
	bytes := r.ReadBytes()
	for _, byte := range bytes {
		sectionType := section.Type(byte)
		s.Set(sectionType.Flag(), true)
	}
}
