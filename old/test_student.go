package main

import (
	"fmt"
)

/*
学生案例：
编写一个Student结构体，包含name、gender、age、id、score字段，分别为string、string、int、int、float64类型。
结构体中声明一个say方法，返回string类型，方法返回信息中包含所有字段值。
在main方法中，创建Student结构体实例(变量)，并访问say方法，并将调用结果打印输出。

*/
type Student struct {
	name   string
	gender string
	age    int
	id     int
	score  float64
}

func (s Student) say() string {
	fmt.Println(s.name, s.gender, s.age, s.id, s.score)
	return "hello"
}

func main() {
	s := Student{"wutao", "nan", 20, 111, 90.2}
	fmt.Println(s.say())
}
