package mqtt

import (
	"bufio"
	"encoding/binary"
	"fmt"
	"io"
)

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

func readVarLength(reader *bufio.Reader) (int, error) {
	count := 0
	var multiplier int = 1
	var value int
	var digit byte = 128
	var err error
	for (digit&128) != 0 && count < 4 {
		digit, err = reader.ReadByte()
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

func writeVarLength(writer io.Writer, data int) error {
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
		_, err := writer.Write(buf[:])
		if err != nil {
			return err
		}
	}
	return nil
}

func read16BitData(reader io.Reader) (uint16, error) {
	var data [2]byte
	_, err := io.ReadAtLeast(reader, data[:], 2)
	if err != nil {
		return 0, nil
	}
	return binary.BigEndian.Uint16(data[:]), nil
}

func write16BitData(writer io.Writer, data uint16) error {
	var buf [2]byte
	binary.BigEndian.PutUint16(buf[:], data)
	n, err := writer.Write(buf[:])
	if n != 2 {
		return fmt.Errorf("insufficient room to write 16bit data into")
	}
	return err
}

func readMqttString(reader *bufio.Reader) (string, error) {
	length, err := read16BitData(reader)
	if err != nil {
		return "", err
	}
	if length == 0 {
		return "", nil
	}
	data := make([]byte, length)
	_, err = io.ReadAtLeast(reader, data, int(length))
	if err != nil {
		return "", err
	}
	//TODO: should be tested if data is a valid utf8 string ?
	return string(data), nil
}

func writeMqqString(writer *bufio.Writer, data string) error {
	err := write16BitData(writer, uint16(len(data)))
	if err != nil {
		return err
	}

	n, err := writer.WriteString(data)
	if n != len(data) {
		return fmt.Errorf("no room to write string data into")
	}
	return err
}

type FixedHeader struct {
	H byte
}

func (h *FixedHeader) Marshal(writer *bufio.Writer) error {
	return writer.WriteByte(h.H)
}

func (h *FixedHeader) Unmarshal(reader *bufio.Reader) (err error) {
	h.H, err = reader.ReadByte()
	return
}

func (h *FixedHeader) MessageType() byte {
	return h.H & 0xf0
}

func (h *FixedHeader) Qos() byte {
	return h.H & 0x06
}

func (h *FixedHeader) Retained() bool {
	return h.H&0x01 != 0
}

func (h *FixedHeader) Duplicated() bool {
	return h.H&0x04 != 0
}

type ConnectMsg struct {
	FixedHeader
	ProtocolLevel  byte
	KeepAlive      uint16
	ClientIdentity string
	WillTopic      string
	WillMessage    string
	UserName       string
	Password       string
}

func (m *ConnectMsg) Marshal(writer *bufio.Writer) (n int, err error) {
	err = m.FixedHeader.Marshal(writer)
	if err != nil {
		return 0, err
	}

	//
}

func (m *ConnectMsg) Unmarshal(reader *bufio.Reader) (err error) {
	err = m.FixedHeader.Unmarshal(reader)
	if err != nil {
		return err
	}

	var length int
	length, err = readVarLength(reader)
	if err != nil {
		return err
	}
	// parse variable header

	// parse protocol name and level
	var protocolName string
	protocolName, err = readMqttString(reader)
	if err != nil {
		return err
	}

	var protocolLevel byte
	protocolLevel, err = reader.ReadByte()
	if err != nil {
		return err
	}

	switch protocolLevel {
	case 3:
	case 4:
	case 5:
	default:
		return fmt.Errorf("unkown mqtt protocol level: %d", protocolLevel)
	}

	var connectFlag byte
	connectFlag, err = reader.ReadByte()
	if err != nil {
		return err
	}
}

func (m *ConnectMsg) parseV3() error {

}

func (m *ConnectMsg) parseV4() error {

}

func (m *ConnectMsg) parsev5() error {

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
