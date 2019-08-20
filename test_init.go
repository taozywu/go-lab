package main

import (
	"calc"
	"fmt"
	// "os"
)

func init() {
	fmt.Println("test_init.init")
}

// 给一个a，b，在给一个+ - * /
func main() {
	calc.Test1(1, 2)
	result := calc.Test2(1, 0, '/')
	fmt.Println(result)

	FuncGold(4)
}

// 9*9乘法
//   *            1
//  ***           3
// *****          5

// 首先 1/3/5/7
// 层：4 星星：最后一层 4*2-1 			第一层 1*2-1
// 层：4 空格：最后一层 （1-1）    		第一层 （4-1）

func FuncGold(f int) {
	var a string = " "
	var b string = "*"

	//	打星星
	for i := 1; i <= f; i++ {

		// 空格，先打满
		// 打星星， 直接覆盖上去即可
		for q := 1; q <= f-i; q++ {
			fmt.Print(a)
		}

		tmp := 2*i - 1
		for j := 1; j <= tmp; j++ {
			fmt.Print(b)
		}
		fmt.Println()
	}
}
