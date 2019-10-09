package main

import (
	"fmt"
	"go-lab/utils"
)

func main() {
	fmt.Println("测试面向对象构造，也可以理解是初始化")

	aa := utils.NewPerson("dddddd")
	fmt.Println(aa)

	aa.SetAge(19)
	fmt.Println(aa)

	fmt.Println(aa.GetAge())
}
