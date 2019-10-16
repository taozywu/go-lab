package main

import (
	"fmt"
	"net"
	"strings"
)

func main() {

	// 监听tcp
	listener, err := net.Listen("tcp", "127.0.0.1:9000")

	if err != nil {
		fmt.Println("err=", err)
	}

	// 关闭
	defer listener.Close()

	// 接受多个用户
	for {
		// 接受
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("err=", err)
		}

		// 处理用户发送过来的数据
		// 每来一个开一个新协程
		go HandleConn(conn)
	}

}

// conn 是接口类型
func HandleConn(conn net.Conn) {

	// 关闭conn
	defer conn.Close()

	// 获取客户端网络地址
	addr := conn.RemoteAddr().String()
	fmt.Println("addr:", addr)

	// 指定读取的长度
	buf := make([]byte, 20)

	for {
		// 读取用户数据
		n, err := conn.Read(buf)

		if err != nil {
			fmt.Println("err=", err)
			return
		}

		// 打印长度
		fmt.Println(len(string(buf[:n])))

		// 判断exit
		if "exit" == string(buf[:n-2]) {
			fmt.Println("exit")
			return
		}

		// 读到数据了，不一定读完了，读取指定的获取的长度
		fmt.Println("read buf=", string(buf[:n]))

		// 将数据再发送给客户端
		conn.Write([]byte(strings.ToUpper(string(buf[:n]))))
	}

}
