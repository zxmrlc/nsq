package nsqlookupd

import (
	"net"
)

//client 包含net 链接及具体信息
type ClientV1 struct {
	net.Conn
	peerInfo *PeerInfo
}

func NewClientV1(conn net.Conn) *ClientV1 {
	return &ClientV1{
		Conn: conn,
	}
}

func (c *ClientV1) String() string {
	return c.RemoteAddr().String()
}
