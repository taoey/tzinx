package main

import (
	"fmt"

	"github.com/taoey/tzinx/ziface"
	"github.com/taoey/tzinx/znet"
)

// demo: zinx 服务启动

func main() {

	// 创建一个服务器
	s := znet.NewServer("[zinx02]", 8999)

	s.AddRouter(&PingRouter{})

	// 启动服务器
	s.Server()

}

type PingRouter struct {
	znet.BaseRouter
}

func (p *PingRouter) PreHandle(request ziface.IRequest) {
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("before ping "))
	if err != nil {
		fmt.Println("pre handle err", err)
	}

}

func (p *PingRouter) Handle(request ziface.IRequest) {
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("ping "))
	if err != nil {
		fmt.Println("handle err", err)
	}
}

func (p *PingRouter) PostHandle(request ziface.IRequest) {
	connection := request.GetConnection()
	_, err := connection.GetTCPConnection().Write([]byte("after ping "))
	if err != nil {
		fmt.Println("after handle err", err)
	}
}
