package main

import (
	"fmt"
	"go-lab/utils"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

//使用Go提供的json-rpc 标准包
func init() {
	fmt.Println("JSON RPC 采用了JSON，而不是 gob编码，和RPC概念一模一样，")
}

func main() {
	// 注册DemoM
	str := new(utils.DemoM)
	rpc.Register(str)

	// 注册DemoM1
	str1 := new(utils.DemoM1)
	rpc.Register(str1)

	// 从上面得知是可行的。这样我们就知道在这里面
	// 是可以监听一个ip+端口方式来处理一个项目里面的多个类文件或多个模块下的方法。
	// 这个里面我们可以封装下，组装成

	/*
		array(
			"uri" => "127.0.0.1:9090",
			"count" => "2"    // 数量
			"register_list" => array("test1", "test1", "test3")
		)
	*/

	tcpAddr, err := net.ResolveTCPAddr("tcp", ":8080")
	if err != nil {
		fmt.Println("err=", err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		fmt.Println("err=", err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		jsonrpc.ServeConn(conn)

	}

}
