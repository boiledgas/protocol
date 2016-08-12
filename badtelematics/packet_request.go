package telematics

import (
	"container/list"
	"protocol/utils"
	"sync/atomic"
	"time"
)

var sequence byte

func Sequence() byte {
	atomic.CompareAndSwapInt32(&sequence, 255, 0)
	return atomic.AddInt32(&sequence, 1)
}

type Request struct {
	utils.Flags8

	Sequence   byte
	Timestamp  int32
	Id         Identification
	Auth       Authentication
	Sup        Supported
	Properties map[byte]modulePropertySection
	Values     map[byte]modulePropertyValueSection
}

func (r *Request) GetSequence() byte {
	return r.Sequence
}

func (r *Request) GetTimestamp() time.Time {
	return time.Unix(int64(r.Timestamp), 0)
}

func (r *Request) SetTimestamp(t time.Time) {
	r.Timestamp = int32(t.Unix())
}

func (r *Request) Sections() []SectionType {
	s := make([]SectionType, 0, 10)
	var ok bool
	var section_type SectionType
	for i, section := range r.sections {
		if section != nil {
			if section_type, ok = ToSectionType(byte(i)); ok {
				s = append(s, section_type)
			}
		}
	}

	return s
}

func (r *Request) Authentication() (a Authentication, ok bool) {
	var s_key byte
	if s_key, ok = SectionFlag(SECTION_AUTHENTICATION); !ok {
		return
	}

	ok = true
	if r.sections[s_key] == nil {
		r.sections[s_key] = NewAuthentication()
		ok = false
	}

	a = r.sections[s_key].(Authentication)
	return
}

func (r *Request) Identification() (id Identification, ok bool) {
	var s_key byte
	if s_key, ok = SectionFlag(SECTION_IDENTIFICATION); !ok {
		return
	}

	ok = true
	if r.sections[s_key] == nil {
		r.sections[s_key] = NewIdentification()
		ok = false
	}

	id = r.sections[s_key].(Identification)
	return
}

func (r *Request) Supported() (s Supported, ok bool) {
	s_key, _ := SectionFlag(SECTION_SUPPORTED)
	ok = true
	if r.sections[s_key] == nil {
		r.sections[s_key] = NewSupported()
		ok = false
	}

	s = r.sections[s_key].(Supported)
	return
}

func (r *Request) SetConf(c Conf) {
	for _, m := range c.Modules() {
		r.pushConf(m)
	}
	for _, p := range c.Properties() {
		r.pushConf(p)
	}
	for _, c := range c.Commands() {
		r.pushConf(c)
	}
	for _, ca := range c.Arguments() {
		r.pushConf(ca)
	}
}

func (r *Request) GetConf() (c Conf, res bool) {
	return
}

func (r *Request) PropertyValue(m Module) ModulePropertyValue {
	var s_key byte
	var ok bool
	if s_key, ok = SectionFlag(SECTION_MODULE_PROPERTY_VALUE); !ok {
		panic("section not founc")
	}
	vs := r.sections[s_key].([]ModulePropertyValue)
	if vs == nil {
		vs = make([]ModulePropertyValue, 0, 10)
		r.sections[s_key] = vs
	}

	v := NewModulePropertyValue(m)
	vs = append(vs, v)
	return v
}

func (r *Request) CommandExecute(c Command) CommandExecute {
	var s_key byte
	var ok bool
	if s_key, ok = SectionFlag(SECTION_COMMAND_EXECUTE); !ok {
		panic("section not founc")
	}
	ces := r.sections[s_key].([]CommandExecute)
	if ces == nil {
		ces = make([]CommandExecute, 0, 10)
		r.sections[s_key] = ces
	}

	ce := NewCommandExecute(c)
	ces = append(ces, ce)
	return ce
}

func (r *Request) getModules(create_not_exist bool) *list.List {
	sid, _ := SectionFlag(SECTION_MODULE)
	if r.sections[sid] == nil {
		if !create_not_exist {
			return nil
		}

		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *Request) getProperties() *list.List {
	sid, _ := SectionFlag(SECTION_MODULE_PROPERTY)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *Request) getCommands() *list.List {
	sid, _ := SectionFlag(SECTION_COMMAND)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *Request) getCommandArguments() *list.List {
	sid, _ := SectionFlag(SECTION_COMMAND_ARGUMENT)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *Request) pushConf(s interface{}) {
	var l *list.List
	switch s.(type) {
	case Module:
		l = r.getModules(true)
	case ModuleProperty:
		l = r.getProperties()
	case Command:
		l = r.getCommands()
	case CommandArgument:
		l = r.getCommandArguments()
	default:
		panic("not configuration section")
	}

	l.PushBack(s)
}

func (r *Request) resetConf() {
	sid, _ := SectionFlag(SECTION_MODULE)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = SectionFlag(SECTION_MODULE_PROPERTY)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = SectionFlag(SECTION_COMMAND)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = SectionFlag(SECTION_COMMAND_ARGUMENT)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
}

func (r *Request) section(d interface{}) {
	switch s := d.(type) {
	case ModulePropertyValue:
		{
			if i, ok := SectionFlag(SECTION_MODULE_PROPERTY_VALUE); ok {
				sections := r.sections[i].([]ModulePropertyValue)
				if sections == nil {
					sections = make([]ModulePropertyValue, 0, 5)
				}
				r.sections[i] = append(sections, s)
			} else {
				panic("section module property value")
			}
		}
	case ModulePropertyDisable:
		{
			if i, ok := SectionFlag(SECTION_MODULE_PROPERTY_DISABLED); ok {
				sections := r.sections[i].([]ModulePropertyDisable)
				if sections == nil {
					sections = make([]ModulePropertyDisable, 0, 5)
				}
				r.sections[i] = append(sections, s)
			} else {
				panic("section module property disable")
			}
		}
	case CommandExecute:
		{
			if i, ok := SectionFlag(SECTION_COMMAND_EXECUTE); ok {
				sections := r.sections[i].([]CommandExecute)
				if sections == nil {
					sections = make([]CommandExecute, 0, 5)
				}
				r.sections[i] = append(sections, s)
			} else {
				panic("section command execute")
			}
		}
	default:
		panic("section not defined")
	}
}
