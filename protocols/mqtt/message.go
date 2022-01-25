package mqtt

const (
	MSG_RESERVED byte = iota
	MSG_CONNECT
	MSG_CONNACK
	MSG_PUBLISH
	MSG_PUBACK
	MSG_PUBREC
	MSG_PUBREL
	MSG_PUBCOMP
	MSG_SUBSCRIBE
	MSG_SUBACK
	MSG_UNSCRIBE
	MSG_UNSUBACK
	MSG_PINGREQ
	MSG_PINGRESP
	MSG_DISCONNECT
	MSG_RESERVED2
)

type FixedHeader struct {
	H byte
}

func (h FixedHeader) MessageType() byte {
	return h.H & 0xf0
}

func (h FixedHeader) Qos() byte {
	return h.H & 0x06
}

func (h FixedHeader) Retained() bool {
	return h.H&0x01 != 0
}

func (h FixedHeader) Duplicated() bool {
	return h.H&0x04 != 0
}

type ConnectMsg struct {
	FixedHeader
}

type PublishMsg struct {
	FixedHeader
}

type DisconnectMsg struct {
	FixedHeader
}

type ConnAckMsg struct {
	FixedHeader
}

type PubAckMsg struct {
	FixedHeader
}

type PubRecMsg struct {
	FixedHeader
}

type PubRelMsg struct {
	FixedHeader
}

type PubCompMsg struct {
	FixedHeader
}
