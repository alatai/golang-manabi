package main

import "fmt"

// MyFunc interface{}是万能数据类型（同Java中的Object比较类型）
func MyFunc(arg interface{}) {
	fmt.Println("MyFunc is called...")
	fmt.Println(arg)

	// interface{}底层如何区别此时的应用数据类型？
	// 类型断言机制
	value, ok := arg.(string)
	if !ok {
		fmt.Println("arg is not a string type")
	} else {
		fmt.Println("arg is a string type, value = ", value)
	}
}

type Book02 struct {
	author string
}
