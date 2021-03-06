package main

import "github.com/taoey/tzinx/znet"

// demo: zinx 服务启动

func main() {

	// 创建一个服务器
	s := znet.NewServer("[zinx01]", 8080)

	// 启动服务器
	s.Server()
}
