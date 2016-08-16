package section

type Type byte

// section types
const (
	SECTION_UNKNOWN                  Type = 0x00
	SECTION_ENDOFPAYLOAD             Type = 0xbb
	SECTION_IDENTIFICATION           Type = 0x01
	SECTION_AUTHENTICATION           Type = 0x02
	SECTION_MODULE                   Type = 0x03
	SECTION_MODULE_PROPERTY          Type = 0x04
	SECTION_MODULE_PROPERTY_VALUE    Type = 0x05
	SECTION_MODULE_PROPERTY_DISABLED Type = 0x06
	SECTION_COMMAND                  Type = 0x07
	SECTION_COMMAND_ARGUMENT         Type = 0x08
	SECTION_COMMAND_EXECUTE          Type = 0x09
	SECTION_SUPPORTED                Type = 0x0A
)
