package main

import "fmt"

// // Go提供的结构体就是把使用各种数据类型定义的不同变量组合起来的高级数据类型
// type Person struct {
// 	Name string
// 	Age  int
// }

// func main() {
// 	fmt.Println(Person{"A", 18})            // 顺序不能错
// 	fmt.Println(Person{Name: "B", Age: 19}) // 这个按照字段来实例
// 	fmt.Println(Person{Name: "C"})          // 部分实例
// }

// 定义一个结构体  => 类似先定义一个类

// 在这个结构体上绑定一个方法  => 类似顶一个这个类的方法

// type Person struct {
// 	A int
// 	B int
// }

// func (p Person) Add() int {
// 	return p.A + p.B
// }

// func main() {
// 	// 实例下这个结构体

// 	p := Person{1, 2}
// 	fmt.Println(p.Add())
// }
func main() {
	A()
	B()
}
func A() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("@@@@@@@@@@")
			fmt.Println(err)
			fmt.Println("@@@@@@@@@@")
		}
	}()
	fmt.Println("Hi A1")
	panic("A") //recover 会跳过panic之后的所有语句
	fmt.Println("Hi A2")
}
func B() {
	fmt.Println("Hi B")
}
