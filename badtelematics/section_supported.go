package telematics

import (
	"errors"
	"protocol/utils"
)

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

type Supported struct {
	utils.Flags16
}

func (s *Supported) Support(sectionType SectionType, value bool) {
	s.Set(sectionType.flag(), value)
}

func (s *Supported) IsSupported(sectionType SectionType) bool {
	return s.Has(sectionType.flag())
}

const (
	FLAG_IDENTIFICATION           uint16 = 0x001
	FLAG_AUTHENTICATION           uint16 = 0x002
	FLAG_SUPPORTED                uint16 = 0x004
	FLAG_MODULE                   uint16 = 0x008
	FLAG_MODULE_PROPERTY          uint16 = 0x010
	FLAG_MODULE_PROPERTY_VALUE    uint16 = 0x020
	FLAG_MODULE_PROPERTY_DISABLED uint16 = 0x040
	FLAG_COMMAND                  uint16 = 0x080
	FLAG_COMMAND_ARGUMENT         uint16 = 0x100
	FLAG_COMMAND_EXECUTE          uint16 = 0x200
)

func (t SectionType) flag() uint16 {
	switch t {
	case SECTION_IDENTIFICATION:
		return FLAG_IDENTIFICATION
	case SECTION_AUTHENTICATION:
		return FLAG_AUTHENTICATION
	case SECTION_SUPPORTED:
		return FLAG_SUPPORTED
	case SECTION_MODULE:
		return FLAG_MODULE
	case SECTION_MODULE_PROPERTY:
		return FLAG_MODULE_PROPERTY
	case SECTION_MODULE_PROPERTY_VALUE:
		return FLAG_MODULE_PROPERTY_VALUE
	case SECTION_MODULE_PROPERTY_DISABLED:
		return FLAG_MODULE_PROPERTY_DISABLED
	case SECTION_COMMAND:
		return FLAG_COMMAND
	case SECTION_COMMAND_ARGUMENT:
		return FLAG_COMMAND_ARGUMENT
	case SECTION_COMMAND_EXECUTE:
		return FLAG_COMMAND_EXECUTE
	default:
		panic("section not supported")
	}
	return
}

func ToSectionType(flag uint16) SectionType {
	switch flag {
	case 0x01:
		t = SECTION_IDENTIFICATION
	case 0x02:
		t = SECTION_AUTHENTICATION
	case 0x04:
		t = SECTION_MODULE
	case 0x08:
		t = SECTION_MODULE_PROPERTY
	case 0x10:
		t = SECTION_MODULE_PROPERTY_VALUE
	case 0x20:
		t = SECTION_MODULE_PROPERTY_DISABLED
	case 0x40:
		t = SECTION_COMMAND
	case 0x80:
		t = SECTION_COMMAND_ARGUMENT
	case 0x100:
		t = SECTION_COMMAND_EXECUTE
	case 0x200:
		t = SECTION_SUPPORTED
	default:
		ok = false
	}

	return
}
