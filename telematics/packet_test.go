package telematics

import (
	"bytes"
	"fmt"
	"testing"
	"time"
)

func Test_Packet_RequestAuthentication(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf, nil)
	w := NewWriter(&buf, nil)

	req := NewRequest()
	req.SetTimestamp(time.Now())
	auth, _ := req.Authentication()
	auth.SetIdentifier("777")
	auth.SetSecret([]byte{0x00, 0x01, 0x02, 0x03})
	id, _ := req.Identification()
	id.SetCodeText("777")
	id.SetHash(77)
	sup, _ := req.Supported()
	sup.Support(SECTION_AUTHENTICATION, true)
	sup.Support(SECTION_IDENTIFICATION, true)
	sup.Support(SECTION_MODULE, true)

	w.WritePacket(req.(*requestStruct))
	res := r.ReadPacket().(*requestStruct)

	if res.GetSequence() != req.GetSequence() {
		t.Errorf("sequence wrong %v != %v", res.GetSequence(), req.GetSequence())
	}
	if res.GetTimestamp() != req.GetTimestamp() {
		t.Errorf("timestamp wrong %v != %v", res.GetSequence(), req.GetSequence())
	}
	if res_auth, ok := res.Authentication(); ok {
		if res_auth_id, ok := res_auth.GetIdentifier(); ok {
			if auth_id, _ := auth.GetIdentifier(); auth_id != res_auth_id {
				panic(fmt.Sprintf("auth_id %v != %v", auth_id, res_auth_id))
			}
		} else {
			panic("no auth identifier")
		}
	}
	if res_sup, ok := res.Supported(); ok {
		if !res_sup.IsSupported(SECTION_AUTHENTICATION) {
			panic("section authentication not supported")
		}
		if !res_sup.IsSupported(SECTION_IDENTIFICATION) {
			panic("section identification not supported")
		}
		if !res_sup.IsSupported(SECTION_MODULE) {
			panic("section module not supported")
		}
		if res_sup.IsSupported(SECTION_COMMAND) {
			panic("section command not supported")
		}
		if res_sup.IsSupported(SECTION_COMMAND_ARGUMENT) {
			panic("section command argument not supported")
		}
		if res_sup.IsSupported(SECTION_COMMAND_EXECUTE) {
			panic("section comand execute not supported")
		}
	}
}

func Test_Packet_RequestConfiguration(t *testing.T) {
	buf := bytes.Buffer{}
	r := NewReader(&buf, nil)
	w := NewWriter(&buf, nil)

	m1_id := byte(1)
	m2_id := byte(2)
	req := NewRequest()

	conf := NewConfiguration()

	m1, _ := conf.GetModule(m1_id)
	//if conf.SetModule(m1) {
	//	t.Fail()
	//}
	m1.SetName("test 1")
	m1.SetDesc("test description 1")

	m3 := NewModule(m2_id)
	//if conf.SetModule(m3) {
	//	t.Fail()
	//}
	m3.SetName("test 3")
	m3.SetDesc("test description 3")

	m2 := NewModule(m2_id)
	//if !conf.SetModule(m2) {
	//	t.Fail()
	//}
	m2.SetName("test 2")
	m2.SetDesc("test description 2")

	p11 := NewModuleProperty(1, m1)
	conf.SetProperty(p11)
	p11.SetName("property 1 1")
	p11.SetDesc("property description 1 1")
	p12 := NewModuleProperty(2, m2)
	conf.SetProperty(p12)
	p12.SetName("property 1 2")
	w.WritePacket(req.(*requestStruct))
	res := r.ReadPacket().(*requestStruct)

	if res.GetSequence() != req.GetSequence() {
		t.Errorf("sequence wrong %v != %v", res.GetSequence(), req.GetSequence())
	}
	if res.GetTimestamp() != req.GetTimestamp() {
		t.Errorf("timestamp wrong %v != %v", res.GetSequence(), req.GetSequence())
	}
	if m, ok := conf.GetModule(m1_id); ok {
		m_name, _ := m.GetName()
		m1_name, _ := m1.GetName()
		if m_name != m1_name {
			panic(fmt.Sprintf("module %v name %v != %v", m.GetId(), m_name, m1_name))
		}
		m_desc, _ := m.GetDesc()
		m1_desc, _ := m1.GetDesc()
		if m_desc != m1_desc {
			panic(fmt.Sprintf("module %v desc %v != %v", m.GetId(), m_desc, m1_desc))
		}
	} else {
		panic(fmt.Sprintf("module %v not exists", m1_id))
	}
}
