package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println(os.Args)
}

// // 声明一个接口
// type Usb interface {
// 	Say()
// }

// // 声明一个结构体，类似一个对象及成员
// type Student struct {
// }

// // 对Student这个结构体绑定一个Say方法， 且用的是指针类型
// func (this *Student) Say() {
// 	fmt.Println("Say()")
// }

// func main() {
// 	var u Usb
// 	// 普通类型并未实现Say方法，所以这个写法是不对的！
// 	// s := Student{}
// 	s := &Student{}
// 	u = s
// 	u.Say()
// }
