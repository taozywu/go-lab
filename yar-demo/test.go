package main

import (
	// "flag"
	"fmt"

	// "net/rpc"
	// "runtime"

	// "sync"
	"time"

	"github.com/gyf19/yar-go/yar"
)

type Args struct {
	A, B int
	C    string
}

func init() {
}

func main() {

	t1 := time.Now()
	var client, err = yar.Dial("tcp", "127.0.0.1:12345", "json") //msgpack json
	if err != nil {
		fmt.Println("Dialing:", err)
	}
	var args = Args{4, 5, "GO"}
	var reply = Args{}
	for i := 0; i < 1000; i++ {

		err := client.Call("Arith.Multiply", args, &reply)
		if err != nil {
			fmt.Println("Call:", err)
		}

		fmt.Printf("Result: %d * %d = %d\n", args.A, args.B, reply)
	}
	t2 := time.Now()
	fmt.Println(" lasted ", t2.Sub(t1))

	var result = Args{}
	err = client.Call("Test.Demo", args, &result)
	if err != nil {
		fmt.Println("Call:", err)
	}
	fmt.Println(result)
}
