package telematics

type SectionType byte

// section types
const (
	SECTION_ENDOFPAYLOAD             SectionType = 0xbb
	SECTION_IDENTIFICATION           SectionType = 0x01
	SECTION_AUTHENTICATION           SectionType = 0x02
	SECTION_MODULE                   SectionType = 0x03
	SECTION_MODULE_PROPERTY          SectionType = 0x04
	SECTION_MODULE_PROPERTY_VALUE    SectionType = 0x05
	SECTION_MODULE_PROPERTY_DISABLED SectionType = 0x06
	SECTION_COMMAND                  SectionType = 0x07
	SECTION_COMMAND_ARGUMENT         SectionType = 0x08
	SECTION_COMMAND_EXECUTE          SectionType = 0x09
	SECTION_SUPPORTED                SectionType = 0x0A
)

type supportedSection struct {
	codes [10]bool
}

type Supported interface {
	Support(SectionType, bool)
	IsSupported(SectionType) bool
	Get() []byte
	Set([]byte)
}

func (s *supportedSection) Support(t SectionType, v bool) {
	if i, ok := sectionIndex(t); ok {
		s.codes[i] = v
	}
}

func (s *supportedSection) IsSupported(t SectionType) bool {
	if i, ok := sectionIndex(t); ok {
		return s.codes[i]
	}

	return false
}

func (s *supportedSection) Get() []byte {
	r := make([]byte, 0, 10)
	for k, v := range s.codes {
		if v {
			if t, ok := sectionType(byte(k)); ok {
				r = append(r, byte(t))
			} else {
				panic("section not defined")
			}
		}
	}
	return r
}

func (s *supportedSection) Set(v []byte) {
	for _, r := range v {
		if i, ok := sectionIndex(SectionType(r)); ok {
			s.codes[i] = true
		}
	}
}

func NewSupported() Supported {
	v := supportedSection{}
	return &v
}
