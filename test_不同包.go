package main

import (
	"calc"
	"fmt"
)

func main() {
	var a, b int
	a = 10
	b = 20
	sum := calc.Test1(a, b)
	fmt.Println("sum", sum)
}
