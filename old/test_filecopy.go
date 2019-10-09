package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	dstFileName := "d:/Work/GoWork/src/go-lab/tt.txt"
	srcFileName := "d:/Work/GoWork/src/go-lab/t.txt"
	_, err := CopyFile(dstFileName, srcFileName)
	if err != nil {
		fmt.Println("拷贝失败")
	} else {
		fmt.Println("拷贝成功")
	}

}

// 自己来实现一个copy的函数
func CopyFile(dstFileName string, srcFileName string) (written int64, err error) {

	srcFile, err := os.Open(srcFileName)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
	}
	defer srcFile.Close()
	//通过srcfile ,获取到 Reader
	reader := bufio.NewReader(srcFile)

	//打开dstFileName
	dstFile, err := os.OpenFile(dstFileName, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Printf("open file err=%v\n", err)
		return
	}

	//通过dstFile, 获取到 Writer
	writer := bufio.NewWriter(dstFile)
	defer dstFile.Close()

	for {
		r, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		writer.WriteString(r)
	}

	writer.Flush()

	return
	// 文本内容貌似拷贝不成功???这个只能拷贝图片？？？？？？？？？？
	// return io.Copy(writer, reader)
}
