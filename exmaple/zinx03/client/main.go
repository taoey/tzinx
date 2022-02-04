package main

import (
	"fmt"
	"net"
	"time"
)

// 客户端模拟
func main() {
	fmt.Println("client running...")

	// 1、 连接远程服务器，得到connect
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("conn err", err)
		return
	}

	// 测试多次调用
	count := 10
	for i := 0; i < count; i++ {

		// 2、 调用write方法
		body := fmt.Sprintf("hello world %d", i)
		fmt.Println(body)
		_, err := conn.Write([]byte(body))
		if err != nil {
			fmt.Println("write err", err)
			return
		}

		// 3、 读取返回信息
		buf := make([]byte, 512)
		cnt, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error")
			return
		}
		fmt.Printf("server call back: %s, len:%d \n", string(buf[0:cnt]), cnt)

		time.Sleep(time.Second)
	}
}
