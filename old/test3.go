package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	// 此类型可以理解是万能类型
	var a interface{} // 声明一个a 接口类型
	b := Point{1, 2}
	a = b // 这样赋值后，将a的类型固定为了Point结构类型

	var c Point
	c = a.(Point)

	fmt.Println(a, b, c)
	fmt.Printf("%T  %T\n", a, c)
}
