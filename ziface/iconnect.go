package ziface

import "net"

// 连接模块
type IConnect interface {
	Start()
	Stop()
	Send([]byte) error
	GetTCPConnection() *net.TCPConn
	GetConnID() uint32
	RemoteAddr() net.Addr
}

// 处理连接业务的方法
// conn: 需要处理的tcp连接
// body: 需要处理的内容
// lenght: 需要的长度
type HandleFunc func(conn *net.TCPConn, body []byte, length int) error
