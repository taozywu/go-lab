package main

import (
	"fmt"
)

func main() {

	inCh := make(chan int, 500)
	chaCh := make(chan int, 1000)
	exitCh := make(chan bool, 6)

	go InsertData(inCh)

	// 开多个协程
	for i := 0; i < 4; i++ {
		go ChaData(inCh, chaCh, exitCh)
	}

	go func() {
		// 确保每个ChaData协程已处理完
		for i := 0; i < 4; i++ {
			<-exitCh
		}
		// 关闭
		close(chaCh)
	}()

	// 读取
	for {

		// 获取
		n, ok := <-chaCh
		if !ok {
			break
		}
		fmt.Println(n)

	}
}

// 往管道插入1000条数据
func InsertData(inCh chan int) {
	for i := 1; i <= 1000; i++ {
		inCh <- i
	}
	close(inCh)
}

// 从1000条数据中找出所有质数
func ChaData(inCh chan int, chaCh chan int, exitCh chan bool) {

	// 质数：除1和他本身之外的因数
	// 因数：a/b 被整除后，b就是a的因数

	for {

		num, ok := <-inCh

		// 说明管道一个数也没有了，则后面的逻辑就不用在处理了。
		if !ok {
			break
		}

		check := true

		for i := 2; i < num; i++ {
			if num%i == 0 { // 说明不是质数
				check = false
				break
			}
		}

		// 排除掉1
		if num != 1 && check {
			chaCh <- num
		}
	}

	// 这个地方不能使用close，因为在main协程中使用了多协程了。
	// 此方法会执行多次。不能因为执行一次后就给关闭掉了。
	// close(chaCh)

	// 需要再引入一个channel
	exitCh <- true
}
