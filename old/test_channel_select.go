package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int, 10)
	strCh := make(chan string, 5)

	for i := 0; i < 10; i++ {
		intCh <- i
		// strCh <- "hello"
	}

	for i := 0; i < 5; i++ {
		strCh <- "hello"
	}

	close(intCh)
	close(strCh)

	for {
		select {
		case v, ok := <-intCh:
			if !ok {
				break
			}
			time.Sleep(time.Second)
			fmt.Println("intCh", v)
		case v, ok := <-strCh:
			if !ok {
				break
			}
			time.Sleep(time.Second)
			fmt.Println("strCh", v)
		}
	}

}
