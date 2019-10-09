package main

import (
	"encoding/json"
	"fmt"
)

type A struct {
	Num int
}
type B struct {
	Num int
}

type Monster struct {
	Name  string `json:"name"` // `json:"name"` 就是 struct tag
	Age   int    `json:"age"`
	Skill string `json:"skill"`
}

func main() {
	var a A
	var b B
	b.Num = 1
	a = A(b) // ? 可以转换，但是有要求，就是结构体的的字段要完全一样(包括:名字、个数和类型！)
	fmt.Println(a, b)

	m := Monster{"wutao", 20, "OKK"}
	fmt.Println(m)

	jsonStr, err := json.Marshal(m)

	if err != nil {
		fmt.Println("err")
	} else {
		fmt.Println(jsonStr)
	}
}
