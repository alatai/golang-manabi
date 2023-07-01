package main

import "fmt"

func main() {
	// defer的执行顺序
	defer func01()
	defer func02()
	defer func03()
	// defer和return的执行顺序
	// return在defer先被调用
	returnAndDefer()
}

func testDefer() {
	// defer关键字，在当前函数体结束之前触发
	defer fmt.Println("main end 01")
	defer fmt.Println("main end 02")

	fmt.Println("main::hello golang 01")
	fmt.Println("main::hello golang 02")
}

// defer的执行顺序
func func01() {
	fmt.Println("A")
}

func func02() {
	fmt.Println("B")
}

func func03() {
	fmt.Println("C")
}

// defer和return的执行顺序
func returnAndDefer() int {
	defer deferFunc()
	return returnFunc()
}

func deferFunc() {
	fmt.Println("deferFunc called...")
}

func returnFunc() int {
	fmt.Println("returnFunc called...")
	return 0
}
