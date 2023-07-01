package main

import "fmt"

// 函数相关
func main() {
	c := func01("func01", 123)
	fmt.Println("c = ", c)

	ret1, ret2 := func02("func02", 456)
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret1, ret2 = func03("func03")
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)

	ret1, ret2 = func04("func04")
	fmt.Println("ret1 = ", ret1)
	fmt.Println("ret2 = ", ret2)
}

// 带返回值的函数
func func01(a string, b int) int {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)
	c := 100

	return c
}

// 多返回值函数（匿名)
func func02(a string, b int) (int, int) {
	fmt.Println("a = ", a)
	fmt.Println("b = ", b)

	return 666, 777
}

// 多返回值函数（别名）
func func03(name string) (r1 int, r2 int) {
	fmt.Println("name = ", name)
	// r1, r2形参初始化默认的值为0
	// r1, r2的作用域为fun2函数体的{}空间
	fmt.Println("r1 = ", r1)
	fmt.Println("r2 = ", r2)

	// 给返回值赋值
	r1 = 100
	r2 = 200

	// return后的返回值可省略
	return
}

// 多返回值函数（类型省略写法）
func func04(name string) (r1, r2 int) {
	fmt.Println("name = ", name)

	r1 = 1000
	r2 = 2000

	return r1, r2
}
