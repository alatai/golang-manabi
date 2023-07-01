package main

import "fmt"

// Golang中操作数组
func main() {
	// 切片的拷贝
	s := []int{1, 2, 3} // len = 3, cap = 3
	// [0, 2)
	s1 := s[0:2]
	fmt.Println("s1 = ", s1) // s1 =  [1 2]

	// s和s1指向同一个存储空间
	s1[0] = 100
	fmt.Println("s = ", s, ", s1 = ", s1)

	// copy可以将底层数组的slice一起进行拷贝
	s2 := make([]int, 3)
	// 将s中的值依次拷贝到s2中
	copy(s2, s)
	s2[0] = 200
	fmt.Println("s2 = ", s2)
	fmt.Println("s = ", s, ", s1 = ", s1)
}

// len()和cap()函数
func sliceOpt02() {
	var numbers = make([]int, 3, 5)
	// cap：表示容量
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 3, cap = 5, slice = [0 0 0]

	// 向numbers追加一个元素1
	numbers = append(numbers, 1)
	numbers = append(numbers, 2)
	// cap超出后会自动追加，追加长度为cap长度（若未设置cap长度，则等于len长度）
	numbers = append(numbers, 3)
	fmt.Printf("len = %d, cap = %d, slice = %v\n", len(numbers), cap(numbers), numbers) // len = 3, cap = 5, slice = [0 0 0]
}

// 切片的声明
func sliceOpt01() {
	// 切片，声明slice01是一个切片，并且初始化，默认值为1，2，3，长度为3
	slice01 := []int{1, 2, 3}
	// %v：表示任何变量的详细信息
	fmt.Printf("len = %d, slice01 = %v\n", len(slice01), slice01) // len = 3, slice = [1 2 3]

	// 声明slice02是一个切片，但是并没有给slice02分配空间
	var slice02 []int
	// 通过make函数开辟空间，默认值为0
	// slice02 = make([]int, 3)
	fmt.Printf("len = %d, slice02 = %v\n", len(slice02), slice02) // len = 3, slice = [0 0 0]

	// 声明slice03，并开辟空间，默认值为0
	// var slice03 []int = make([]int, 3)
	slice03 := make([]int, 4)                                     // 常用
	fmt.Printf("len = %d, slice03 = %v\n", len(slice03), slice03) // len = 3, slice = [0 0 0]

	// 判断一个slice是否为空
	if slice02 == nil {
		fmt.Println("slice02是一个空切片")
	} else {
		fmt.Println("slice02已经开辟空间")
	}
}

// 数组基础
func arrayFoundationOpt() {
	// 固定长度的数组
	// 声明数组，默认值0, 0, 0...
	var myArr01 [10]int

	for i := 0; i < 10; i++ {
		myArr01[i] = i
	}

	fmt.Println("myArr01 = ", myArr01)

	// 声明并初始化数组
	myArr02 := [8]int{0, 1, 2, 3, 4}
	fmt.Println("myArr02 = ", myArr02)

	// 利用range函数打印数组
	for index, value := range myArr02 {
		fmt.Println("index = ", index, ", value = ", value)
	}

	// 查看数组的数据类型
	// 注意：[10]int与[8]int为两种不同类型的数据类型
	fmt.Printf("myArray01 type = %T\n", myArr01) // [10]int
	fmt.Printf("myArray02 type = %T\n", myArr02) // [8]int

	// 动态数组，切片slice
	myArr03 := []int{1, 2, 3, 4, 5}
	printArray02(myArr03)
	fmt.Println("myArr03 = ", myArr03)
}

// 此时只允许[4]int类型的参数，实际是值拷贝
func printArray01(myArr [4]int) {
	for index, value := range myArr {
		fmt.Println("index = ", index, ", value = ", value)
	}
}

// 此时参数为动态数组类型，为引用传递
func printArray02(myArr []int) {
	// _：表示匿名变量
	for _, value := range myArr {
		fmt.Println(", value = ", value)
	}

	myArr[0] = 100
}
