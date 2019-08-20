package main

import (
	"strconv"
	"strings"

	// "strings"

	// "calc"
	"fmt"
	// "os"
)

func main() {

	// 字母和数字占一个字节、中文因为是utf8，所以一个中文占3个字节
	str1 := "123"        // len 9
	str2 := []rune(str1) // len 5

	str3 := make([]string, 5)

	for i := 0; i < len(str2); i++ {
		// fmt.Println(string(str2[i]))

		// if string
		str3[i] = string(str2[i])

		// fmt.Println(i, str3)
	}

	fmt.Println(str3, len(str3), cap(str3))

	// 字符串转整数
	str4, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(str4)
	}

	// 整数转字符串
	str6 := 123123
	str5 := strconv.Itoa(str6)
	fmt.Printf("%T %v", str5, str5)

	fmt.Println("")
	var bytes = []byte("中hello go")
	fmt.Printf("bytes=%v\n", bytes)

	fmt.Println(string(bytes))

	b := strings.Contains("seafood", "mary")
	fmt.Printf("b=%v\n", b)

	num := strings.Count("ceheEse", "e")
	fmt.Printf("num=%v\n", num)
}
