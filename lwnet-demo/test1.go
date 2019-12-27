package main

import (
	"fmt"
	"net"
	"sync"
)

func test() {
	conn, err := net.Dial("tcp", "127.0.0.1:6000")
	if err != nil {
		fmt.Println("客户端建立连接失败", err)
		return
	}
	defer conn.Close()
	// data := make([]byte, 5)
	conn.Write([]byte("hello"))
}

func main() {
	wg := sync.WaitGroup{}
	go func() {
		wg.Add(1)
		test()
		wg.Done()
	}()
	wg.Wait()
}
