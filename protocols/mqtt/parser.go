package mqtt

import (
	"bufio"
	"io"
)

type IMqttAction interface {
	OnConnected(*ConnectMsg) error

	OnPublish(*PublishMsg) error

	OnDisconnected(*DisconnectMsg) error
}

type MQTTCodec struct {
	reader *bufio.Reader
	writer *bufio.Writer
	cb     IMqttAction
}

func NewMQTTCodec(reader io.Reader, writer io.Writer) *MQTTCodec {
	//create buffer io with default buffer size(4K)
	codec := &MQTTCodec{
		reader: bufio.NewReader(reader),
		writer: bufio.NewWriter(writer)}
	return codec
}

func (p *MQTTCodec) Process() error {
	return nil
}
