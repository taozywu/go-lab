package handler

import (
	"fmt"
)

type Test int

type Args1 struct {
	A, B int
	C    string
}

func init() {
	fmt.Println("test1.init")
}

func (t *Test) Demo(args *Args1, reply *Args1) error {
	reply.A = 1
	reply.C = "test"
	return nil
}
