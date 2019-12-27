package main

import (
	// "io"
	"net"
)

import (
	"fmt"
)

func main() {
	// listen
	listen, err := net.Listen("tcp", ":8888")

	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}

	fmt.Println("开启tcp服务成功！")

	// 关闭监听
	defer listen.Close()

	for {
		// 如果客户端连接到server端了，就会创建一个新的conn
		conn, err := listen.Accept()
		fmt.Println("客户端连上server了")

		if err != nil {
			fmt.Println("listen.Accept err=", err)
			fmt.Println("连接失败")
			return
		}

		go ProcessConn(conn)

	}

}

func ProcessConn(conn net.Conn) {
	// 关闭掉客户端连上的连接
	defer conn.Close()

	// 初始化接受字节的长度限制
	buf := make([]byte, 10)
	allBytes := []byte{}

	for {
		n, err := conn.Read(buf)
		if err != nil {
			// 说明客户端关闭了
			// if err == io.EOF {
			// 	return
			// }
			fmt.Println("conn.Read err=", err)
			fmt.Println("接受数据错误")
			return
		}
		allBytes = append(allBytes, buf[:n]...)
	}

	fmt.Println("接受到数据了")
	fmt.Println(string(allBytes))
}
