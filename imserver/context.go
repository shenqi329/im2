package imserver

import (
	"errors"
	"github.com/golang/protobuf/proto"
	imError "im/imserver/error"
	protocolBean "im/protocol/bean"
	protocolCoder "im/protocol/coder"
	"log"
	"net"
)

type (
	Context interface {
		IMServer() *Server
		UDPConn() *net.UDPConn
		UDPAddr() *net.UDPAddr
		ProtoMessage() proto.Message
		ConnId() uint64
		WraperProtoMessage(messageType protocolBean.MessageType, message proto.Message) error
	}

	context struct {
		imServer     *Server
		udpConn      *net.UDPConn
		udpAddr      *net.UDPAddr
		protoMessage proto.Message
		connId       uint64
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

func (c *context) ConnId() uint64 {
	return c.connId
}

func (c *context) IMServer() *Server {
	return c.imServer
}

func (c *context) WraperProtoMessage(messageType protocolBean.MessageType, message proto.Message) error {

	buffer, err := protocolCoder.EncoderProtoMessage((int)(messageType), message)
	if err != nil {
		log.Println(err)
		return err
	}

	//包装数据后返回
	wraperMessage := &protocolBean.WraperMessage{
		ConnId:  c.ConnId(),
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

func NewCommonResponseWithError(err error, rid uint64) *protocolBean.CommonResponse {

	imErr, ok := err.(*imError.IMError)
	if ok {
		response := &protocolBean.CommonResponse{
			Rid:  rid,
			Code: imErr.Code,
			Desc: imErr.Desc,
		}
		return response
	}

	response := &protocolBean.CommonResponse{
		Rid:  rid,
		Code: imError.CommonInternalServerError,
		Desc: imError.ErrorCodeToText(imError.CommonInternalServerError),
	}
	return response
}
