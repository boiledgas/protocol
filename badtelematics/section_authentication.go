package telematics

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

func (s *Authentication) GetIdentifier() (id string, ok bool) {
	id = s.Identifier
	ok = s.Has(AUTHENTICATION_FLAGS_IDENTIFIER)
	return
}

func (s *Authentication) SetIdentifier(id string) {
	s.Identifier = id
	s.Set(AUTHENTICATION_FLAGS_IDENTIFIER, len(id) > 0)
}

func (s *Authentication) GetSecret() (secret []byte, ok bool) {
	secret = s.Secret
	ok = s.Has(AUTHENTICATION_FLAGS_SECRET)
	return
}

func (s *Authentication) SetSecret(secret []byte) {
	s.Secret = secret
	s.Set(AUTHENTICATION_FLAGS_SECRET, len(secret) > 0)
}
