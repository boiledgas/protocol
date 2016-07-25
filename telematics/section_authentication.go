package telematics

// authentication flags
const (
	AUTHENTICATION_FLAGS_IDENTIFIER byte = 0x01
	AUTHENTICATION_FLAGS_SECRET          = 0x02
)

type authenticationSection struct {
	baseSection

	Identifier string
	Secret     []byte
}

type Authentication interface {
	Section

	GetIdentifier() (string, bool)
	SetIdentifier(id string)

	GetSecret() ([]byte, bool)
	SetSecret(secret []byte)
}

func NewAuthentication() Authentication {
	res := authenticationSection{}
	return &res
}

func (s *authenticationSection) GetIdentifier() (id string, ok bool) {
	id = s.Identifier
	ok = s.hasFlag(AUTHENTICATION_FLAGS_IDENTIFIER)
	return
}

func (s *authenticationSection) SetIdentifier(id string) {
	s.Identifier = id
	s.setFlag(AUTHENTICATION_FLAGS_IDENTIFIER, len(id) > 0)
}

func (s *authenticationSection) GetSecret() (secret []byte, ok bool) {
	secret = s.Secret
	ok = s.hasFlag(AUTHENTICATION_FLAGS_SECRET)
	return
}

func (s *authenticationSection) SetSecret(secret []byte) {
	s.Secret = secret
	s.setFlag(AUTHENTICATION_FLAGS_SECRET, len(secret) > 0)
}
