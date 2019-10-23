package main

import (
	"log"

	"github.com/athanb/workerman"
)

func main() {
	client := workerman.NewClient("tcp://127.0.0.1:9000")
	client.SetEvent(&clientevent{})
	client.ListenAndServe()
}

type clientevent struct {
}

func (event *clientevent) OnStart(listen workerman.ListenTcp) {

}

func (*clientevent) OnConnect(connect workerman.Connect) {
}

func (*clientevent) OnMessage(connect workerman.Connect, message []byte) {
	log.Println(string(message))
	connect.Send("[123]")
}

func (*clientevent) OnClose(connect workerman.Connect) {

}

func (*clientevent) OnError(listen workerman.ListenTcp, err error) {

}
