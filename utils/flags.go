package utils

func GetFlags8(val uint8) []uint8 {
	flags := make([]uint8, 0, 8)
	var flag uint8
	for i := uint8(0); i < 8; i++ {
		flag = uint8(1 << i)
		if val&flag > 0 {
			flags = append(flags, flag)
		}
	}

	return flags
}

func GetFlags16(val uint16) []uint16 {
	flags := make([]uint16, 0, 16)
	var flag uint16
	for i := uint16(0); i < 16; i++ {
		flag = uint16(1 << i)
		if val&flag > 0 {
			flags = append(flags, flag)
		}
	}

	return flags
}

func SetFlags8(flags []uint8, val *uint8) {
	*val = 0
	for _, flag := range flags {
		*val = *val | flag
	}
}

func SetFlag8(flag uint8, val *uint8) {
	*val = *val | flag
}

func SetFlags16(flags []uint16, val *uint16) {
	*val = 0
	for _, flag := range flags {
		*val = *val | flag
	}
}
