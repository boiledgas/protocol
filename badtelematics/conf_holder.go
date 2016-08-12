package telematics

import (
	"bytes"
	"encoding/gob"
)

type ConfHolder struct {
	hash byte
	conf Configuration
}

func NewHolder() *ConfHolder {
	conf := NewConfiguration()
	holder := &ConfHolder{conf: *conf}
	return holder
}

func (h *ConfHolder) Conf(c *Configuration) {
	if h.hash != c.hash {
		copy_conf(&h.conf, c)
	}
}

func (h *ConfHolder) SetConf(c *Configuration) {
	copy_conf(c, &h.conf)
}

func copy_conf(src *Configuration, dst *Configuration) {
	var buf bytes.Buffer
	enc := gob.NewEncoder(&buf)
	dec := gob.NewDecoder(&buf)
	if err := enc.Encode(src); err != nil {
		panic(err)
	}
	if err := dec.Decode(dst); err != nil {
		panic(err)
	}
}
