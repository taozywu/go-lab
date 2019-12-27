package main

import (
	"fmt"
)

// 定义一个Options结构体
type Options struct {
	Network string
	Address string
}

// 定义一个函数类型，函数Option，参数是Options结构体引用
type Option func(*Options)

func newOptions(opt ...Option) *Options {
	opts := Options{}
	// opt是一个函数
	for _, o := range opt {
		// 地址
		o(&opts)
	}
	// 需要返回的是：结构体指针
	return &opt
}

func Network(n string) Option {
	// 传参：指针
	return func(o *Options) {
		o.Network = n
	}
}

// Address server 监听地址
func Address(a string) Option {
	return func(o *Options) {
		o.Address = a
	}
}

func newServer(opts ...Option) {
	options := newOptions(opts...)
	fmt.Println(options)
}

func main() {
	fmt.Println("ok")
	newServer(Network("a"), Address("b"))
}
