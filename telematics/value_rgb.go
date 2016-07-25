package telematics

type RgbStruct struct {
	r byte
	g byte
	b byte
}

type RgbValue interface {
	GetR() byte
	GetG() byte
	GetB() byte
}

func (v *RgbStruct) GetR() byte {
	return v.r
}

func (v *RgbStruct) GetG() byte {
	return v.g
}

func (v *RgbStruct) GetB() byte {
	return v.b
}
