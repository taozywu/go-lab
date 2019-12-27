package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

func main() {
	// 客户端连接redis
	conn, err := redis.Dial("tcp", "192.168.1.235:7008")

	if err != nil {
		fmt.Println("redis.Dial err=", err)
		return
	}
	// 关闭redis
	defer conn.Close()

	// _, err1 := conn.Do("set", "name3", "w1t")

	// if err1 != nil {
	// 	fmt.Println("Set err=", err1)
	// 	return
	// }

	// r, err2 := redis.String(conn.Do("get", "name3"))

	// if err1 != nil {
	// 	fmt.Println("Get err=", err2)
	// 	return
	// }

	// fmt.Println(r)

	// 进行hash操作
	// _, err1 := conn.Do("hSet", "user", "name", "wt")

	// if err1 != nil {
	// 	fmt.Println("hSet err=", err1)
	// 	return
	// }

	// _, err1 = conn.Do("hSet", "user", "age", 18)

	// if err1 != nil {
	// 	fmt.Println("hSet err=", err1)
	// 	return
	// }

	// r1, err1 := redis.String(conn.Do("hGet", "user", "name"))

	// if err1 != nil {
	// 	fmt.Println("hGet err=", err1)
	// 	return
	// }

	// fmt.Println(r1)

	// r2, err1 := redis.Int(conn.Do("hGet", "user", "age"))

	// if err1 != nil {
	// 	fmt.Println("hGet err=", err1)
	// 	return
	// }

	// fmt.Println(r2)

	_, err1 := conn.Do("hMSet", "user", "name", "wt", "age", 18)

	if err1 != nil {
		fmt.Println("hMSet err=", err1)
		return
	}

	r12, err1 := redis.Strings(conn.Do("hMGet", "user", "name", "age"))

	if err1 != nil {
		fmt.Println("hMGet err=", err1)
		return
	}

	for i, v := range r12 {
		fmt.Println(i, v)
	}
}
