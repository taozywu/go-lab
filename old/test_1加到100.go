package main

import (
	"fmt"
)

// 1 加到100
func main() {
	var sum int
	for a := 1; a <= 100; a++ {
		sum += a
	}
	fmt.Println(sum)
}
