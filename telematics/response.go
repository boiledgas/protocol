package telematics

type ResponseFlag byte

const (
	RESPONSE_OK            ResponseFlag = 0x00
	RESPONSE_AUTHORIZATION ResponseFlag = 0x01
	RESPONSE_DESCRIPTION   ResponseFlag = 0x02
	RESPONSE_ERROR         ResponseFlag = 0x80
)

type Response struct {
	Flags    ResponseFlag
	Sequence byte
	Crc      byte
}
