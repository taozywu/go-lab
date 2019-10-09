package main

import (
	"fmt"
	// "sync"
	// "time"
)

var myMap = make(map[int]int, 10)

// var lock sync.Mutex // 互斥

func main() {
	// 需求：现在要计算 1-200 的各个数的阶乘，并且把各个数的阶乘放入到map中。
	// 最后显示出来。要求使用goroutine完成

	// 先搞懂什么是阶乘: 简单理解；n的阶乘=1*2*3...*n；同时0的阶乘是1
	go Jiecheng(3)

	// time.Sleep(time.Second * 3)

	for i, v := range myMap {
		fmt.Println(i, "-", v)
	}
}

func Jiecheng(n int) {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	myMap[n] = res
}
