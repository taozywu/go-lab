package handler

import (
	"fmt"
)

type Arith int

type Args struct {
	A, B int
	C    string
}

func init() {
	fmt.Println("test.init")
}

func (t *Arith) Multiply(args *Args, reply *Args) error {
	reply.A = args.A * args.B
	reply.C = args.C + "_hello"
	return nil
}
