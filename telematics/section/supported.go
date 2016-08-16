package section

import (
	"protocol/utils"
)

type Supported struct {
	utils.Flags16
}

func (s *Supported) Support(sectionType Type, value bool) {
	s.Set(sectionType.Flag(), value)
}

func (s *Supported) IsSupported(sectionType Type) bool {
	return s.Has(sectionType.Flag())
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

func (t Type) Flag() uint16 {
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
	return 0
}

func ToSectionType(flag uint16) (t Type) {
	switch flag {
	case FLAG_IDENTIFICATION:
		t = SECTION_IDENTIFICATION
	case FLAG_AUTHENTICATION:
		t = SECTION_AUTHENTICATION
	case FLAG_MODULE:
		t = SECTION_MODULE
	case FLAG_MODULE_PROPERTY:
		t = SECTION_MODULE_PROPERTY
	case FLAG_MODULE_PROPERTY_VALUE:
		t = SECTION_MODULE_PROPERTY_VALUE
	case FLAG_MODULE_PROPERTY_DISABLED:
		t = SECTION_MODULE_PROPERTY_DISABLED
	case FLAG_COMMAND:
		t = SECTION_COMMAND
	case FLAG_COMMAND_ARGUMENT:
		t = SECTION_COMMAND_ARGUMENT
	case FLAG_COMMAND_EXECUTE:
		t = SECTION_COMMAND_EXECUTE
	case FLAG_SUPPORTED:
		t = SECTION_SUPPORTED
	default:
		t = SECTION_UNKNOWN
	}

	return
}
