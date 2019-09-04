package utils

import (
	"fmt"
)

type person struct {
	Name string
	age  int //其它包不能直接访问..
	sal  float64
}

//这个函数类似对象初始化构造方法
func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

// set方法
func (this *person) SetAge(age int) {
	if age > 0 {
		this.age = age
	} else {
		fmt.Println("age 不能为小于0")
	}
}

// get方法
func (this *person) GetAge() int {
	return this.age
}
