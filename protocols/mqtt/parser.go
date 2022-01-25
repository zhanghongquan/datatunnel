package mqtt

import (
	"io"

	v5 "github.com/eclipse/paho.golang/packets"
)

type IMqttActionV3 interface {
	OnConnected(*ConnectMsg) error

	OnPublish(*PublishMsg) error

	OnDisconnected(*DisconnectMsg) error
}

type MQTTParserV3 struct {
	reader io.Reader
	cb     IMqttActionV3
	buf    [128]byte
	// bufDataStart和bufDataEnd用来表示buf所存放的可用数据的开始和结束地址
	// 当数据被使用以后, bufDataStart不会断的向后移动。
	bufDataStart int
	bufDataEnd   int
}

func NewMQTTParser(reader io.Reader) *MQTTParserV3 {
	v3 := &MQTTParserV3{reader: reader}
	return v3
}

func (p *MQTTParserV3) Process() error {
	v5.ReadPacket(p.reader)
}

func (p *MQTTParserV3) readVarLength() (int, error) {
	msb := true
	buffer := p.buf[:]
	for msb {

	}
}

func (p *MQTTParserV3) read16BitLength() (uint16, error) {

}

func (p *MQTTParserV3) readRawData(wanted int) ([]byte, error) {
	if p.bufDataEnd <= p.bufDataStart || (p.bufDataEnd-p.bufDataStart) < wanted {
		//read more data into buffer due to data insufficient
	}
}
