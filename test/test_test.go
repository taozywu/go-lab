package test

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	var a int = 1
	var b int = 1

	fmt.Print(a)
	for i := 0; i < 5; i++ {
		fmt.Print(" ", b)
		tmp := a
		a = b
		b = tmp + a
	}
}
