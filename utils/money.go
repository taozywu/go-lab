package utils

import (
	"fmt"
)

type MyFamilyAccount struct {

	// 定义一个字段
	loop bool
	// 用于接收用户的输入
	key string
	// 记录用户的收入和支出情况，该字符串会拼接
	details string
	// 保存账号的金额
	balance float64
	// 定义一个标识符
	flag bool
	// 定义一个金额
	money float64
	// 声明一个说明
	note string
}

func NewMyFamilyAccount() *MyFamilyAccount {
	return &MyFamilyAccount{
		loop:    true,
		key:     "",
		details: "nihao",
		balance: 122.122,
		flag:    true,
		money:   2000,
		note:    "",
	}
}

func (this *MyFamilyAccount) MainAccount() {

	for {
		// 1. 先输出这个主菜单
		fmt.Println("-----------------家庭收支记账软件-----------------")
		fmt.Println("			1 收支明细")
		fmt.Println("			2 登记收入")
		fmt.Println("			3 登记支出")
		fmt.Println("			4 退出")
		fmt.Print("请选择(1-4):")
		fmt.Scanln(&this.key)

		switch this.key {
		case "1":
			this.show()
		case "2":
			this.income()
		case "3":
			this.outcome()
		case "4":
			this.loop = this.exit()
		}

		if !this.loop {
			break
		}

	}
}

func (this *MyFamilyAccount) show() {
	fmt.Println(this.details)
}

func (this *MyFamilyAccount) income() {
	fmt.Print("本次输入金额:")
	fmt.Scanln(&this.money)

	this.balance += this.money
	fmt.Println("总和：", this.balance)
}

func (this *MyFamilyAccount) outcome() {
	fmt.Print("本次支出金额:")
	fmt.Scanln(&this.money)

	if this.balance < this.money {
		fmt.Println("钱不够呀")
		return
	} else {
		this.balance -= this.money
	}
	fmt.Println("还剩下：", this.balance)
}

func (this *MyFamilyAccount) exit() bool {
	// 修改loop
	fmt.Println("确定要退出吗？y/n")

	choice := ""
	//循环判断，直到输入y 或者 n

	for {
		fmt.Scanln(&choice)
		if choice == "y" || choice == "n" {
			break
		}
		fmt.Print("你的输入有误，请输入y/n:")
	}
	if choice == "y" {
		return false
	} else {
		return true
	}
}
