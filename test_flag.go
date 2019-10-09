package main

import (
	"flag"
	"fmt"
)

var (
	confPath string
)

func main() {

	// 将conf参数注册给一个新变量
	flag.StringVar(&confPath, "conf", "", "default config path")

	// 解析
	flag.Parse()

	fmt.Println(confPath)
}
