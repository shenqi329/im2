package accessserver

import (
	coder "im/protocol/coder"
	"net"
)

type (
	Context interface {
		Server() *Server
		Request() *Request
		Conn() *net.TCPConn
		Message() *coder.Message
		CloseChan() chan<- uint32
		ReqChan() chan<- *Request
		Close()
	}
	context struct {
		server    *Server
		request   *Request
		conn      *net.TCPConn
		message   *coder.Message
		closeChan chan<- uint32
		reqChan   chan<- *Request
	}
)

func (c *context) Server() *Server {
	return c.server
}
func (c *context) Request() *Request {
	return c.request
}
func (c *context) Conn() *net.TCPConn {
	return c.conn
}
func (c *context) Message() *coder.Message {
	return c.message
}

func (c *context) CloseChan() chan<- uint32 {
	return c.closeChan
}
func (c *context) ReqChan() chan<- *Request {
	return c.reqChan
}
func (c *context) Close() {
	c.request.conn.Close()
	c.closeChan <- c.request.connId
}
