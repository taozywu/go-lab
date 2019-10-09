package main

import (
	"calc"
	"fmt"
)

func main() {

	fmt.Println("ddd")

	var m = calc.Test11(1, "ww", 30)
	// fmt.Println(*m)

	fmt.Println((*m).Test22(), (*m).Age)
}
