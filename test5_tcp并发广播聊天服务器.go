package main

import (
	"fmt"
	"net"
	"time"
)

type Client struct {
	C    chan string // 用于发送数据的管道
	Name string      // 用户名
	Addr string      // 地址
}

// 保存在线用户 cliAddr ===> Client
var onlineMap map[string]Client

var message = make(chan string)

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C {
		conn.Write([]byte(msg + "\n")) // 给conn发送消息
	}
}

// 处理用户连接
func HandleConn(conn net.Conn) {
	// 用完就关闭
	defer conn.Close()

	// 获取网络地址
	cliAddr := conn.RemoteAddr().String()

	// 创建一个结构体
	cli := Client{make(chan string), cliAddr, cliAddr}

	// 将结构体加入到map中：记录在线用户
	onlineMap[cliAddr] = cli

	// 在开个一个协程，专门给当前客户端发送消息
	go WriteMsgToClient(cli, conn)

	// 广播当前客户端在线
	message <- cliAddr + ", login"

	isQuit := make(chan bool) // 是否退出
	isData := make(chan bool) // 是否有数据交互

	// 新建一个协程，接受用户发送过来的消息
	go func() {
		buf := make([]byte, 2048)

		for {
			n, err := conn.Read(buf)
			if err != nil { // 对方断开 或 出问题
				isQuit <- true
				fmt.Println("conn.Read err ", err)
				return
			}
			msg := string(buf[:n-1]) // 获取数据
			// 转发
			message <- cli.Addr + ", " + msg

			// 说明有数据
			isData <- true
		}
	}()

	// 不让结束
	for {
		select {
		case <-isQuit: // 说明退出
			delete(onlineMap, cliAddr)
			message <- cliAddr + "， logout" // 告诉退出了
			return
		case <-isData: // 说明有数据

		case <-time.After(30 * time.Second): // 超过30s就处理
			delete(onlineMap, cliAddr)
			message <- cliAddr + "， logout" // 告诉退出了
			return
		}
	}
}

//在开一个协程：转发消息，只要消息来了，遍历map，给map每个成员发送消息
func Manager() {
	// 分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <-message // 接受消息

		// 遍历在线用户
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func main() {
	// 监听
	listener, err := net.Listen("tcp", ":9000")

	if err != nil {
		fmt.Println("net.Listen err ", err)
		return
	}

	defer listener.Close()

	// 在开一个协程：转发消息，只要消息来了，遍历map，给map每个成员发送消息
	go Manager()

	// 循环等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listener.Accept err ", err)
			continue // 失败了不让终止
		}

		go HandleConn(conn) // 处理用户连接
	}
}
