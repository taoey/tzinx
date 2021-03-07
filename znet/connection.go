package znet

import (
	"fmt"
	"github.com/taoey/tzinx/ziface"
	"net"
)

// 连接模块
type Connection struct {
	Conn *net.TCPConn

	ConnId uint32

	isClose bool

	HandleAPI ziface.HandleFunc

	Exit chan bool
}

// 初始化连接模块方法
func NewConnection(conn *net.TCPConn, connID uint32, callbackAPI ziface.HandleFunc) *Connection {
	c := &Connection{
		Conn:      conn,
		ConnId:    connID,
		HandleAPI: callbackAPI,
		isClose:   false,
		Exit:      make(chan bool, 1),
	}
	return c
}

// 连接的读取业务
func (c *Connection) StartReader() {
	fmt.Println("read goroutine is running...")
	defer c.Stop()

	for {
		// 读取客户端的数据到buf中
		buf := make([]byte, 512)
		cnt, err := c.Conn.Read(buf)
		if err != nil {
			fmt.Println("recv read err", err)
			continue
		}

		// 调用当前连接所绑定的api
		err = c.HandleAPI(c.Conn, buf, cnt)
		if err != nil {
			fmt.Println("conn handle err", err)
			break
		}

	}
}

func (c *Connection) Start() {
	fmt.Println("conn start... connid=", c.ConnId)

	// 读取数据
	go c.StartReader()

}

func (c *Connection) Stop() {
	fmt.Println("conn stop... connid=", c.ConnId)

	if c.isClose == true {
		return
	}

	c.isClose = true

	// 关闭socket连接
	c.Conn.Close()

	close(c.Exit)

}

func (c *Connection) Send([]byte) error {

	return nil
}

func (c *Connection) GetTCPConnection() *net.TCPConn {
	return c.Conn
}

func (c *Connection) GetConnID() uint32 {
	return c.ConnId
}

func (c *Connection) RemoteAddr() net.Addr {
	return c.Conn.RemoteAddr()
}
