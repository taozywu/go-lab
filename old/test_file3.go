package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	filePath := "d:/Work/GoWork/src/go-lab/t.txt"
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println(err)
	}

	// 退出时记得关闭file
	defer file.Close()

	// 连续写入hellworld
	writer := bufio.NewWriter(file)
	for i := 0; i < 5; i++ {
		writer.WriteString("hellworld223\n")
	}
	//因为writer是带缓存，因此在调用WriterString方法时，其实
	//内容是先写入到缓存的,所以需要调用Flush方法，将缓冲的数据
	//真正写入到文件中， 否则文件中会没有数据!!!
	writer.Flush()
}
