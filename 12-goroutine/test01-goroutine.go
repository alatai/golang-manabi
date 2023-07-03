package main

import (
	"fmt"
	"time"
)

/* goroutine（协程并发） */

func testGoroutine() {
	// 创建一个goroutine执行newTask
	// go newTask()

	// 子goroutine会随着主goroutine的退出而终止
	// fmt.Println("main goroutine exit")

	/*
		i := 0
		for {
			i++
			fmt.Printf("main goroutine : i = %d\n", i)
			time.Sleep(1 * time.Second)
		}
	*/

	// 使用go创建一个形参为空，返回值为空的goroutine
	/*
		go func() {
			defer fmt.Println("A.defer")

			// 匿名函数
			func() {
				defer fmt.Println("B.defer")
				// 退出当前goroutine
				runtime.Goexit()
				fmt.Println("B")
			}()

			fmt.Println("A")
		}()
	*/

	// 子goroutine与主goroutine是异步执行，不能通过以下方式获取子goroutine的返回值
	// flag := go func(a int, b int) bool {
	go func(a int, b int) bool {
		fmt.Println("a = ", a, ", b = ", b)
		return true
	}(10, 20)

	// 无限循环
	for {
		time.Sleep(1 * time.Second)
	}
}

// 子（从）goroutine
func newTask() {
	i := 0
	for {
		i++
		fmt.Printf("new Goroutine : i = %d\n", i)
		time.Sleep(1 * time.Second)
	}
}
