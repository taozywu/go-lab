package main

import (
	"fmt"
)

// "fmt"

func main() {

	// 后面的10 有跟没有不一样
	// 有了，说明可以缓存进去，10代表10个容量。
	ch := make(chan int, 10)

	// 往ch写一个1
	ch <- 1

	// 关闭，如果缓存中的数据还有，则可以继续读，直到结束，
	// 结束后返回的类型跟初始化定义类型一样。
	// 不能两次close。会提示close一个已close的channel
	close(ch)

	// 读取一个ch
	i, ok := <-ch
	if !ok {
		fmt.Println("end")
	}
	fmt.Println(i)
}
