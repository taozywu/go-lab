package main

import (
	"fmt"
	// "strconv"
	"time"
)

// /需求，每隔1秒中打印一个数字，打印到100时就退出
//需求2: 每隔0.1秒中打印一个数字，打印到100时就退出

func main() {

	var n int
	for {
		n++
		fmt.Println(n)
		time.Sleep(time.Second) // 睡1秒
		if n >= 100 {
			break
		}
	}
}
