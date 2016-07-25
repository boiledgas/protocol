package telematics

import "fmt"

type modulesMap struct {
	count byte
	items [255]Module
}

// Device configuration
type conf struct {
	hash       byte
	modules    modulesMap
	prop_count uint16
	properties map[byte]map[byte]ModuleProperty
	cmd_count  uint16
	commands   map[byte]map[byte]Command
	arg_count  uint16
	arguments  map[byte]map[byte]map[byte]CommandArgument
}

type Conf interface {
	Modules() []Module
	Properties() []ModuleProperty
	Commands() []Command
	Arguments() []CommandArgument
}

func (c *conf) Reset() {
	c.modules = modulesMap{}
	c.properties = make(map[byte]map[byte]ModuleProperty)
	c.commands = make(map[byte]map[byte]Command)
	c.arguments = make(map[byte]map[byte]map[byte]CommandArgument)
}

func NewConfiguration() *conf {
	c := conf{}
	c.Reset()
	return &c
}

func (c *conf) SetModule(m Module) {
	mid := m.GetId()
	if c.modules.items[mid] != nil {
		panic(fmt.Sprintf("module with id %d exists", mid))
	}

	c.modules.items[mid] = m
}

func (c *conf) GetModule(id byte) (m Module, ok bool) {
	m = c.modules.items[id]
	ok = m != nil
	return
}

func (c *conf) Modules() []Module {
	r := make([]Module, 0, int(c.modules.count))
	i := 0
	for _, m := range c.modules.items {
		if m != nil {
			r[i] = m
			i++
		}
	}
	return r
}

func (c *conf) SetProperty(p ModuleProperty) {
	mid := p.GetModuleId()
	if c.modules.items[mid] == nil {
		panic(fmt.Sprintf("module with id %d not found", mid))
	}

	if _, ok := c.properties[mid]; !ok {
		c.properties[mid] = make(map[byte]ModuleProperty)
	}

	pid := p.GetId()
	if _, ok := c.properties[mid][pid]; ok {
		panic(fmt.Sprintf("module property with id %d exists", pid))
	}

	c.properties[mid][pid] = p
	c.prop_count++
}

func (c *conf) GetProperty(mid byte, pid byte) (p ModuleProperty, ok bool) {
	var pmap map[byte]ModuleProperty
	if pmap, ok = c.properties[mid]; !ok {
		return
	}

	p, ok = pmap[pid]
	return
}

func (c *conf) Properties() []ModuleProperty {
	r := make([]ModuleProperty, c.prop_count)
	i := 0
	for _, pmap := range c.properties {
		for _, p := range pmap {
			r[i] = p
			i++
		}
	}
	return r
}

func (c *conf) SetCommand(mc Command) {
	mid := mc.GetModuleId()
	if c.modules.items[mid] == nil {
		panic(fmt.Sprintf("module with id %d not found", mid))
	}

	if _, ok := c.commands[mid]; !ok {
		c.commands[mid] = make(map[byte]Command)
	}

	cid := mc.GetId()
	if _, ok := c.properties[mid][cid]; ok {
		panic(fmt.Sprintf("module command with id %d exists", cid))
	}

	c.commands[mid][cid] = mc
	c.cmd_count++
}

func (c *conf) GetCommand(mid byte, cid byte) (mc Command, ok bool) {
	var cmap map[byte]Command
	if cmap, ok = c.commands[cid]; !ok {
		return
	}

	mc, ok = cmap[cid]
	return
}

func (c *conf) Commands() []Command {
	r := make([]Command, c.cmd_count)
	i := 0
	for _, cmap := range c.commands {
		for _, c := range cmap {
			r[i] = c
			i++
		}
	}
	return r
}

func (c *conf) SetArgument(ca CommandArgument) {
	mid := ca.GetModuleId()
	cid := ca.GetCommandId()
	if _, ok := c.commands[mid]; ok {
		if _, ok := c.commands[mid][cid]; !ok {
			panic(fmt.Sprintf("command with id %d not found", cid))
		}
	} else {
		panic(fmt.Sprintf("command with id %d not found", cid))
	}

	if _, ok := c.arguments[mid]; !ok {
		c.arguments[mid] = make(map[byte]map[byte]CommandArgument)
	}
	if _, ok := c.arguments[mid][cid]; !ok {
		c.arguments[mid][cid] = make(map[byte]CommandArgument)
	}

	ca_id := ca.GetId()
	if _, ok := c.arguments[mid][cid][ca_id]; ok {
		panic(fmt.Sprintf("command argument with id %d exists", ca_id))
	}

	c.arguments[mid][cid][ca_id] = ca
	c.arg_count++
}

func (c *conf) GetArgument(mid byte, cid byte, ca_id byte) (ca CommandArgument, ok bool) {
	if _, ok = c.arguments[mid]; ok {
		if _, ok = c.arguments[mid][cid]; ok {
			if ca, ok = c.arguments[mid][cid][ca_id]; ok {
				return
			}
		}
	}

	ok = false
	return
}

func (c *conf) Arguments() []CommandArgument {
	r := make([]CommandArgument, c.arg_count)
	i := 0
	for _, amap := range c.arguments {
		for _, c := range amap {
			for _, ca := range c {
				r[i] = ca
				i++
			}
		}
	}
	return r

}
