package section

import (
	"protocol/telematics/value"
	"protocol/utils"
	"fmt"
)

// module property flags
const (
	MODULE_PROPERTY_FLAGS_MIN         byte = 0x01
	MODULE_PROPERTY_FLAGS_MAX              = 0x02
	MODULE_PROPERTY_FLAGS_LIST             = 0x04
	MODULE_PROPERTY_FLAGS_ACCESS           = 0x08
	MODULE_PROPERTY_FLAGS_NAME             = 0x10
	MODULE_PROPERTY_FLAGS_DESCRIPTION      = 0x20
)

type PropertyAccess byte

// property access
const (
	PROPERTYACCESS_NOTSET   PropertyAccess = 0x00
	PROPERTYACCESS_READ     PropertyAccess = 0x01
	PROPERTYACCESS_WRITE    PropertyAccess = 0x02
	PROPERTYACCESS_CONFIG   PropertyAccess = 0x04
	PROPERTYACCESS_DISABLED PropertyAccess = 0x08
)

type ModuleProperty struct {
	utils.Flags8

	ModuleId byte
	Id       byte
	Type     value.DataType
	Min      interface{}
	Max      interface{}
	List     []value.NameValue
	Access   PropertyAccess
	Name     string
	Desc     string
}

func (m ModuleProperty) String() string {
	return fmt.Sprintf("{Id:%v; ModuleId:%v; Name:%v, Type:%v}", m.Id, m.ModuleId, m.Name, m.Type.String())
}
