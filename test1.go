package main

import (
	"fmt"
	"time"
	// "time"
)

// func main() {
// 	ch := make(chan int)

// 	go func() {
// 		for i := 0; i < 5; i++ {
// 			ch <- i // 网通道写数据
// 		}
// 		// 写完及时关闭，关闭后无法在写入了。
// 		close(ch)
// 	}()

// 	// 读取channel数据
// 	// for {
// 	// 	if num, ok := <-ch; ok == true {
// 	// 		fmt.Println(num)
// 	// 	} else {
// 	// 		break
// 	// 	}
// 	// }

// 	for num := range ch {
// 		fmt.Println(num)
// 	}

// }
// func main01() {
// 	//设置一个定时器，时间为2s，到了时间后会触发且只触发一次。
// 	// timer := time.NewTimer(2 * time.Second)
// 	// t := <-timer.C //读取timer.C管道

// 	// time.Sleep(2 * time.Second)
// 	<-time.After(2 * time.Second)
// 	fmt.Println("时间到")
// }

// 斐波那此数列
// func main() {

// 	ch := make(chan int)
// 	quit := make(chan bool)

// 	// 消费者
// 	go func() {
// 		for i := 0; i < 8; i++ {
// 			num := <-ch
// 			fmt.Println(num)
// 		}
// 		quit <- true
// 	}()

// 	// 产生者
// 	Chan(ch, quit)
// }

// func Chan(ch chan<- int, quit <-chan bool) {

// 	x, y := 1, 1

// 	for {
// 		select {
// 		case ch <- x:
// 			x, y = y, x+y
// 		case flag := <-quit:
// 			fmt.Println(flag)
// 			return
// 		}
// 	}
// }

func main() {

	ch := make(chan int)
	quit := make(chan bool)

	go func() {
		for {
			select {
			case num := <-ch:
				fmt.Println(num)
			case <-time.After(3 * time.Second):
				fmt.Println("超市")
				quit <- true
			}
		}
	}()

	for i := 0; i < 5; i++ {
		ch <- i
		time.Sleep(time.Second)
	}

	<-quit
	fmt.Println("over")
}
