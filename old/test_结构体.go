package main

import (
	"fmt"
)

type Person struct {
	Name string
}

// 给Person增加一个方法
func (p Person) test() {
	p.Name = "wt"
	fmt.Println(p)
}

func (p *Person) test1() {
	(*p).Name = "wt1"
	fmt.Println(*p)
}

func main() {
	var p Person
	p.Name = "wutao"
	p.test()
	fmt.Println(p)

	p.test1()
	fmt.Println(p)
}
