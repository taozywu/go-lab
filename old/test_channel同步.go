package main

import (
	"fmt"
	"time"
)

func main() {
	go person1()
	go person2()

	for {

	}
}

// 创建全局channel
var ch = make(chan int)

func person1() {
	Printer("hello")
	ch <- 666 // 给管道写数据
}

func person2() {
	<-ch // 从管道接受，如果没数据就阻塞
	Printer("world")
}

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}
