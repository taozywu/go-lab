package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"

	"github.com/lonerwolf/lwnet-go"
	ringbuffer "github.com/lonerwolf/lwnet-go/ring-buffer"
)

type TestServer struct {
}

func (svr *TestServer) OnConnect(c lwnet.Connection) {
	fmt.Println("OnConnect:" + c.RemoteAddr().String())
	return
}

func (svr *TestServer) OnMessage(c lwnet.Connection, ptBuff *ringbuffer.RingBuffer) {
	fmt.Println("OnMessage")
	tmp_buf := make([]byte, ptBuff.Length())
	ptBuff.Read(tmp_buf)
	fmt.Println("buff=", string(tmp_buf))
	c.Write(tmp_buf)
	return
}

func (svr *TestServer) OnClose(c lwnet.Connection) {
	fmt.Println("OnClose")
	return
}

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU())

	go func() {
		http.ListenAndServe("0.0.0.0:8080", nil)
	}()

	log.Fatal(lwnet.Serve(&TestServer{}, fmt.Sprintf("tcp://:%d", 6000),
		lwnet.WithAcceptCoroutineNum(1),
		lwnet.WithEventCoroutineNum(4),
		lwnet.WithReusePort(false)))
}
