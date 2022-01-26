package mqtt

import (
	"bytes"
	"fmt"
	"testing"
)

func TestCodec(t *testing.T) {
	writer := &bytes.Buffer{}
	m := NewMQTTCodec(writer, writer)
	m.writeVarLength(1048575)
	fmt.Printf("%v\n", writer.Bytes())
	reader := bytes.NewBuffer(writer.Bytes())
	m = NewMQTTCodec(reader, reader)
	n, err := m.readVarLength()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(n)
}
