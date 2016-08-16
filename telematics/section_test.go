package telematics

import (
	"bytes"
	"protocol/telematics/section"
	"protocol/telematics/value"
	"testing"
)

func Test_Identification(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	codeText := "huizhu"
	hash := byte(235)
	deviceType := section.DEVICETYPE_APPLICATION
	s := section.Identification{
		Hash:       hash,
		CodeText:   codeText,
		Type: deviceType,
	}
	s.Set(section.IDENTIFICATION_FLAGS_CODETEXT, true)
	s.Set(section.IDENTIFICATION_FLAGS_DEVICETYPE, true)
	s.Set(section.IDENTIFICATION_FLAGS_DEVICEHASH, true)

	w.WriteIdentification(&s)
	res := section.Identification{}
	r.ReadIdentification(&res)

	if res.Has(section.IDENTIFICATION_FLAGS_DEVICEHASH) {
		if hash != res.Hash {
			t.Errorf("hash wrong: %s != %s", hash, res.Hash)
		}
	} else {
		t.Error("Hash not exists")
	}
	if res.Has(section.IDENTIFICATION_FLAGS_CODETEXT) {
		if codeText != res.CodeText {
			t.Errorf("codeText wrong: %s != %s", codeText, res.CodeText)
		}
	} else {
		t.Error("Code text not exists")
	}
	if res.Has(section.IDENTIFICATION_FLAGS_DEVICETYPE) {
		if deviceType != res.Type {
			t.Errorf("deviceType wrong: %s != %s", deviceType, res.Type)
		}
	} else {
		t.Error("deviceType not exists")
	}
}

func Test_Authentication(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	s := section.Authentication{Identifier: "myId", Secret: []byte{0x01, 0x02, 0x03, 0x04}}
	s.Set(section.AUTHENTICATION_FLAGS_IDENTIFIER, true)
	s.Set(section.AUTHENTICATION_FLAGS_SECRET, true)

	w.WriteAuthentication(&s)
	res := section.Authentication{}
	r.ReadAuthentication(&res)

	if res.Has(section.AUTHENTICATION_FLAGS_IDENTIFIER) {
		if s.Identifier != res.Identifier {
			t.Errorf("Identitfier wrong: %s != %s", s.Identifier, res.Identifier)
		}
	} else {
		t.Error("Identitfier not exists")
	}
	if res.Has(section.AUTHENTICATION_FLAGS_SECRET) {
		if !bytes.Equal(s.Secret, res.Secret) {
			t.Errorf("Secret wrong: %v != %v", s.Secret, res.Secret)
		}
	} else {
		t.Error("Secret not exists")
	}
}

func Test_Module(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	s := section.Module{Id: byte(1), Name: "moduleName", Description: "moduleDesc"}
	s.Set(section.MODULE_FLAGS_NAME, true)
	s.Set(section.MODULE_FLAGS_DESCRIPTION, true)

	w.WriteModule(&s)
	res := section.Module{}
	r.ReadModule(&res)

	if s.Id != res.Id {
		t.Errorf("id wrong: %v != %v", s.Id, res.Id)
	}
	if res.Has(section.MODULE_FLAGS_NAME) {
		if s.Name != res.Name {
			t.Errorf("name wrong: %v != %v (%v)", s.Name, res.Name, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if res.Has(section.MODULE_FLAGS_DESCRIPTION) {
		if s.Description != res.Description {
			t.Errorf("desc wrong: %v != %v (%v)", s.Description, res.Description, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
}

func Test_ModuleProperty(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	m := section.Module{Id: byte(10), Name: "module1"}
	s := section.ModuleProperty{Id: byte(1), ModuleId: m.Id, Type: value.GPS, Name: "gps1", Desc: "module property Desc", Access: section.PROPERTYACCESS_READ | section.PROPERTYACCESS_WRITE}
	s.Set(section.MODULE_PROPERTY_FLAGS_NAME, true)
	s.Set(section.MODULE_PROPERTY_FLAGS_DESCRIPTION, true)
	s.Set(section.MODULE_PROPERTY_FLAGS_ACCESS, true)

	w.WriteModuleProperty(&s)
	res := section.ModuleProperty{}
	r.ReadModuleProperty(&res)

	if s.ModuleId != res.ModuleId {
		t.Errorf("moduleId wrong: %v != %v", s.ModuleId, res.ModuleId)
	}
	if s.Id != res.Id {
		t.Errorf("id wrong: %v != %v", s.Id, res.Id)
	}
	if s.Type != res.Type {
		t.Errorf("moduleType wrong: %v != %v", s.Type, res.Type)
	}

	if res.Has(section.MODULE_PROPERTY_FLAGS_NAME) {
		if s.Name != res.Name {
			t.Errorf("name wrong: %v != %v (%v)", s.Name, res.Name, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if res.Has(section.MODULE_PROPERTY_FLAGS_DESCRIPTION) {
		if s.Desc != res.Desc {
			t.Errorf("desc wrong: %v != %v (%v)", s.Desc, res.Desc, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
	if res.Has(section.MODULE_PROPERTY_FLAGS_ACCESS) {
		if s.Access != res.Access {
			t.Errorf("acc wrong: %v != %v (%v)", s.Access, res.Access, buf.Bytes())
		}
	} else {
		t.Errorf("acc not exists (%v)", buf.Bytes())
	}
}

func Test_ModulePropertyValue(t *testing.T) {
	m1 := section.Module{Id: 1, Name: "moduleName1"}
	p1 := section.ModuleProperty{Id: 1, ModuleId: m1.Id, Type: value.GPS, Name: "propertyName1", Desc: "propertyName1"}
	p1.Set(section.MODULE_PROPERTY_FLAGS_NAME, true)
	p1.Set(section.MODULE_PROPERTY_FLAGS_DESCRIPTION, true)

	p2 := section.ModuleProperty{Id: 2, ModuleId: m1.Id, Type: value.Int24, Name: "propertyName2"}
	p2.Set(section.MODULE_PROPERTY_FLAGS_NAME, true)

	conf := Configuration{Modules: []section.Module{m1}, Properties: []section.ModuleProperty{p1, p2}}

	buf := bytes.Buffer{}
	r := NewReader(&buf)
	r.Configuration = &conf
	w := TelematicsWriter{Writer: &buf}
	w.Configuration = &conf

	gps := value.Gps{Latitude: 35.55, Longitude: 55.55, Sat: byte(12)}
	gps.Set(value.GPS_FLAG_LATLNG, true)
	gps.Set(value.GPS_FLAG_SATELLITES, true)
	s := section.ModulePropertyValue{ModuleId: m1.Id, Values: map[byte]interface{}{p1.Id: gps, p2.Id: int32(12)}}

	w.WriteModulePropertyValue(&s)
	res := section.ModulePropertyValue{}
	r.ReadModulePropertyValue(&res)

	if m1.Id != res.ModuleId {
		t.Errorf("moduleId wrong: %v != %v", m1.Id, res.ModuleId)
	}

	if val, ok := res.Values[p1.Id]; ok {
		gps1 := val.(value.Gps)
		if !gps1.Has(value.GPS_FLAG_LATLNG) {
			t.Errorf("property %v not exists lat, lng (%v)", p1.Id, buf.Bytes())
		} else {
			if gps1.Latitude != gps.Latitude || gps1.Longitude != gps.Longitude {
				t.Errorf("latlng wrong %v != %v, %v != %v (%v)", gps.Latitude, gps1.Latitude, gps.Longitude, gps1.Longitude, buf.Bytes())
			}
		}
		if !gps1.Has(value.GPS_FLAG_SATELLITES) {
			t.Errorf("property %v not exists sattelites (%v)", p1.Id, buf.Bytes())
		} else {
			if gps1.Sat != gps.Sat {
				t.Errorf("sattelites wrong %v != %v (%v)", gps.Sat, gps1.Sat, buf.Bytes())
			}
		}
		if gps1.Has(value.GPS_FLAG_COURSE) {
			t.Errorf("property %v exists course (%v)", p1.Id, buf.Bytes())
		}
		if gps1.Has(value.GPS_FLAG_SPEED) {
			t.Errorf("property %v exists speed (%v)", p1.Id, buf.Bytes())
		}
		if gps1.Has(value.GPS_FLAG_ALTITUDE) {
			t.Errorf("property %v exists altitude (%v)", p1.Id, buf.Bytes())
		}
	} else {
		t.Errorf("property %v not set (%v)", p1.Id, buf.Bytes())
	}

	if val, ok := res.Values[p2.Id]; ok {
		if val != res.Values[p2.Id] {
			t.Errorf("p2 wrong %v != %v (%v)", val, res.Values[p2.Id], buf.Bytes())
		}
	} else {
		t.Errorf("property %v not set (%v)", p2.Id, buf.Bytes())
	}
}

func Test_ModulePropertyDisabled(t *testing.T) {
}

func Test_Command(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	m := section.Module{Id: byte(11)}
	s := section.Command{Id: byte(1), ModuleId: m.Id, Name: "moduleName", Description: "moduleDesc"}
	s.Set(section.COMMAND_FLAGS_NAME, true)
	s.Set(section.COMMAND_FLAGS_DESCRIPTION, true)

	w.WriteCommand(&s)
	res := section.Command{}
	r.ReadCommand(&res)

	if s.Id != res.Id {
		t.Errorf("id wrong: %v != %v", s.Id, res.Id)
	}
	if s.ModuleId != res.ModuleId {
		t.Errorf("moduleId wrong: %v != %v", s.ModuleId, res.ModuleId)
	}
	if s.Has(section.COMMAND_FLAGS_NAME) {
		if s.Name != res.Name {
			t.Errorf("name wrong: %v != %v (%v)", res.Name, s.Name, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if s.Has(section.COMMAND_FLAGS_DESCRIPTION) {
		if res.Description != s.Description {
			t.Errorf("desc wrong: %v != %v (%v)", res.Description, s.Description, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
}

func Test_CommandArgument(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	m := section.Module{Id: byte(10)}
	c := section.Command{Id: byte(1), ModuleId: m.Id}
	s := section.CommandArgument{Id: byte(19), CommandId: c.Id, Name: "arg1", Desc: "arg1desc", Type: value.Int24, Required: byte(1)}
	s.Set(section.COMMAND_ARGUMENT_FLAGS_NAME, true)
	s.Set(section.COMMAND_ARGUMENT_FLAGS_DESCRIPTION, true)
	s.Set(section.COMMAND_ARGUMENT_FLAGS_REQUIRED, true)

	w.WriteCommandArgument(&s)
	res := section.CommandArgument{}
	r.ReadCommandArgument(&res)

	if s.ModuleId != res.ModuleId {
		t.Errorf("module_id wrong: %v != %v", s.ModuleId, res.ModuleId)
	}
	if s.CommandId != res.CommandId {
		t.Errorf("command_id wrong: %v != %v", s.CommandId, res.CommandId)
	}
	if s.Id != res.Id {
		t.Errorf("id wrong: %v != %v", s.Id, res.Id)
	}
	if s.Type != res.Type {
		t.Errorf("dataType wrong: %v != %v", s.Type, res.Type)
	}

	if res.Has(section.COMMAND_ARGUMENT_FLAGS_NAME) {
		if s.Name != res.Name {
			t.Errorf("name wrong: %v != %v (%v)", s.Name, res.Name, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if res.Has(section.COMMAND_ARGUMENT_FLAGS_DESCRIPTION) {
		if s.Desc != res.Desc {
			t.Errorf("desc wrong: %v != %v (%v)", s.Desc, res.Desc, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
	if res.Has(section.COMMAND_ARGUMENT_FLAGS_REQUIRED) {
		if s.Required != res.Required {
			t.Errorf("req wrong: %v != %v (%v)", s.Required, res.Required, buf.Bytes())
		}
	} else {
		t.Errorf("req not exists (%v)", buf.Bytes())
	}
}

func Test_CommandExecute(t *testing.T) {
	m := section.Module{Id: byte(70)}
	c := section.Command{Id: byte(7), ModuleId: m.Id}
	ca1 := section.CommandArgument{Id: byte(77), CommandId: c.Id, ModuleId: c.ModuleId, Type: value.Byte}
	ca2 := section.CommandArgument{Id: byte(78), CommandId: c.Id, ModuleId: c.ModuleId, Type: value.GPS}

	conf := Configuration{
		Modules:   []section.Module{m},
		Commands:  []section.Command{c},
		Arguments: []section.CommandArgument{ca1, ca2},
	}

	s := section.CommandExecute{CommandId: c.Id, ModuleId: c.ModuleId, Arguments: make(map[byte]interface{})}
	s.Arguments[ca1.Id] = byte(12)
	arg2_value := value.Gps{Latitude: 35.77, Longitude: 77.77}
	arg2_value.Set(value.GPS_FLAG_LATLNG, true)
	s.Arguments[ca2.Id] = arg2_value

	buf := bytes.Buffer{}
	r := NewReader(&buf)
	r.Configuration = &conf
	w := TelematicsWriter{Writer: &buf}
	w.Configuration = &conf

	w.WriteCommandExecute(&s)
	res := section.CommandExecute{Arguments: make(map[byte]interface{})}
	r.ReadCommandExecute(&res)

	if s.ModuleId != res.ModuleId {
		t.Errorf("module_id wrong: %v != %v", s.ModuleId, res.ModuleId)
	}
	if s.CommandId != res.CommandId {
		t.Errorf("command_id wrong: %v != %v", s.CommandId, res.CommandId)
	}
	if arg, _ := res.Arguments[ca1.Id]; arg != s.Arguments[ca1.Id] {
		t.Errorf("arg1 wrong: %v != %v (%v)", s.Arguments[ca1.Id], res.Arguments[ca1.Id], buf.Bytes())
	}
	if s.Arguments[ca2.Id] != res.Arguments[ca2.Id] {
		t.Errorf("arg2 wrong: %v != %v (%v)", s.Arguments[ca2.Id], res.Arguments[ca2.Id], buf.Bytes())
	}
}

func Test_Supported(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf)
	w := TelematicsWriter{Writer: &buf}

	s := section.Supported{}
	s.Support(section.SECTION_IDENTIFICATION, true)
	s.Support(section.SECTION_MODULE, true)
	s.Support(section.SECTION_MODULE_PROPERTY, true)
	s.Support(section.SECTION_MODULE_PROPERTY_VALUE, true)

	w.writeSupported(&s)
	res := section.Supported{}
	r.ReadSupported(&res)

	var flags [16]uint16
	res.Load(&flags)
	var id, m, mp, mpv bool
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		sectionType := section.ToSectionType(flag)
		switch sectionType {
		case section.SECTION_IDENTIFICATION:
			id = true
		case section.SECTION_MODULE:
			m = true
		case section.SECTION_MODULE_PROPERTY:
			mp = true
		case section.SECTION_MODULE_PROPERTY_VALUE:
			mpv = true
		default:
			t.Errorf("Cant support %v", sectionType)
		}
	}

	if !id {
		t.Error("section SECTION_IDENTIFICATION not supported")
	}
	if !m {
		t.Error("section SECTION_MODULE not supported")
	}
	if !mp {
		t.Error("section SECTION_MODULE_PROPERTY not supported")
	}
	if !mpv {
		t.Error("section SECTION_MODULE_PROPERTY_VALUE not supported")
	}
}
