package main

import (
	"fmt"
	"sort"
)

func main() {
	/*
		要求：使用一个map来记录monster的信息 name 和 age, 也就是说一个
		monster对应一个map,并且妖怪的个数可以动态的增加=>map切片
	*/
	var monsters []map[string]string        // map切片
	monsters = make([]map[string]string, 4) // 使用make先复制切片
	monsters[0] = make(map[string]string)   // 每个切片里面需要再次复制map
	monsters[0]["name"] = "wt"
	monsters[0]["age"] = "18"
	m := make(map[string]string) // 定义个新的map
	m["name"] = "wutao"
	m["age"] = "20"
	monsters = append(monsters, m)
	fmt.Println(monsters)

	//map的排序，本身是乱序的，如果想排序，需要变成切片，本身切片操作的是底层的数组。
	map1 := make(map[int]int, 10)
	map1[10] = 100
	map1[1] = 13
	map1[4] = 56
	map1[8] = 90

	fmt.Println(map1)

	s := make([]int, 4)
	fmt.Println(s)

	i := 0
	for k, _ := range map1 {
		s[i] = k
		i++
	}

	// 排序
	sort.Ints(s)
	fmt.Println(s)

	//案例
	/*
		课堂练习：演示一个key-value 的value是map的案例
		比如：我们要存放3个学生信息, 每个学生有 name和sex 信息
		思路:   map[string]map[string]string

	*/

	mm := make(map[string]map[string]string)
	mm["wutao"] = make(map[string]string)
	mm["wutao"]["name"] = "wt"
	mm["wutao"]["sex"] = "1"
	fmt.Println(mm)
}
