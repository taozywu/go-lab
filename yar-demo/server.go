package main

import (
	"fmt"
	"go-lab/yar-demo/handler"
	"net"
	"runtime"
	"sync"

	"github.com/gyf19/yar-go/yar"
)

var worker = runtime.NumCPU()

func main() {
	runtime.GOMAXPROCS(worker)

	var server = yar.NewServer()
	// 调用系统rpc-server的注册方法：目的将需要绑定的结构体给注册进来
	server.Register(new(handler.Arith))
	server.Register(new(handler.Test))

	var wg sync.WaitGroup
	wg.Add(1)
	// 开一个协程+闭包
	go func() {
		// 新建一个listener
		listener, err := net.Listen("tcp", ":12345")
		if err != nil {
			fmt.Println(err)
			return
		}
		wg.Done()
		server.Accept(listener)
	}()
	wg.Wait()
	fmt.Scanln()

}
