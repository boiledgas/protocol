package telematics

import "github.com/boiledgas/protocol/utils"

const (
	FLAG_REQUEST  byte = 0x01
	FLAG_RESPONSE byte = 0x02
)

type Packet struct {
	utils.Flags8
	Request  Request
	Response Response
}
