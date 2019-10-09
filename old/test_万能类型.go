package main

import (
	"fmt"
)

func Test(str ...interface{}) {

	for _, v := range str {
		switch v.(type) {
		case bool:
			fmt.Println("是一个bool类型")
		case string:
			fmt.Println("是一个string类型")
		}
	}
}

func main() {
	Test(true)
}
