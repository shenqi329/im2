package server

import (
	"errors"
	"github.com/golang/protobuf/proto"
	imServiceResponse "im/imserver/response"
	imService "im/imserver/service"
	protocalBean "im/protocal/bean"
	"im/protocal/coder"
	"log"
	"net"
	"os"
	"reflect"
)

type Request struct {
	isCancel bool
	reqPkg   []byte
	rspChan  chan<- []byte
}

type Server struct {
	localUdpAddr string
}

func NEWServer(localUdpAddr string) *Server {
	return &Server{
		localUdpAddr: localUdpAddr,
	}
}

func (s *Server) Run() {
	s.listenOnUdpPort(s.localUdpAddr)
}

func (s *Server) listenOnUdpPort(localUdpAddr string) {

	addr, err := net.ResolveUDPAddr("udp", localUdpAddr)

	if err != nil {
		log.Println("net.ResolveUDPAddr fail.", err)
		os.Exit(1)
	}

	conn, err := net.ListenUDP("udp", addr)
	defer conn.Close()

	if err != nil {
		log.Println("net.ListenUDP fail.", err)
		os.Exit(1)
	}

	log.Println("net.ListenUDP", addr)

	reqChan := make(chan *Request, 1000)
	var recvAndSendCount uint32 = 0

	for true {
		buf := make([]byte, 1024)
		rlen, remote, err := conn.ReadFromUDP(buf)
		if err != nil {
			log.Println("读取数据失败!", err.Error())
			continue
		}
		recvAndSendCount++
		//log.Println("recvAndSendCount:", recvAndSendCount, " rlen:", rlen)
		go s.processHandler(conn, remote, buf[:rlen], reqChan)
	}
}

func (s *Server) processHandler(conn *net.UDPConn, remote *net.UDPAddr, msg []byte, reqChan chan<- *Request) {

	decoder := coder.NEWDecoder()
	beanWraperMessages, err := decoder.Decode(msg)
	if err != nil {
		log.Println(err.Error())
		return
	}
	//处理
	for _, beanWraperMessage := range beanWraperMessages {
		if beanWraperMessage.Type != protocalBean.MessageTypeWraper {
			continue
		}

		wraperMessage := &protocalBean.WraperMessage{}
		if err := proto.Unmarshal(beanWraperMessage.Body, wraperMessage); err != nil {
			log.Println(err)
			continue
		}

		decoder.Reset()
		beanMessages, err := decoder.Decode(wraperMessage.Message)
		if err != nil {
			log.Println(err.Error())
			continue
		}
		for _, beanMessage := range beanMessages {
			s.handleMessage(conn, remote, beanMessage, reqChan, wraperMessage.Rid)
		}
	}
}

func (s *Server) handleMessage(conn *net.UDPConn, addr *net.UDPAddr, message *coder.Message, reqChan chan<- *Request, rid uint64) {

	protoMessage := protocalBean.Factory((protocalBean.MessageType)(message.Type))
	if err := proto.Unmarshal(message.Body, protoMessage); err != nil {
		log.Println(err.Error())
		return
	}
	switch message.Type {
	case protocalBean.MessageTypeDeviceRegisteRequest:
		{
			s.handleRegisteRequest(conn, addr, protoMessage.(*protocalBean.DeviceRegisteRequest), rid)
		}
	case protocalBean.MessageTypeDeviceLoginRequest:
		{
			s.handleLoginRequest(conn, addr, protoMessage.(*protocalBean.DeviceLoginRequest), rid)
		}
	}
}

func (s *Server) handleRegisteRequest(conn *net.UDPConn, addr *net.UDPAddr, deviceRegisteRequest *protocalBean.DeviceRegisteRequest, rid uint64) error {

	tokenBean, err := imService.HandleRegisteRequest(deviceRegisteRequest)

	if err != nil {
		return s.ServiceHandleError(err, conn, addr, deviceRegisteRequest, protocalBean.MessageTypeDeviceRegisteResponse)
	}

	return nil
	// response := &bean.DeviceRegisteResponse{
	// 	Rid:   deviceRegisteRequest.Rid,
	// 	Code:  "00000001",
	// 	Desc:  "success",
	// 	Token: fmt.Sprintf("%d", rid),
	// }

	// client := &http.Client{}
	// httpRequest, err := http.NewRequest(http.MethodGet, "http://172.17.0.3:8081/user/info", nil)
	// httpRequest.Header.Set("token", deviceRegisteRequest.SsoToken)
	// resp, err := client.Do(httpRequest)
	// defer resp.Body.Close()

	// responseBean := &imResponse.Response{}
	// json.NewDecoder(req.Body()).Decode(responseBean)

	// if !strings.EqualFold(responseBean.Code, imResponse.CommonSuccess) {
	// 	log.Println(responseBean.Desc)
	// }

	// buffer, err := coder.EncoderProtoMessage(bean.MessageTypeDeviceRegisteResponse, response)
	// if err != nil {
	// 	log.Println(err)
	// 	return
	// }
	// s.wraperMessageAndSendBack(conn, addr, reqChan, rid, buffer)
}

func (s *Server) ServiceHandleError(err error, conn *net.UDPConn, addr *net.UDPAddr, deviceRegisteRequest *protocalBean.DeviceRegisteRequest, messageType protocalBean.MessageType) error {

	e := reflect.Indirect(reflect.ValueOf(err))

	response := &imServiceResponse.Response{}

	code := e.FieldByName("Code")
	if code.Kind() == reflect.String {
		response.Code = code.String()
	}
	desc := e.FieldByName("Desc")
	if desc.Kind() == reflect.String {
		response.Desc = desc.String()
	}

	registerResponse := &protocalBean.DeviceRegisteResponse{
		Rid:  deviceRegisteRequest.Rid,
		Code: response.Code,
		Desc: response.Desc,
	}

	buffer, err := coder.EncoderProtoMessage((int)(messageType), registerResponse)
	if err != nil {
		log.Println(err)
		return err
	}
	return s.wraperMessageAndSendBack(conn, addr, rid, buffer)
}

func (s *Server) handleLoginRequest(conn *net.UDPConn, addr *net.UDPAddr, request *bean.DeviceLoginRequest, rid uint64) error {

	response := &bean.DeviceLoginResponse{
		Rid:  request.Rid,
		Code: "00000001",
		Desc: "success",
	}

	buffer, err := coder.EncoderProtoMessage((int)(bean.MessageTypeDeviceLoginResponse), response)
	if err != nil {
		log.Println(err)
		return err
	}
	return s.wraperMessageAndSendBack(conn, addr, rid, buffer)
}

func (s *Server) wraperMessageAndSendBack(conn *net.UDPConn, addr *net.UDPAddr, rid uint64, buffer []byte) error {
	wraperMessage := &bean.WraperMessage{
		Rid:     rid,
		Message: buffer,
	}
	buffer, err := coder.EncoderProtoMessage(bean.MessageTypeWraper, wraperMessage)
	if err != nil {
		log.Println(err)
		return err
	}

	count, err := conn.WriteTo(buffer, addr)
	if count != len(buffer) {
		err = errors.New("写入数据失败")
		log.Println(err)
		return err
	}
	return nil
}
