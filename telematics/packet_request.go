package telematics

import (
	"container/list"
	"time"
)

var sequence byte

type requestStruct struct {
	readonly  bool
	sequence  byte
	timestamp int32
	id        Identification
	auth      Authentication
	sections  [10]interface{}
}

type Request interface {
	GetSequence() byte
	GetTimestamp() time.Time
	SetTimestamp(time.Time)

	Authentication() (Authentication, bool)
	Identification() (Identification, bool)
	Supported() (Supported, bool)

	SetConf(c Conf)
	GetConf() (c Conf, res bool)

	PropertyValue(Module) ModulePropertyValue
	CommandExecute(Command) CommandExecute
}

func NewRequest() Request {
	r := requestStruct{sequence: sequence}
	sequence++
	return &r
}

func (r *requestStruct) GetSequence() byte {
	return r.sequence
}

func (r *requestStruct) GetTimestamp() time.Time {
	return time.Unix(int64(r.timestamp), 0)
}

func (r *requestStruct) SetTimestamp(t time.Time) {
	r.timestamp = int32(t.Unix())
}

func (r *requestStruct) Sections() []SectionType {
	s := make([]SectionType, 0, 10)
	var ok bool
	var section_type SectionType
	for i, section := range r.sections {
		if section != nil {
			if section_type, ok = sectionType(byte(i)); ok {
				s = append(s, section_type)
			}
		}
	}

	return s
}

func (r *requestStruct) Authentication() (a Authentication, ok bool) {
	var s_key byte
	if s_key, ok = sectionIndex(SECTION_AUTHENTICATION); !ok {
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

func (r *requestStruct) Identification() (id Identification, ok bool) {
	var s_key byte
	if s_key, ok = sectionIndex(SECTION_IDENTIFICATION); !ok {
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

func (r *requestStruct) Supported() (s Supported, ok bool) {
	s_key, _ := sectionIndex(SECTION_SUPPORTED)
	ok = true
	if r.sections[s_key] == nil {
		r.sections[s_key] = NewSupported()
		ok = false
	}

	s = r.sections[s_key].(Supported)
	return
}

func (r *requestStruct) SetConf(c Conf) {
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

func (r *requestStruct) GetConf() (c Conf, res bool) {
	return
}

func (r *requestStruct) PropertyValue(m Module) ModulePropertyValue {
	var s_key byte
	var ok bool
	if s_key, ok = sectionIndex(SECTION_MODULE_PROPERTY_VALUE); !ok {
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

func (r *requestStruct) CommandExecute(c Command) CommandExecute {
	var s_key byte
	var ok bool
	if s_key, ok = sectionIndex(SECTION_COMMAND_EXECUTE); !ok {
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

func (r *requestStruct) getModules(create_not_exist bool) *list.List {
	sid, _ := sectionIndex(SECTION_MODULE)
	if r.sections[sid] == nil {
		if !create_not_exist {
			return nil
		}

		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *requestStruct) getProperties() *list.List {
	sid, _ := sectionIndex(SECTION_MODULE_PROPERTY)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *requestStruct) getCommands() *list.List {
	sid, _ := sectionIndex(SECTION_COMMAND)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *requestStruct) getCommandArguments() *list.List {
	sid, _ := sectionIndex(SECTION_COMMAND_ARGUMENT)
	if r.sections[sid] == nil {
		r.sections[sid] = list.New()
	}
	return r.sections[sid].(*list.List)
}

func (r *requestStruct) pushConf(s interface{}) {
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

func (r *requestStruct) resetConf() {
	sid, _ := sectionIndex(SECTION_MODULE)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = sectionIndex(SECTION_MODULE_PROPERTY)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = sectionIndex(SECTION_COMMAND)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
	sid, _ = sectionIndex(SECTION_COMMAND_ARGUMENT)
	if r.sections[sid] != nil {
		r.sections[sid] = nil
	}
}

func (r *requestStruct) section(d interface{}) {
	switch s := d.(type) {
	case ModulePropertyValue:
		{
			if i, ok := sectionIndex(SECTION_MODULE_PROPERTY_VALUE); ok {
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
			if i, ok := sectionIndex(SECTION_MODULE_PROPERTY_DISABLED); ok {
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
			if i, ok := sectionIndex(SECTION_COMMAND_EXECUTE); ok {
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

func sectionIndex(t SectionType) (i byte, ok bool) {
	ok = true
	switch t {
	case SECTION_IDENTIFICATION:
		i = 0
	case SECTION_AUTHENTICATION:
		i = 1
	case SECTION_SUPPORTED:
		i = 9
	case SECTION_MODULE:
		i = 2
	case SECTION_MODULE_PROPERTY:
		i = 3
	case SECTION_MODULE_PROPERTY_VALUE:
		i = 4
	case SECTION_MODULE_PROPERTY_DISABLED:
		i = 5
	case SECTION_COMMAND:
		i = 6
	case SECTION_COMMAND_ARGUMENT:
		i = 7
	case SECTION_COMMAND_EXECUTE:
		i = 8
	default:
		ok = false
	}

	return
}

func sectionType(i byte) (t SectionType, ok bool) {
	ok = true
	switch i {
	case 0:
		t = SECTION_IDENTIFICATION
	case 1:
		t = SECTION_AUTHENTICATION
	case 2:
		t = SECTION_MODULE
	case 3:
		t = SECTION_MODULE_PROPERTY
	case 4:
		t = SECTION_MODULE_PROPERTY_VALUE
	case 5:
		t = SECTION_MODULE_PROPERTY_DISABLED
	case 6:
		t = SECTION_COMMAND
	case 7:
		t = SECTION_COMMAND_ARGUMENT
	case 8:
		t = SECTION_COMMAND_EXECUTE
	case 9:
		t = SECTION_SUPPORTED
	default:
		ok = false
	}

	return
}
