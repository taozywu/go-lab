// package main

// import (
// 	"fmt"
// 	// "os"
// 	"runtime"
// )

// func main() {
// 	fmt.Println("ceshi 文件")

// 	// file, err := os.Open("d:/Work/GoWork/src/go-lab/t.txt")
// 	// if err != nil {
// 	// 	fmt.Println("打开错误：", err)
// 	// } else {
// 	// 	fmt.Printf("file=%v\n", file)
// 	// 	err := file.Close()
// 	// 	if err != nil {
// 	// 		fmt.Println("关闭错误：", err)
// 	// 	} else {
// 	// 		fmt.Println("关闭成功")
// 	// 	}
// 	// }

// 	// cpuNum := runtime.NumCPU()
// 	// fmt.Println("cpuNum=", cpuNum)

// 	// //可以自己设置使用多个cpu
// 	// runtime.GOMAXPROCS(cpuNum - 1)
// 	// fmt.Println("ok")

// }
package main

import (
	"fmt"
	"strconv"
	"time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出 "hello,world"
func test() {
	for i := 1; i <= 10; i++ {
		fmt.Println("tesst () hello,world " + strconv.Itoa(i))
		time.Sleep(time.Second)
	}
}

func main() {

	go test() // 开启了一个协程

	for i := 1; i <= 10; i++ {
		fmt.Println(" main() hello,golang" + strconv.Itoa(i)) // 将数字转字符串
		time.Sleep(time.Second)
	}
}
