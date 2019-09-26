package main

import (
	"fmt"
	"time"
)

func main() {
	// //演示一下管道的使用
	// //1. 创建一个可以存放3个int类型的管道
	// var intChan chan int
	// intChan = make(chan int, 3)

	// //2. 看看intChan是什么
	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// //3. 向管道写入数据
	// intChan <- 10
	// fmt.Println(intChan)
	// num := 211
	// intChan <- num
	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	// intChan <- 50
	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	// // //如果从channel取出数据后，可以继续放入
	// <-intChan
	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)
	// intChan <- 98 //注意点, 当我们给管写入数据时，不能超过其容量
	// fmt.Printf("intChan 的值=%v intChan本身的地址=%p\n", intChan, &intChan)

	// //4. 看看管道的长度和cap(容量)
	// fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 3, 3

	// var num2 int
	// num2 = <-intChan
	// fmt.Println("num2=", num2)
	// fmt.Printf("channel len= %v cap=%v \n", len(intChan), cap(intChan)) // 2, 3

	// num3 := <-intChan
	// num4 := <-intChan

	// //num5 := <-intChan

	// fmt.Println("num3=", num3, "num4=", num4 /*, "num5=", num5*/)

	// 默认管道是双向， 可写可读

	// var ch chan int
	// ch = make(chan int, 3)
	// ch <- 1
	// m := <-ch
	// fmt.Println(m)

	// // var ch1 chan<- int
	// ch1 := make(chan<- int, 3)
	// ch1 <- 2
	// ch1 <- 1
	// // ch1 <- 3
	// // ch1 <- 4
	// // n := <-ch1
	// fmt.Println(len(ch1), cap(ch1))

	go sayHello()
	go test()

	for i := 0; i < 10; i++ {
		fmt.Println("main() ok=", i)
		time.Sleep(time.Second)
	}
}

//函数
func sayHello() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("hello,world")
	}
}

//函数
func test() {
	//这里我们可以使用defer + recover
	defer func() {
		//捕获test抛出的panic
		if err := recover(); err != nil {
			fmt.Println("test() 发生错误", err)
		}
	}()
	//定义了一个map
	var myMap map[int]string
	// myMap = make(map[int]string)
	myMap[0] = "golang" //error
	fmt.Println(myMap)
}
