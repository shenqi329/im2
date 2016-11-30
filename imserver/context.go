package imserver

import (
	"errors"
	"github.com/golang/protobuf/proto"
	protocolBean "im/protocol/bean"
	protocolCoder "im/protocol/coder"
	"log"
	"net"
)

type (
	Context interface {
		UDPConn() *net.UDPConn
		UDPAddr() *net.UDPAddr
		ProtoMessage() proto.Message
		Rid() uint64
		WraperProtoMessage(messageType protocolBean.MessageType, message proto.Message) error
	}

	context struct {
		udpConn      *net.UDPConn
		udpAddr      *net.UDPAddr
		protoMessage proto.Message
		rid          uint64
	}
)

func (c *context) UDPConn() *net.UDPConn {
	return c.udpConn
}

func (c *context) UDPAddr() *net.UDPAddr {
	return c.udpAddr
}

func (c *context) ProtoMessage() proto.Message {
	return c.protoMessage
}

func (c *context) Rid() uint64 {
	return c.rid
}

func (c *context) WraperProtoMessage(messageType protocolBean.MessageType, message proto.Message) error {

	buffer, err := protocolCoder.EncoderProtoMessage((int)(messageType), message)
	if err != nil {
		log.Println(err)
		return err
	}

	//包装数据后返回
	wraperMessage := &protocolBean.WraperMessage{
		Rid:     c.Rid(),
		Message: buffer,
	}

	buffer, err = protocolCoder.EncoderProtoMessage(protocolBean.MessageTypeWraper, wraperMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	count, err := c.UDPConn().WriteTo(buffer, c.udpAddr)

	if count != len(buffer) {
		err = errors.New("写入数据失败")
		log.Println(err)
		return err
	}
	return nil
}
