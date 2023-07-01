package main

import "fmt"

func main() {
	a := 10
	b := 20

	swap(&a, &b)
	fmt.Println("a = ", a, ", b = ", b)

	// 一级指针
	var p *int
	p = &a
	fmt.Println("&a = ", &a, ", p = ", p)
	// 二级指针
	var pp **int
	pp = &p
	fmt.Println("&p = ", &p, ", pp = ", pp)
}

// 通过引用传递进行值交换
func swap(a *int, b *int) {
	tmp := *a
	*a = *b
	*b = tmp
}
