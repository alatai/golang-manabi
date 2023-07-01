package main

import "fmt"

// const来定义枚举类型
// 可以再const()中添加一个关键字：iota
// 每行的iota都会累加1，第一行的iota的默认值是0
const (
	BEIJING = iota * 10
	// TOKYO   = iota(可以省略)
	TOKYO
)

const (
	a, b = iota + 1, iota + 2
	c, d
	e, f
	g, h = iota * 2, iota * 3
	i, k
)

// 常量
func main() {
	const length int = 10
	fmt.Println("length = ", length)

	fmt.Println("BEIJING = ", BEIJING)
	fmt.Println("TOKYO = ", TOKYO)

	fmt.Println("a = ", a, ", b = ", b)
	fmt.Println("c = ", c, ", d = ", d)
	fmt.Println("e = ", e, ", f = ", f)
	fmt.Println("g = ", g, ", h = ", h)
	fmt.Println("i = ", i, ", k = ", k)

	// iota只能配合const()一起使用，只有在const进行累加效果
	// var a int = iota
}
