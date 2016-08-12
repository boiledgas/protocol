package section

import "protocol/utils"

// authentication flags
const (
	AUTHENTICATION_FLAGS_IDENTIFIER byte = 0x01
	AUTHENTICATION_FLAGS_SECRET          = 0x02
)

type Authentication struct {
	utils.Flags8
	Identifier string
	Secret     []byte
}
