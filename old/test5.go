// package main

// import (
// 	"fmt"
// 	// "os"
// 	"runtime"
// )

// func main() {
// 	fmt.Println("ceshi 文件")

// 	// file, err := os.Open("d:/Work/GoWork/src/go-lab/t.txt")
// 	// if err != nil {
// 	// 	fmt.Println("打开错误：", err)
// 	// } else {
// 	// 	fmt.Printf("file=%v\n", file)
// 	// 	err := file.Close()
// 	// 	if err != nil {
// 	// 		fmt.Println("关闭错误：", err)
// 	// 	} else {
// 	// 		fmt.Println("关闭成功")
// 	// 	}
// 	// }

// 	// cpuNum := runtime.NumCPU()
// 	// fmt.Println("cpuNum=", cpuNum)

// 	// //可以自己设置使用多个cpu
// 	// runtime.GOMAXPROCS(cpuNum - 1)
// 	// fmt.Println("ok")

// }
package main

import (
	"fmt"
	"reflect"
	// "time"
	// "strconv"
	// "time"
)

// 在主线程(可以理解成进程)中，开启一个goroutine, 该协程每隔1秒输出 "hello,world"
// 在主线程中也每隔一秒输出"hello,golang", 输出10次后，退出程序
// 要求主线程和goroutine同时执行

//编写一个函数，每隔1秒输出 "hello,world"
// func test() {
// 	for i := 1; i <= 10; i++ {
// 		fmt.Println("tesst () hello,world " + strconv.Itoa(i))
// 		time.Sleep(time.Second)
// 	}
// }

func main() {

	// go test() // 开启了一个协程

	// for i := 1; i <= 10; i++ {
	// 	fmt.Println(" main() hello,golang" + strconv.Itoa(i)) // 将数字转字符串
	// 	time.Sleep(time.Second)
	// }

	// a()

	// func() {
	// 	sum := 1
	// 	for i := 1; i <= 2; i++ {
	// 		fmt.Println(i)
	// 		sum += i
	// 	}
	// 	fmt.Println(sum)
	// }()

	// fmt.Println(f())

	// var g int
	// go func(i int) {
	// 	s := 0
	// 	for j := 0; j < i; j++ {
	// 		fmt.Println(j, "@@@")
	// 		s += j
	// 	}
	// 	g = s
	// }(10)

	// time.Sleep(time.Second * 4)

	// fmt.Println(g) // 45

	// var arr [2]int
	// for i := 0; i < 2; i++ {
	// 	arr[i] = i + 1
	// }
	// fmt.Println(arr)

	// a := [...]string{"a", "b", "c", "d"}
	// for i := range a {
	// 	fmt.Println("Array item", i, "is", a[i])
	// }

	var str string = "tom"      //ok
	fs := reflect.ValueOf(&str) //ok fs -> string
	fs.Elem().SetString("jack") //ok
	fmt.Printf("%v\n", str)     // jack
}

// func a() {
// 	i := 0
// 	defer fmt.Println(i)
// 	i++
// 	return
// }

// func f() {
// 	g := func(i int) { fmt.Printf("%d ", i) } //此例子中只是为了演示匿名函数可分配不同的
// 	//内存地址，在现实开发中，不应该把该部分信息放置到循环中。
// 	for i := 0; i < 4; i++ {
// 		g(i)
// 		// fmt.Printf(" - g is of type %T and has value %v\n", g, g)
// 	}
// }

// ret = 2;
