package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("ceshi 文件")

	file, err := os.Open("d:/Work/GoWork/src/go-lab/t.txt")
	if err != nil {
		fmt.Println("打开错误：", err)
	} else {
		fmt.Printf("file=%v\n", file)
		err := file.Close()
		if err != nil {
			fmt.Println("关闭错误：", err)
		} else {
			fmt.Println("关闭成功")
		}
	}

}
