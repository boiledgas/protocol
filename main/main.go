package main

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"math"
	"math/rand"
	"protocol/badtelematics"
	"protocol/utils"
	"reflect"
	"strings"
)

func parsing() {
	//packets := []string{
	//	"aa020074250d57013a0f2b3728343935293236392d31313131650064000dbb40",
	//	"aa020174250d570303010343505504373937570303020347534d0442475332030303034750530c75424c4f582d4e454f2d364d0430010131033132560b45787420766f6c74616765043001023004342e3356044c69506f043001032304435055741b696e7465726e616c2074656d70657261747572652073656e736f7204300104390341636314332d6178697320616363656c65726f6d65746572043801420502064265657065720a426565706572286d532904300201380843656c6c496e666f0d47534d2063656c6c20696e666f043003013708506f736974696f6e0c47505320706f736974696f6e04380106010507636f6e6669673012636f6e66206465736372697074696f6e203004380107010607636f6e6669673112636f6e66206465736372697074696f6e203104380108010707636f6e6669673212636f6e66206465736372697074696f6e20320703030109434d445f314152474e0b626c6120626c6120626c610703030210434d445f574954484f55545f415247530b4c4f4c204c4f4c204c4f4c081003010008056172676e30bb48",
	//	"aa020274250d5705010401046d3a00000204940e0000037b00040f5bfcc6ff93026400050201010000000000000000050301011faa3d342172b25f16f9ff404009bbae",
	//	"aa020376250d570501040104b22f000002045f100000037b00040fab00840168036400050201010000000000000000050301011fa93d342162b25f1629ff59f209bbcf"}
	//data := strings.Join(packets, "")
	//bs, _ := hex.DecodeString(data)
	//br := bytes.NewReader(bs)
	//bw := &bytes.Buffer{}
	//r := telematics.NewReader(br, nil)
	//w := telematics.NewWriter(bw, nil)
	//f := func(r *telematics.TelematicsReader) bool {
	//	p := r.ReadPacket()
	//	if req, ok := p.(*telematics.Request); ok {
	//		resp := telematics.Response{}
	//		resp.Sequence = req.Sequence
	//		utils.SetFlags8([]byte{telematics.RESPONSE_OK}, &resp.Flags)
	//		w.WritePacket(resp)
	//		fmt.Println("request", req)
	//		fmt.Println("response", hex.EncodeToString(bw.Bytes()))
	//		bw.Reset()
	//	} else if resp, ok := p.(*telematics.Response); ok {
	//		fmt.Println("response", resp)
	//	}
	//	return p != nil
	//}
	//for f(r) {
	//}
}

type Record struct {
	id  uint16
	val uint32
}

func main() {
	chs := make([]chan *Record, 100)
	stat := make(map[uint16]uint16)
	for i := 0; i < len(chs); i++ {
		chs[i] = make(chan *Record)
		go func(id uint16, ch chan *Record) {
			stat[id] = 0
			for i := 0; true; i++ {
				ch <- &Record{
					id:  id,
					val: i,
				}
			}
		}(i, chs[i])
	}

	go func() {
		cases := make([]reflect.SelectCase, len(chs))
		for i, ch := range chs {
			cases[i] = reflect.SelectCase{Dir: reflect.SelectRecv, Chan: reflect.ValueOf(ch)}
		}
		for {
			chosen, ch, ok := reflect.Select(cases)
			if !ok {
				chs[chosen] = nil
			}
			for {
				if rec, ok := <-ch; !ok {
					stat[rec.id] ++
					break
				}
			}
		}
	}()
}
