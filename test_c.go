package main

import (
	"net"
)

import (
	"fmt"
)

func main() {

	// 连接tcp
	conn, err := net.Dial("tcp", "127.0.0.1:8888")

	if err != nil {
		fmt.Println("net.Dial err=", err)
		fmt.Println("连接失败")
		return
	}

	defer conn.Close()

	// 发送数据
	_, err = conn.Write([]byte("aaaaaaaaaaaaaaaaaaaacccccccccccccccccc\n"))
	if err != nil {
		fmt.Println("conn.Write err=", err)
		fmt.Println("发送数据错误")
		return
	}
	fmt.Println("发送成功")

}
