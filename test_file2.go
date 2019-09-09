package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// 这种一次读完。对小文件可以，对大文件不行。
	// 因为不涉及到open，所以也不会有close
	r, err := ioutil.ReadFile("d:/Work/GoWork/src/go-lab/t.txt")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(r))

	// 将文件一次性写入
	err = ioutil.WriteFile("d:/Work/GoWork/src/go-lab/tt.txt", r, 0666)
	if err != nil {
		fmt.Printf("write file error=%v\n", err)
	}
}
