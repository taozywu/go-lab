package main

import (
	"fmt"
)

type Calcuator struct {
	num1 int
	num2 int
}

func (cal Calcuator) GetRes(c string) {

	var s int

	switch c {
	case "+":
		s = cal.num1 + cal.num2
	case "-":
		s = cal.num1 - cal.num2
	case "*":
		s = cal.num1 * cal.num2
	case "/":
		s = cal.num1 / cal.num2
	}
	fmt.Println(s)
}

func main() {
	// 加减乘除 1）需要两个数
	var cal Calcuator
	cal.num1 = 2
	cal.num2 = 2
	cal.GetRes("+")
	cal.GetRes("-")
	cal.GetRes("*")
	cal.GetRes("/")
	// fmt.Println(cal.GetRes("+"))
	// fmt.Println(cal.GetRes("-"))
	// fmt.Println(cal.GetRes("*"))
	// fmt.Println(cal.GetRes("/"))
}
