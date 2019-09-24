package main

import (
	"fmt"
	"time"
)

func main() {

	inChan := make(chan int, 10)
	exitChan := make(chan int, 1)

	go WriteData(inChan)
	go ReadData(inChan, exitChan)

	for {
		i, ok := <-exitChan
		if !ok {
			fmt.Println("main exitChan break")
			break
		}
		fmt.Println("exitChan=", i)
	}
}

func WriteData(inChan chan int) {
	for i := 0; i < 10; i++ {
		inChan <- i
		fmt.Println("WriteData=", i)
	}
	fmt.Println("WriteData end")
	close(inChan)
}

func ReadData(inChan chan int, exitChan chan int) {

	for {
		i, ok := <-inChan
		if !ok {
			fmt.Println("ReadData break")
			break
		}
		time.Sleep(time.Second)
		fmt.Println("ReadData=", i)
	}
	fmt.Println("ReadData end")
	exitChan <- 2
	close(exitChan)
}
