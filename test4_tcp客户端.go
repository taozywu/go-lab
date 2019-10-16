package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	// 客户端主动连接服务器
	conn, err := net.Dial("tcp", "127.0.0.1:9000")

	if err != nil {
		fmt.Println("net.Dial err", err)
		return
	}

	// 关闭连接
	defer conn.Close()

	// 接受键盘输入
	// 这个开一个协程
	go func() {
		// 切片缓存
		str := make([]byte, 1024)

		// 不停的接受，需要用for
		for {
			// 读取键盘输入
			n, err := os.Stdin.Read(str)
			if err != nil {
				fmt.Println("os.Stdin.Read err", err)
				return
			}

			conn.Write(str[:n]) // 将输入的信息发送给服务器
		}
	}()

	buf := make([]byte, 1024)

	// 不停的接受，需要用for
	// 不断读取服务器回复的数据
	for {
		n, err := conn.Read(buf)

		// 当输入exit后，服务器会直接return掉后；此时再读取则会出现err后终止。
		if err != nil {
			fmt.Println("conn.Read err", err)
			return
		}

		if n == 0 {
			fmt.Println("n=0")
			return
		}

		fmt.Println(string(buf[:n])) // 打印读取的内容
	}
}
