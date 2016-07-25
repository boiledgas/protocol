package telematics

import "protocol/utils"

type baseSection struct {
	flags byte
}

type Section interface {
	getFlags() []byte
	setFlag(flag byte, val bool)
}

func (s *baseSection) getFlags() []byte {
	return utils.GetFlags8(s.flags)
}

func (s *baseSection) hasFlag(flag byte) bool {
	return s.flags&flag > 0
}

func (s *baseSection) setFlag(flag byte, val bool) {
	if val {
		s.flags |= flag
	} else {
		s.flags ^= flag
	}
}
