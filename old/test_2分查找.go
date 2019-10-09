package main

import (
	"fmt"
)

func main() {
	// 二分查找

	// 1） 得有一个数组、2）需要知道你要查找的数
	// 找到左右-中间那个索引对应的值，进行判断
	// 小： 说明 left，mid-1
	// 大： 说明 mid+1,right
	var arr [5]int = [5]int{1, 3, 4, 6, 2}
	BinarySearch(&arr, 0, len(arr)-1, 12)
}

func BinarySearch(arr *[5]int, left int, right int, findVal int) {

	if left > right {
		fmt.Println("找不到")
		return
	}

	mid := (left + right) / 2

	if (*arr)[mid] > findVal {
		BinarySearch(arr, left, mid-1, findVal)
	} else if (*arr)[mid] < findVal {
		BinarySearch(arr, mid+1, right, findVal)
	} else {
		fmt.Println("找到了", mid)
	}

}
