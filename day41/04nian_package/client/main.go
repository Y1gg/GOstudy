package main

import (
	"fmt"
	proto "github/Y1gg/day41/05protocal"
	"net"
)

// socket_stick/client/main.go

func main() {
	conn, err := net.Dial("tcp", "127.0.0.1:30000")
	if err != nil {
		fmt.Println("dial failed, err", err)
		return
	}
	defer conn.Close()
	for i := 0; i < 20; i++ {
		msg := `Hello, Hello. How are you?`
		//调用协议编码数据
		b, err := proto.Encode(msg)
		if err != nil {
			fmt.Println("encode failed err:", err)
			return
		}
		conn.Write(b)
		// time.Sleep(time.Second)//一秒发送一次
	}
}
