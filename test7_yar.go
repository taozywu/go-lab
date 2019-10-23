package main

import (
	"flag"
	"fmt"
	"net/rpc"
	"runtime"
	"sync"
	"time"

	"github.com/gyf19/yar-go/yar"
)

type Arith int

type Args struct {
	A, B int
	C    string
}

var worker = runtime.NumCPU()

var (
	Num  int
	Conn int
	Addr string
)

const RS = 1

func init() {
	flag.IntVar(&Num, "num", 10, "")
	flag.IntVar(&Conn, "conn", 1, "")
	flag.StringVar(&Addr, "addr", "127.0.0.1:12345", "")
	flag.Parse()
}

func main() {
	runtime.GOMAXPROCS(worker)
	//f, _ := os.Create("profile_file")
	//pprof.StartCPUProfile(f)     // 开始cpu profile，结果写到文件f中
	//defer pprof.StopCPUProfile() // 结束profile

	clients := make([]*rpc.Client, 0, 10)
	for i := 0; i < Conn; i++ {
		var client, err = yar.Dial("tcp", Addr, "json") //msgpack json
		if err != nil {
			fmt.Println(err)
		}
		clients = append(clients, client)
	}
	index := 0
	var wg sync.WaitGroup
	start := time.Now()
	for i := 1; i <= Num; i++ {
		wg.Add(1)
		var c *rpc.Client
		if index < len(clients) {
			c = clients[index]
			index++
		} else {
			index = 0
			c = clients[index]
		}

		go func(cli *rpc.Client, i int) {
			defer wg.Done()
			var args = &Args{4, 5, "GO"}
			var reply = &Args{}
			fmt.Println("goroutine start...")
			st := time.Now()
			for n := 0; n < RS; n++ {
				if err := cli.Call("Arith.Multiply", args, reply); err != nil {
					fmt.Println(i, n, err)
				}
				fmt.Println(i, reply)
			}
			fmt.Println(time.Now().Sub(st))
		}(c, i)
	}
	wg.Wait()

	total := RS * Num
	timeSub := (int)(time.Now().Sub(start) / 1000000)

	fmt.Printf("concurrency: %d\n", Num)
	fmt.Printf("total: %d\n", total)
	fmt.Printf("seconds: %d\n", timeSub)
	fmt.Printf("qps: %d\n", total/timeSub*1000)
}

/*
func testClient(wg *sync.WaitGroup) {
	t1 := time.Now()
	var client, err = yar.Dial("tcp", "192.168.125.185:12345", "json") //msgpack json
	if err != nil {
		fmt.Println("Dialing:", err)
	}
	for i := 0; i < 1000; i++ {
		var args = &Args{4, 5, "GO"}
		var reply = &Args{}

		err := client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			fmt.Println("Call:", err)
		}

		//fmt.Printf("Result: %d * %d = %d", args.A, args.B, reply)
	}
	t2 := time.Now()
	fmt.Println(" lasted ", t2.Sub(t1))
	wg.Done()
} */
