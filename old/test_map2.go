package main

import (
	"fmt"
)

/*
1)使用 map[string]map[string]sting 的map类型
2)key: 表示用户名，是唯一的，不可以重复
3)如果某个用户名存在，就将其密码修改"888888"，如果不存在就增加这个用户信息,
（包括昵称nickname 和 密码pwd）。
4)编写一个函数 modifyUser(users map[string]map[string]sting, name string) 完成上述功能
*/
func main() {

	// var m map[string]map[string]string
	// m = make(map[string]map[string]string)
	m := make(map[string]map[string]string) // 这个定义只能定义最外层。

	m["ok"] = make(map[string]string)
	m["ok"]["name"] = "wt"
	m["ok"]["pwd"] = "123456"
	m["ok"]["nickname"] = "wutao"

	ModifyOk(m, "ok")
	ModifyOk(m, "wt1")
	ModifyOk(m, "wt2")

	fmt.Println(m)

	//第二种方式
	cities := make(map[string]string)
	cities["no1"] = "北京"
	cities["no2"] = "天津"
	cities["no3"] = "上海"
	fmt.Println(cities)

	//因为 no3这个key已经存在，因此下面的这句话就是修改
	cities["no3"] = "上海~"
	fmt.Println(cities)

	//演示删除
	delete(cities, "no1")
	fmt.Println(cities)
	//当delete指定的key不存在时，删除不会操作，也不会报错
	delete(cities, "no4")
	fmt.Println(cities)

	//map的声明和注意事项
	var a map[string]string
	//在使用map前，需要先make , make的作用就是给map分配数据空间
	a = make(map[string]string, 1)
	a["no1"] = "宋江" //ok?
	a["no2"] = "吴用" //ok?
	a["no1"] = "武松" //ok?
	a["no3"] = "吴用" //ok?
	// fmt.Println(cap(a))

	map1 := make(map[int]int, 2)
	map1[1] = 90
	map1[2] = 88
	map1[10] = 1
	map1[20] = 2
	modify(map1)
	// 看看结果， map1[10] = 900 ,说明map是引用类型
	fmt.Println(map1)

	mm := make(map[int]Student, 6)
	mm[1] = Student{"wutao", 20}
	fmt.Println(mm)
}

// map的value是一个结构体
type Student struct {
	Name string
	Age  int
}

func modify(m map[int]int) {
	m[2] = 1000
}

func ModifyOk(m map[string]map[string]string, name string) {
	// map取值判断用此方法。
	if _, err := m[name]; err != false {
		m[name]["pwd"] = "888888"
	} else {
		// m[name] 必须再声明map。
		m[name] = make(map[string]string)
		m[name]["name"] = name
		m[name]["nickname"] = "随机一个昵称"
		m[name]["pwd"] = "123456"
	}
	// fmt.Println(m)
}
