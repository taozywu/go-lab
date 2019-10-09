package main

import (
	"fmt"
)

// 将函数当一个参数来处理
func main() {
	sum := Test(Test1, 1, 2)
	fmt.Println(sum)
}

func Test1(a int, b int) int {
	return a + b
}

func Test(a func(int, int) int, b int, c int) int {
	return a(b, c)
}
