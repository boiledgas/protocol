package telematics

type GsmStruct struct {
	mcc    string
	mnc    string
	lac    uint16
	cid    uint16
	signal int8
}

type GsmValue interface {
	GetMCC() string
	SetMCC(string)

	GetMNC() string
	SetMNC(string)

	GetLAC() uint16
	SetLAC(uint16)

	GetCID() uint16
	SetCID(uint16)

	GetSignal() int8
	SetSignal(int8)
}

func (c *GsmStruct) GetMCC() (val string) {
	val = c.mcc
	return
}

func (c *GsmStruct) SetMCC(val string) {
	c.mcc = val
}

func (c *GsmStruct) GetMNC() (val string) {
	return c.mnc
}

func (c *GsmStruct) SetMNC(val string) {
	c.mnc = val
}

func (c *GsmStruct) GetLAC() (val uint16) {
	return c.lac
}

func (c *GsmStruct) SetLAC(val uint16) {
	c.lac = val
}

func (c *GsmStruct) GetCID() (val uint16) {
	return c.cid
}

func (c *GsmStruct) SetCID(val uint16) {
	c.cid = val
}

func (c *GsmStruct) GetSignal() (val int8) {
	return c.signal
}

func (c *GsmStruct) SetSignal(val int8) {
	c.signal = val
}
