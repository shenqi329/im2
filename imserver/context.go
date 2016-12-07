package imserver

import (
	"errors"
	"github.com/golang/protobuf/proto"
	imServerBean "im/imserver/bean"
	dao "im/imserver/dao"
	imError "im/imserver/error"
	protocolClient "im/protocol/client"
	protocolCoder "im/protocol/coder"
	protocolServer "im/protocol/server"
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
		NeedWraper() bool
		SendProtoMessage(messageType protocolClient.MessageType, message proto.Message) error
		ConnInfoChan() chan<- *ConnInfo
		TokenConnInfoChan() chan<- int64
	}

	context struct {
		imServer          *Server
		udpConn           *net.UDPConn
		udpAddr           *net.UDPAddr
		protoMessage      proto.Message
		connId            uint64
		needWraper        bool
		connInfoChan      chan<- *ConnInfo
		tokenConnInfoChan chan<- int64
	}
)

func (c *context) ConnInfoChan() chan<- *ConnInfo {
	return c.connInfoChan
}
func (c *context) TokenConnInfoChan() chan<- int64 {
	return c.tokenConnInfoChan
}

func (c *context) NeedWraper() bool {
	return c.needWraper
}

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

func (c *context) SendProtoMessage(messageType protocolClient.MessageType, message proto.Message) error {

	buffer, err := protocolCoder.EncoderProtoMessage((int)(messageType), message)
	if err != nil {
		log.Println(err)
		return err
	}

	if c.needWraper {
		//包装数据后返回
		//log.Println("包装数据")
		wraperMessage := &protocolServer.WraperMessage{
			ConnId:  c.ConnId(),
			Message: buffer,
		}

		buffer, err = protocolCoder.EncoderProtoMessage(protocolServer.MessageTypeWraper, wraperMessage)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	count, err := c.UDPConn().WriteTo(buffer, c.udpAddr)

	if count != len(buffer) {
		err = errors.New("写入数据失败")
		log.Println(err)
		return err
	}
	return nil
}

func NewCommonResponseWithError(err error, rid uint64) *protocolClient.CommonResponse {

	imErr, ok := err.(*imError.IMError)
	if ok {
		response := &protocolClient.CommonResponse{
			Rid:  rid,
			Code: imErr.Code,
			Desc: imErr.Desc,
		}
		return response
	}

	response := &protocolClient.CommonResponse{
		Rid:  rid,
		Code: imError.CommonInternalServerError,
		Desc: imError.ErrorCodeToText(imError.CommonInternalServerError),
	}
	return response
}

//发送同步通知
func SendSyncInform(udpAddr *net.UDPAddr, udpConn *net.UDPConn, connId uint64, userId string) {

	var sessionMaps []*imServerBean.SessionMap

	err := dao.NewDao().Find(&sessionMaps, &imServerBean.SessionMap{
		UserId: userId,
	})
	if err != nil {
		log.Println(err)
		return
	}

	for _, sessionMap := range sessionMaps {
		SendSyncInformWithSessionMap(udpAddr, udpConn, sessionMap, connId)
	}
}

func SendSyncInformWithSessionMap(udpAddr *net.UDPAddr, udpConn *net.UDPConn, sessionMap *imServerBean.SessionMap, connId uint64) {

	var messages []*imServerBean.Message

	err := dao.NewDao().Find(&messages, &imServerBean.Message{
		SessionId: sessionMap.SessionId,
	})

	latestIndex, err := dao.MessageMaxIndex(sessionMap.SessionId)

	if sessionMap.ReadIndex >= latestIndex {
		//log.Println("不需发送同步通知")
		return
	}

	if err != nil {
		log.Println(err)
		return
	}
	log.Println(latestIndex)

	syncInfo := &protocolClient.SyncInform{
		SessionId:   sessionMap.SessionId,
		LatestIndex: latestIndex,
		ReadIndex:   sessionMap.ReadIndex,
	}

	SendProtoMessage(udpAddr, udpConn, protocolClient.MessageTypeSyncInform, syncInfo, connId, true)
}

func SendProtoMessage(udpAddr *net.UDPAddr, udpConn *net.UDPConn, messageType protocolClient.MessageType, message proto.Message, connId uint64, needWraper bool) error {

	buffer, err := protocolCoder.EncoderProtoMessage((int)(messageType), message)
	if err != nil {
		log.Println(err)
		return err
	}

	if needWraper {
		//包装数据后返回
		//log.Println("包装数据")
		wraperMessage := &protocolServer.WraperMessage{
			ConnId:  connId,
			Message: buffer,
		}

		buffer, err = protocolCoder.EncoderProtoMessage(protocolServer.MessageTypeWraper, wraperMessage)
		if err != nil {
			log.Println(err)
			return err
		}
	}

	count, err := udpConn.WriteTo(buffer, udpAddr)

	if count != len(buffer) {
		err = errors.New("写入数据失败")
		log.Println(err)
		return err
	}
	return nil
}
