package telematics

import (
	"bytes"
	"encoding/binary"
	"encoding/hex"
	"fmt"
	"log"
	"protocol/telematics/section"
	"testing"
)

func TestPacket1(t *testing.T) {
	var buf bytes.Buffer
	bytes, _ := hex.DecodeString("aa020000000000013a0f383638303432303230323439353536010001000dbb09")
	buf.Write(bytes)

	reader := NewReader(&buf)
	req := Request{}
	var pt byte
	if err := binary.Read(reader.reader, binary.BigEndian, &pt); err != nil {
		t.Error(err)
	}
	if pt != PACKET_TYPE_REQUEST {
		t.Error("not request")
	}
	if err := reader.ReadRequest(&req); err != nil {
		t.Error(err)
	}
	log.Println(req)
}

func TestPacket2(t *testing.T) {
	var buf bytes.Buffer
	bytes, _ := hex.DecodeString("aa020200000000030301034350551b41524d76372050726f636573736f72207265762035202876376c290303020347534d1f53494d3a20756e646566696e656420284d43433a20302c204d4e433a2030290303030347505309756e646566696e6564043001013105506f77657205506f776572043001023007426174746572790742617474657279043001032304435055740f4350552074656d706572617475726504300104390c416363656c65726174696f6e14332d6178697320616363656c65726f6d657465720430014205064265657065720a426565706572286d532904300201380843656c6c496e666f0d47534d2063656c6c20696e666f043003013708506f736974696f6e0c47505320506f736974696f6ebb80")
	buf.Write(bytes)

	reader := NewReader(&buf)
	var pt byte
	if err := binary.Read(reader.reader, binary.BigEndian, &pt); err != nil {
		t.Error(err)
	}
	if pt != PACKET_TYPE_REQUEST {
		t.Error("not request")
	}

	req := Request{}
	if err := reader.ReadRequest(&req); err != nil {
		t.Error(err)
	}
	if !req.HasConfiguration() {
		t.Error("no configuration found")
	} else {
		fmt.Println(req.Conf)
	}
}

func TestPacket3(t *testing.T) {
	var buf bytes.Buffer
	bytes, _ := hex.DecodeString("aa020000000000013a0f383638303432303230323439353536010001000dbb09")
	buf.Write(bytes)
	bytes, _ = hex.DecodeString("aa020200000000030301034350551b41524d76372050726f636573736f72207265762035202876376c290303020347534d1f53494d3a20756e646566696e656420284d43433a20302c204d4e433a2030290303030347505309756e646566696e6564043001013105506f77657205506f776572043001023007426174746572790742617474657279043001032304435055740f4350552074656d706572617475726504300104390c416363656c65726174696f6e14332d6178697320616363656c65726f6d657465720430014205064265657065720a426565706572286d532904300201380843656c6c496e666f0d47534d2063656c6c20696e666f043003013708506f736974696f6e0c47505320506f736974696f6ebb80")
	buf.Write(bytes)
	bytes, _ = hex.DecodeString("aa021f00000000050301011fd56834213c7171160000000000bb46")
	buf.Write(bytes)

	var hash byte
	reader := NewReader(&buf)
	for {
		var pt byte
		if err := binary.Read(reader.reader, binary.BigEndian, &pt); err != nil {
			break
		}
		if pt != PACKET_TYPE_REQUEST {
			t.Error("not request")
		}

		req := Request{}
		if err := reader.ReadRequest(&req); err != nil {
			t.Error(err)
			break
		}
		if req.Has(section.FLAG_IDENTIFICATION) && req.Id.Has(section.IDENTIFICATION_FLAGS_DEVICEHASH) {
			if reader.Configuration == nil || req.Id.Hash != reader.Configuration.Hash {
				hash = req.Id.Hash
			}
		}
		if req.HasConfiguration() {
			conf := &req.Conf
			conf.Hash = hash
			reader.Configuration = conf

		}

		log.Println(req.String())
	}
}

func TestResponse(t *testing.T) {
	var buf bytes.Buffer
	writer := NewWriter(&buf)
	sequence := Sequence()
	resp := Response{Sequence: sequence, Flags: RESPONSE_OK | RESPONSE_DESCRIPTION}
	writer.WriteResponse(&resp)
	log.Println(buf.Bytes())
}
