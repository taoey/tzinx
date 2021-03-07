package znet

import (
	"fmt"
	"github.com/taoey/tzinx/ziface"
	"net"
)

type Server struct {
	Name      string // 服务器名称
	IPVersion string // 服务器IP版本
	IP        string // 服务器监听IP
	Port      int    // 监听端口号
}

func NewServer(name string, port int) ziface.IServer {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      port,
	}
	return s
}

//-- iserver接口实现

func (s *Server) Start() {

	// 单独启动一个协程，进行监听
	go func() {

		//1、 获取tcp addr
		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp err:", err)
			return
		}

		// 2、监听服务器地址
		listener, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("listen tcp err:", err)
			return
		}

		fmt.Println("start Zinx server  ", s.Name, s.Port, " succ, now listenning...")

		// 3、 阻塞等待客户端连接，处理客户端连接业务
		for {

			// 如果有客户端连接过来，阻塞会返回数据
			conn, err := listener.AcceptTCP()
			if err != nil {
				fmt.Println("Accept err：", err)
				return
			}
			fmt.Println("Get conn remote addr = ", conn.RemoteAddr().String())

			// 已经建立连接，启动go协程处理相关业务
			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("recv buf err", err)
						continue
					}
					fmt.Println("recv client:", string(buf[0:cnt]))

					// 回显
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err", err)
						continue
					}

				}
			}()

		}
	}()

}

func (s *Server) Server() {
	s.Start()

	// TODO 服务器启动之后需要做的任务

	// 阻塞，防止主进程退出导致监听退出
	select {}
}

func (s *Server) Stop() {
	// TODO 回收服务器资源
}
