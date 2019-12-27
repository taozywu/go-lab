package main

import (
	"log"

	"github.com/ctfang/network"
)

func main() {
	client := network.NewClient("tcp://127.0.0.1:8080")
	client.SetEvent(&clientevent{})
	client.ListenAndServe()
}

type clientevent struct {
}

func (event *clientevent) OnStart(listen network.ListenTcp) {

}

func (*clientevent) OnConnect(connect network.Connect) {
}

func (*clientevent) OnMessage(connect network.Connect, message []byte) {
	log.Println(string(message))
	connect.Send("[123]")
}

func (*clientevent) OnClose(connect network.Connect) {

}

func (*clientevent) OnError(listen network.ListenTcp, err error) {

}
