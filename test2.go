package main

import (
	"fmt"
)

// 1) 查询余额 2）取款 3）存款

type Account struct {
	Name  string  // 名称
	Pwd   string  // 密码
	Money float64 // 余额
}

func (a *Account) Cha(pwd string) {
	if pwd != a.Pwd {
		fmt.Println("密码不对")
		return
	}

	fmt.Println(a.Name, a.Money)
}

func (a *Account) Qu(money float64, pwd string) {
	if pwd != a.Pwd {
		fmt.Println("密码不对")
		return
	}
	if money > a.Money {
		fmt.Println("余额不足")
		return
	}

	a.Money -= money
	fmt.Println(a.Name, a.Money)
}

func (a *Account) Cun(money float64, pwd string) {
	if pwd != a.Pwd {
		fmt.Println("密码不对")
		return
	}
	if money < 1 {
		fmt.Println("存钱不能小于1")
		return
	}
	a.Money += money
	fmt.Println(a.Name, a.Money)
}

func main() {
	a := Account{"wt", "111111", 123.1}
	fmt.Println(a)
	a.Cha("1233")
	a.Cha("111111")
	a.Qu(100, "111111")
	a.Cun(1000, "111111")
}
