package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {

	file, err := os.Open("d:/Work/GoWork/src/go-lab/t.txt")
	if err != nil {
		fmt.Println("错误", err)
	}

	// 函数退出了就关闭它
	defer file.Close()

	reader := bufio.NewReader(file)

	for {
		r, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf(r)
	}
}
