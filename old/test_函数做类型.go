package main

import (
	"fmt"
)

// 函数当成一个类型
func getSum(n1 int, n2 int) int {
	return n1 + n2
}

// 将函数当成一个形参
func myFun(funvar func(int, int) int, num1 int, num2 int) int {
	return funvar(num1, num2)
}

// 如上可以简化处理
type myFunType func(int, int) int

func myFun2(funvar myFunType, num1 int, num2 int) int {
	return funvar(num1, num2)
}

func main() {

	a, b := 1, 2

	// 将函数getSum作为实参带进去
	sum := myFun(getSum, a, b)
	fmt.Println(sum)

	sum1 := myFun2(getSum, a, b)

	fmt.Println(sum1)
}
