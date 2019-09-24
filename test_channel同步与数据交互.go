package main

import (
	"fmt"
	// "time"
)

func main() {

	ch := make(chan string)

	defer fmt.Println("OKKK2")

	go func() {
		defer fmt.Println("子调用完毕")

		for i := 0; i < 2; i++ {
			fmt.Println("子", i)
			// time.Sleep(time.Second)
		}

		ch <- "我是子，工作即将完毕"
	}() //闭包调用

	// 1）使用下面的来让 go func能正常操作
	// for {

	// }

	// 2）也可以使用channel来实现。
	str := <-ch // 没有数据前，阻塞
	fmt.Println(str)

	fmt.Println("OKK1")

}
