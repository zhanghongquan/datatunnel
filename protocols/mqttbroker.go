package protocols

import (
	"net"

	"github.com/zhanghongquan/datatunnel/protocols/mqtt"
)

type IMqttAction interface {
	OnAuth(*mqtt.Auth) error

	OnConAck(*mqtt.Connack) error

	OnConnect(*mqtt.Connect) error

	OnDisconnect(*mqtt.Disconnect) error

	OnPingReq(*mqtt.Pingreq) error

	OnPingResp(*mqtt.Pingresp) error

	OnPubAck(*mqtt.Puback) error

	OnPubComp(*mqtt.Pubcomp) error

	OnPublish(*mqtt.Publish) error

	OnPubRec(*mqtt.Pubrec) error

	OnPubRel(*mqtt.Pubrel) error

	OnSubAck(*mqtt.Suback) error

	OnSubscribe(*mqtt.Subscribe) error

	OnUnsubAck(*mqtt.Unsuback) error

	OnUnsubscribe(*mqtt.Unsubscribe) error
}

type MQTTConfig struct {
	Listen string
}

type MQTTBroker struct {
	config   *MQTTConfig
	listener net.Listener
}

func NewMQTTBroker(config *MQTTConfig) (*MQTTBroker, error) {

}

func (m *MQTTBroker) Start() error {

}

func (m *MQTTBroker) Stop() {

}

func (m *MQTTBroker) run() error {

}
