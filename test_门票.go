package main

import (
	"fmt"
)

// 景区门票案例

// 一个景区根据游人的年龄收取不同价格的门票，比如年龄大于18，收费20元，其它情况门票免费.
// 请编写Visitor结构体，根据年龄段决定能够购买的门票价格并输出

type Visitor struct {
	age int
}

func (v Visitor) output() {
	if v.age > 18 {
		fmt.Println("交20")
	} else {
		fmt.Println("免费")
	}
}

func main() {
	var v Visitor
	v.age = 21
	v.output()
}
