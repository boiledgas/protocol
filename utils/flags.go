package utils

type Flags8 uint8
type Flags16 uint16

func (f *Flags8) Has(flag uint8) bool {
	return uint8(*f)&flag > 0
}

func (f *Flags8) Set(flag uint8, value bool) {
	if value {
		*f |= Flags8(flag)
	} else {
		*f &= Flags8(^flag)
	}
}

func (f *Flags8) Load(flags *[8]byte) {
	var flag uint8
	for i := uint8(0); i < 8; i++ {
		flag = uint8(1 << i)
		if uint8(*f)&flag > 0 {
			flags[i] = flag
		}
	}
}

func (f *Flags16) Has(flag uint16) bool {
	return uint16(*f)&flag > 0
}

func (f *Flags16) Set(flag uint16, value bool) {
	if value {
		*f |= Flags16(flag)
	} else {
		*f &= Flags16(^flag)
	}
}

func (f *Flags16) Load(flags *[16]uint16) {
	var flag uint16
	for i := uint16(0); i < 16; i++ {
		flag = uint16(1 << i)
		if uint16(*f)&flag > 0 {
			flags[i] = flag
		}
	}
}
