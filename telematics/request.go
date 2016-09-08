package telematics

import (
	"github.com/boiledgas/protocol/telematics/section"
	"github.com/boiledgas/protocol/utils"
	"sync/atomic"
	"bytes"
)

var sequence int32

func Sequence() byte {
	atomic.CompareAndSwapInt32(&sequence, 255, 0)
	return byte(atomic.AddInt32(&sequence, 1))
}

type Request struct {
	utils.Flags16

	Sequence  byte
	Timestamp int32
	Id        section.Identification
	Auth      section.Authentication
	Sup       section.Supported
	Conf      Configuration
	Values    []section.ModulePropertyValue
	Executes  []section.CommandExecute
	Disabled  []section.ModulePropertyDisable
}

func (r *Request) HasConfiguration() bool {
	return r.Has(section.FLAG_MODULE) ||
		r.Has(section.FLAG_MODULE_PROPERTY) ||
		r.Has(section.FLAG_COMMAND) ||
		r.Has(section.FLAG_COMMAND_ARGUMENT)
}

func (r Request) String() string {
	var flags [16]uint16
	r.Load(&flags)
	var buf bytes.Buffer
	for _, flag := range flags {
		if flag == 0 {
			continue
		}
		sectionType := section.ToSectionType(flag)
		switch sectionType {
		case section.SECTION_IDENTIFICATION:
			buf.WriteString(r.Id.String())
		case section.SECTION_AUTHENTICATION:
			buf.WriteString(r.Auth.String())
		case section.SECTION_COMMAND_EXECUTE:
			for _, s := range r.Executes {
				buf.WriteString(s.String())
			}
		case section.SECTION_MODULE_PROPERTY_DISABLED:
			for _, s := range r.Disabled {
				buf.WriteString(s.String())
			}
		case section.SECTION_MODULE_PROPERTY_VALUE:
			for _, s := range r.Values {
				buf.WriteString(s.String())
			}
		}
	}
	if r.HasConfiguration() {
		buf.WriteString(r.Conf.String())
	}
	return buf.String()
}