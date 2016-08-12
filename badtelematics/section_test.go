package telematics

import (
	"bytes"
	"testing"
)

func Test_Identification(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	r.Configure(Configuration{})
	codeText := "huizhu"
	hash := byte(235)
	deviceType := DEVICETYPE_APPLICATION
	s := Identification{}
	s.SetHash(hash)
	s.SetCodeText(codeText)
	s.SetDeviceType(deviceType)

	w.WriteIdentification(s)
	res := Identification{}
	r.ReadIdentification(&res)
	if val, ok := res.GetHash(); ok {
		if hash != val {
			t.Errorf("hash wrong: %s != %s", hash, val)
		}
	} else {
		t.Error("Hash not exists")
	}
	if val, ok := res.GetCodeText(); ok {
		if codeText != val {
			t.Errorf("codeText wrong: %s != %s", codeText, val)
		}
	} else {
		t.Error("Code text not exists")
	}
	if val, ok := res.GetDeviceType(); ok {
		if deviceType != val {
			t.Errorf("deviceType wrong: %s != %s", deviceType, val)
		}
	} else {
		t.Error("deviceType not exists")
	}
}

func Test_Authentication(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	id := "myId"
	secret := []byte{0x01, 0x02, 0x03, 0x04}
	s := Authentication{}
	s.SetIdentifier(id)
	s.SetSecret(secret)

	w.writeAuthentication(s)
	res := Authentication{}
	r.ReadAuthentication(&res)

	if val, ok := res.GetIdentifier(); ok {
		if id != val {
			t.Errorf("Identitfier wrong: %s != %s", id, val)
		}
	} else {
		t.Error("Identitfier not exists")
	}
	if val, ok := res.GetSecret(); ok {
		if !bytes.Equal(secret, val) {
			t.Errorf("Secret wrong: %v != %v", secret, val)
		}
	} else {
		t.Error("Secret not exists")
	}
}

func Test_Module(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	id := byte(1)
	name := "moduleName"
	desc := "moduleDesc"
	s := NewModule(id)
	s.SetName(name)
	s.SetDesc(desc)

	w.writeModule(s.(*Module))
	res := &Module{}
	r.readModule(res)

	if id != res.GetId() {
		t.Errorf("id wrong: %v != %v", id, res.GetId())
	}
	if val, ok := res.GetName(); ok {
		if name != val {
			t.Errorf("name wrong: %v != %v (%v)", name, val, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetDesc(); ok {
		if desc != val {
			t.Errorf("desc wrong: %v != %v (%v)", desc, val, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
}

func Test_ModuleProperty(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	moduleId := byte(10)
	m := NewModule(moduleId)
	m.SetName("module1")
	id := byte(1)
	propType := GPS
	name := "gps1"
	desc := "module property Desc"
	acc := PROPERTYACCESS_READ | PROPERTYACCESS_WRITE

	s := NewModuleProperty(id, m)
	s.SetType(propType)
	s.SetName(name)
	s.SetDesc(desc)
	s.SetAccess(acc)

	w.writeModuleProperty(s.(*modulePropertySection))
	res := &modulePropertySection{}
	r.readModuleProperty(res)

	if moduleId != res.GetModuleId() {
		t.Errorf("moduleId wrong: %v != %v", moduleId, res.GetModuleId())
	}
	if id != res.GetId() {
		t.Errorf("id wrong: %v != %v", id, res.GetId())
	}
	if propType != res.GetType() {
		t.Errorf("moduleType wrong: %v != %v", propType, res.GetType())
	}

	if val, ok := res.GetName(); ok {
		if name != val {
			t.Errorf("name wrong: %v != %v (%v)", name, val, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetDesc(); ok {
		if desc != val {
			t.Errorf("desc wrong: %v != %v (%v)", desc, val, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetAccess(); ok {
		if acc != val {
			t.Errorf("acc wrong: %v != %v (%v)", acc, val, buf.Bytes())
		}
	} else {
		t.Errorf("acc not exists (%v)", buf.Bytes())
	}
}

func Test_ModulePropertyValue(t *testing.T) {
	m1 := NewModule(1)
	m1.SetName("moduleName1")

	p1 := NewModuleProperty(1, m1)
	p1.SetType(GPS)
	p1.SetName("propertyName1")
	p1.SetDesc("propertyName1")

	p2 := NewModuleProperty(2, m1)
	p2.SetType(Int24)
	p2.SetName("propertyName2")

	conf := NewConfiguration()
	conf.SetModule(m1)
	conf.SetProperty(p1)
	conf.SetProperty(p2)

	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	lat := 35.55
	lng := 55.55
	sat := byte(12)
	gps := GpsStruct{}
	gps.SetLatLng(lat, lng)
	gps.SetSat(sat)

	s := NewModulePropertyValue(p1)
	s.SetValue(p1.GetId(), gps)
	p2_val := int32(12)
	s.SetValue(p2.GetId(), p2_val)

	w.writeModulePropertyValue(s.(*modulePropertyValueSection))
	res := &modulePropertyValueSection{}
	r.readModulePropertyValue(res)

	if m1.GetId() != res.GetModuleId() {
		t.Errorf("moduleId wrong: %v != %v", m1.GetId(), res.GetModuleId())
	}

	if val, ok := res.GetValue(p1.GetId()); ok {
		gps1 := val.(GpsStruct)
		if lat1, lng1, ok := gps1.GetLatLng(); !ok {
			t.Errorf("property %v not exists lat, lng (%v)", p1.GetId(), buf.Bytes())
		} else {
			if lat1 != lat || lng1 != lng {
				t.Errorf("latlng wrong %v != %v, %v != %v (%v)", lat, lat1, lng, lng1, buf.Bytes())
			}
		}
		if sat1, ok := gps1.GetSat(); !ok {
			t.Errorf("property %v not exists sattelites (%v)", p1.GetId(), buf.Bytes())
		} else {
			if sat1 != sat {
				t.Errorf("sattelites wrong %v != %v (%v)", sat, sat1, buf.Bytes())
			}
		}
		if _, ok := gps1.GetCourse(); ok {
			t.Errorf("property %v exists course (%v)", p1.GetId(), buf.Bytes())
		}
		if _, ok := gps1.GetSpeed(); ok {
			t.Errorf("property %v exists speed (%v)", p1.GetId(), buf.Bytes())
		}
		if _, ok := gps1.GetAltitude(); ok {
			t.Errorf("property %v exists altitude (%v)", p1.GetId(), buf.Bytes())
		}
	} else {
		t.Errorf("property %v not set (%v)", p1.GetId(), buf.Bytes())
	}

	if val, ok := res.GetValue(p2.GetId()); ok {
		if val != p2_val {
			t.Errorf("p2 wrong %v != %v (%v)", val, p2_val, buf.Bytes())
		}
	} else {
		t.Errorf("property %v not set (%v)", p2.GetId(), buf.Bytes())
	}
}

func Test_ModulePropertyDisabled(t *testing.T) {
}

func Test_Command(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	id := byte(1)
	moduleId := byte(11)
	name := "moduleName"
	desc := "moduleDesc"
	m := NewModule(moduleId)
	s := NewCommand(id, m)
	s.SetName(name)
	s.SetDesc(desc)

	w.writeCommand(s.(*commandSection))
	res := &commandSection{}
	r.readCommand(res)

	if id != res.GetId() {
		t.Errorf("id wrong: %v != %v", id, res.GetId())
	}
	if moduleId != res.GetModuleId() {
		t.Errorf("moduleId wrong: %v != %v", moduleId, res.GetModuleId())
	}
	if val, ok := res.GetName(); ok {
		if name != val {
			t.Errorf("name wrong: %v != %v (%v)", name, val, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetDesc(); ok {
		if desc != val {
			t.Errorf("desc wrong: %v != %v (%v)", desc, val, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
}

func Test_CommandArgument(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	module_id := byte(10)
	m := NewModule(module_id)
	command_id := byte(1)
	c := NewCommand(command_id, m)
	name := "arg1"
	desc := "arg1desc"
	id := byte(19)
	dataType := Int24
	req := byte(1)
	s := NewCommandArgument(id, c)
	s.SetName(name)
	s.SetDesc(desc)
	s.SetType(dataType)
	s.SetRequired(req)

	w.writeCommandArgument(s.(*commandArgumentStruct))
	res := &commandArgumentStruct{}
	r.readCommandArgument(res)

	if module_id != res.GetModuleId() {
		t.Errorf("module_id wrong: %v != %v", module_id, res.GetModuleId())
	}
	if command_id != res.GetCommandId() {
		t.Errorf("command_id wrong: %v != %v", command_id, res.GetCommandId())
	}
	if id != res.GetId() {
		t.Errorf("id wrong: %v != %v", id, res.GetId())
	}
	if dataType != res.GetType() {
		t.Errorf("dataType wrong: %v != %v", dataType, res.GetType())
	}

	if val, ok := res.GetName(); ok {
		if name != val {
			t.Errorf("name wrong: %v != %v (%v)", name, val, buf.Bytes())
		}
	} else {
		t.Errorf("name not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetDesc(); ok {
		if desc != val {
			t.Errorf("desc wrong: %v != %v (%v)", desc, val, buf.Bytes())
		}
	} else {
		t.Errorf("desc not exists (%v)", buf.Bytes())
	}
	if val, ok := res.GetRequired(); ok {
		if req != val {
			t.Errorf("req wrong: %v != %v (%v)", req, val, buf.Bytes())
		}
	} else {
		t.Errorf("req not exists (%v)", buf.Bytes())
	}
}

func Test_CommandExecute(t *testing.T) {
	module_id := byte(70)
	command_id := byte(7)

	m := NewModule(module_id)
	c := NewCommand(command_id, m)
	arg1_id := byte(77)
	ca1 := NewCommandArgument(arg1_id, c)
	ca1.SetType(Byte)
	arg2_id := byte(78)
	ca2 := NewCommandArgument(arg2_id, c)
	ca2.SetType(GPS)

	conf := NewConfiguration()
	conf.SetModule(m)
	conf.SetCommand(c)
	conf.SetArgument(ca1)
	conf.SetArgument(ca2)

	s := NewCommandExecute(c)
	arg1_value := byte(12)
	s.SetArgument(arg1_id, arg1_value)
	arg2_value := GpsStruct{}
	arg2_value.SetLatLng(35.77, 77.77)
	s.SetArgument(arg2_id, arg2_value)

	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	w.writeCommandExecute(s.(*commandExecuteStruct))
	res := &commandExecuteStruct{}
	r.readCommandExecute(res)

	if module_id != s.GetModuleId() {
		t.Errorf("module_id wrong: %v != %v", module_id, res.GetModuleId())
	}
	if command_id != res.GetCommandId() {
		t.Errorf("command_id wrong: %v != %v", command_id, res.GetCommandId())
	}
	if arg1_res, ok := res.GetArgument(arg1_id); ok {
		if arg1_res != arg1_value {
			t.Errorf("arg1 wrong: %v != %v (%v)", arg1_res, arg1_value, buf.Bytes())
		}
	} else {
		t.Errorf("arg1 %v not exists (%v)", arg1_id, buf.Bytes())
	}
	if arg2_res, ok := res.GetArgument(arg2_id); ok {
		if arg2_res != arg2_value {
			t.Errorf("arg2 wrong: %v != %v (%v)", arg2_res, arg2_value, buf.Bytes())
		}
	} else {
		t.Errorf("arg2 %v not exists (%v)", arg2_id, buf.Bytes())
	}
}

func Test_Supported(t *testing.T) {
	buf := bytes.Buffer{}
	r := TelematicsReader{Reader: &buf}
	w := TelematicsWriter{Writer: &buf}

	s := NewSupported()
	s.Support(SECTION_IDENTIFICATION, true)
	s.Support(SECTION_MODULE, true)
	s.Support(SECTION_MODULE_PROPERTY, true)
	s.Support(SECTION_MODULE_PROPERTY_VALUE, true)

	w.writeSupported(s.(*Supported))
	res := &Supported{}
	r.ReadSupported(res)

	if len(res.Get()) != 4 {
		t.Errorf("section count: %v", len(res.Get()))
	}

	if !s.IsSupported(SECTION_IDENTIFICATION) {
		t.Error("section SECTION_IDENTIFICATION not supported")
	}
	if !s.IsSupported(SECTION_MODULE) {
		t.Error("section SECTION_MODULE not supported")
	}
	if !s.IsSupported(SECTION_MODULE_PROPERTY) {
		t.Error("section SECTION_MODULE_PROPERTY not supported")
	}
	if !s.IsSupported(SECTION_MODULE_PROPERTY_VALUE) {
		t.Error("section SECTION_MODULE_PROPERTY_VALUE not supported")
	}
}
