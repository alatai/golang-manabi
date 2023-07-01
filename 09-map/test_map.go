package main

import "fmt"

func main() {
	// map的基本操作
	cityMap := make(map[string]string)
	// 添加
	cityMap["China"] = "Beijing"
	cityMap["Japan"] = "Tokyo"
	cityMap["USA"] = "NewYork"
	fmt.Println("cityMap = ", cityMap)

	// 遍历
	printMap(cityMap)
	// 删除
	delete(cityMap, "China")
	// 修改
	cityMap["USA"] = "DC"
	fmt.Println("cityMap = ", cityMap)
}

// 引用传递
func printMap(myMap map[string]string) {
	for key, value := range myMap {
		fmt.Println("key = ", key, ", value = ", value)
	}
}

// map的声明
func mapOpt01() {
	// 声明方式一
	// 声明myMap01是一个map类型的变量，key是string、value是string，
	var myMap01 map[string]string
	// 此时myMap01尚未初始化，为nil
	fmt.Println(myMap01 == nil) // true

	// 再使用map前，需要先用make给map分配存储空间
	myMap01 = make(map[string]string, 3)
	myMap01["one"] = "java"
	myMap01["two"] = "golang"
	myMap01["three"] = "python"

	// 底层为哈希，乱序
	fmt.Println("myMap01 = ", myMap01)

	// 声明方式二
	// 存储空间可以省略（动态）
	myMap02 := make(map[int]string)
	myMap02[1] = "java"
	myMap02[2] = "golang"
	myMap02[3] = "python"
	fmt.Println("myMap02 = ", myMap02)

	// 声明方式三
	// 声明并初始化
	myMap03 := map[string]string{
		"one":   "java",
		"two":   "golang",
		"three": "python",
	}
	fmt.Println("myMap03 = ", myMap03)
}
