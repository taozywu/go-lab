package main

import (
	"fmt"
	"log"

	// "log"
	"github.com/athanb/workerman"
)

func main() {
	fmt.Println("test6")
	server := workerman.NewServer("tcp://127.0.0.1:9000")
	server.SetEvent(&tcpserverevent{})
	server.ListenAndServe()
}

type tcpserverevent struct {
}

func (*tcpserverevent) OnStart(listen workerman.ListenTcp) {

}

func (*tcpserverevent) OnConnect(connect workerman.Connect) {
	connect.SendString("OnConnect")
}

func (*tcpserverevent) OnMessage(connect workerman.Connect, message []byte) {
	log.Println(string(message))
	connect.SendString("OnMessage")
}

func (*tcpserverevent) OnClose(connect workerman.Connect) {
	log.Println("OnClose")
}

func (*tcpserverevent) OnError(listen workerman.ListenTcp, err error) {
	log.Println("OnError")
}
