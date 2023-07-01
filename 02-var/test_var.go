package main

import "fmt"

// 声明全局变量，方法一、方法二、方法三
var gA int = 100
var gB int = 200

// 声明全局变量，方法四
// gC := 300

// 四种变量的声明方式
func main() {
	// 方式一：声明一个变量，默认的值是0
	var a int
	fmt.Println("a = ", a)

	// 方式二：声明一个变量，初始化一个值
	var b int = 100
	fmt.Println("b = ", b)

	var bb string = "abcd"
	fmt.Printf("type of bb = %T\n", bb)

	// 方式三：在初始化的时候，可以省去数据类型，通过值自动匹配当前的变量的数据类型
	var c = 100
	fmt.Println("c = ", c)

	// 格式化输入类型
	fmt.Printf("type of a = %T\n", a)
	fmt.Printf("type of b = %T\n", b)
	fmt.Printf("type of c = %T\n", c)

	var cc = "abcd"
	fmt.Printf("type of cc = %T\n", cc)

	// 方法四：（常用方法）省去var关键字，直接自动匹配
	e := 100
	fmt.Println("e = ", e)
	fmt.Printf("type of e = %T\n", e)

	g := 3.14
	fmt.Printf("type of g = %T\n", g)

	fmt.Println("gA = ", gA)
	fmt.Println("gB = ", gB)
	// fmt.Println("gC = ", gC)

	// 声明多个变量
	var xx, yy = 100, 200
	fmt.Println("xx = ", xx, ", yy = ", yy)
	var kk, ll = 100, "abc"
	fmt.Println("kk = ", kk, ", ll = ", ll)

	// 多行的多变量声明
	var (
		vv int  = 100
		jj bool = true
	)
	fmt.Println("vv = ", vv, ", jj = ", jj)
}
