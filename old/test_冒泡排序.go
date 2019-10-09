package main

import (
	"fmt"
)

// 冒泡排序
// 1) 需要一个数组
// 每次先取一个数，然后循环另外几个数跟此数进行对比，比他大则互换位置。

func main() {

	var arr [5]int = [5]int{2, 1, 4, 5, 3}

	for i := 0; i < len(arr)-1; i++ {
		var flag bool = true
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				tmp := arr[j]
				arr[j] = arr[j+1]
				arr[j+1] = tmp
				flag = false
			}
		}

		if flag {
			break
		}
	}

	fmt.Println(arr)

}
