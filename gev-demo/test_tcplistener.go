package main

import (
	"fmt"
	"net"

	"syscall"
	// "golang.org/x/sys/unix"
)

func main() {
	var listener net.Listener
	// 新建监听
	listener, err := net.Listen("tcp", ":3000")
	if err != nil {
		fmt.Println("net.Listen err=", err)
		return
	}
	fmt.Println(listener)
	// 判断是否是tcp监听
	l, ok := listener.(*net.TCPListener)
	if !ok {
		fmt.Println("listener is not tcp listener")
		return
	}
	// 获取socket描述符
	file, err := l.File()
	if err != nil {
		fmt.Println("could not get file descriptor")
		return
	}
	// fd := int(file.Fd())
	// 设置非阻塞
	// if err = unix.SetNonblock(fd, true); err != nil {
	// 	return
	// }
	// nfd, sa, err := syscall.Accept(file.Fd())
	if err = syscall.SetNonblock(file.Fd(), true); err != nil { //设置非阻塞模式
		fmt.Println("syscall.SetNonblock err=", err)
		return
	}
}
