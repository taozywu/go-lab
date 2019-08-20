package main

import (
	"fmt"
)

/*
	一个养鸡场有6只鸡，它们的体重分别是3kg,5kg,1kg,
	3.4kg,2kg,50kg 。请问这六只鸡的总体重是多少?平
	均体重是多少? 请你编一个程序。=》数组

*/
func main() {
	// 总体重 循环加起来

	// 得出总体重 / 6

	// var c = [6]float64{}
	// c[0] = 3.0
	// c[1] = 5.0
	// c[2] = 1.0
	// c[3] = 3.4
	// c[4] = 2.0
	// c[5] = 50.0
	// fmt.Println(c)

	// var sum float64
	// for _, v := range c {
	// 	sum += v
	// }
	// fmt.Println(sum)

	// fmt.Println(int(sum / 6))

	// // var intArr [3]int //int占8个字节
	// // //当我们定义完数组后，其实数组的各个元素有默认值 0
	// // fmt.Println(intArr)
	// // intArr[0] = 10
	// // intArr[1] = 20
	// // intArr[2] = 30
	// // fmt.Println(intArr)
	// // fmt.Printf("intArr的地址=%p intArr[0] 地址%p intArr[1] 地址%p intArr[2] 地址%p\n",
	// // 	&intArr, &intArr[0], &intArr[1], &intArr[2])

	// // 找最大值
	// var intArr [6]int = [...]int{1, -1, 9, 90, 11, 9000}
	// var maxNo int
	// var maxNoIndex int

	// for i := 0; i < 6; i++ {
	// 	if maxNo < intArr[i] {
	// 		maxNo = intArr[i]
	// 		maxNoIndex = i
	// 	}
	// }

	// fmt.Println(maxNo, maxNoIndex)

	// //使用常规的for循环遍历切片
	// var arr [5]int = [...]int{10, 20, 30, 40, 50} // 数组
	// //slice := arr[1:4] // 20, 30, 40
	// slice := arr[1:4]
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Printf("slice[%v]=%v ", i, slice[i])
	// }

	// fmt.Println()

	// //使用for--range 方式遍历切片
	// for i, v := range slice {
	// 	fmt.Printf("i=%v v=%v \n", i, v)
	// }

	// slice2 := slice[1:2] //  slice [ 20, 30, 40]    [30]
	// slice2[0] = 100      // 因为arr , slice 和slice2 指向的数据空间是同一个，因此slice2[0]=100，其它的都变化

	// fmt.Println("slice2=", slice2)

	// fmt.Println("slice=", slice)
	// fmt.Println("arr=", arr)

	//演示切片的使用 make
	// var slice []float64 = make([]float64, 5, 10)
	// slice[1] = 10
	// slice[3] = 20
	// //对于切片，必须make使用.
	// fmt.Println(slice)
	// fmt.Println("slice的size=", len(slice))
	// fmt.Println("slice的cap=", cap(slice))

	//string底层是一个byte数组，因此string也可以进行切片处理
	//string修改：需要先转[]byte()；如果有中文需要使用[]rune()
	str := "hello@atguigu"
	//使用切片获取到 atguigu
	slice := str[6:]
	fmt.Println("slice=", slice)

	arr1 := []byte(str)
	arr1[0] = 'z'
	str = string(arr1)
	fmt.Println("str=", str)

}
