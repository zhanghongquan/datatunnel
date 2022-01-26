package mqtt

import (
	"bufio"
	"fmt"
	"io"
)

type IMqttActionV3 interface {
	OnConnected(*ConnectMsg) error

	OnPublish(*PublishMsg) error

	OnDisconnected(*DisconnectMsg) error
}

type MQTTCodec struct {
	reader bufio.Reader
	writer bufio.Writer
	cb     IMqttActionV3
}

func NewMQTTParser(reader io.Reader) *MQTTCodec {
	//create buffer io with default buffer size(4K)
	v3 := &MQTTCodec{reader: *bufio.NewReader(reader)}
	return v3
}

func (p *MQTTCodec) Process() error {
}

func (p *MQTTCodec) readVarLength() (int, error) {
	count := 0
	var multiplier int = 1
	var value int
	var digit byte = 128
	for (digit&128) != 0 && count < 4 {
		digit, err := p.reader.ReadByte()
		if err != nil {
			return 0, err
		}
		value += int(digit&127) * multiplier
		multiplier *= 128
		count++
	}

	if digit&128 != 0 {
		return 0, fmt.Errorf("bad encodings, more than 4 bytes used to represent a number")
	}
	return value, nil
}

func (p *MQTTCodec) writeVarLength(data int) error {
	if data > 268435455 {
		return fmt.Errorf("data too huge to encoding: %d", data)
	}
	var buf [1]byte
	for data > 0 {
		digit := byte(data % 128)
		data = data / 128
		if data > 0 {
			digit = digit | 0x80
		}
		buf[0] = digit
		_, err := p.writer.Write(buf[:])
		if err != nil {
			return err
		}
	}
	return nil
}

func (p *MQTTCodec) read16BitLength() (uint16, error) {

}
