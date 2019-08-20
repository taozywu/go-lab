package calc

import (
	"fmt"
)

func init() {
	fmt.Println("calc.init")
}

func test1(a int, b int) int {
	return a + b
}

func Test1(a int, b int) int {
	return a + b
}

func Test2(a int, b int, c byte) int {
	var result int
	switch c {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			fmt.Println("分母不能为0")
		} else {
			result = a / b
		}
	}
	return result
}
