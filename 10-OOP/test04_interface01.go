package main

import "fmt"

/*
多态的基本要素
1.有一个父类（有接口）
2.有子类（实现了父类的全部接口方法）
3.父类类型的变量（指针）执行（引用）子类的具体数据变量
*/

// IAnimal （接口）本质是一个指针
// 实现三个方法就等同于实现了接口（不需要显示写出来）
type IAnimal interface {
	Sleep()
	GetColor() string // 获取动物的颜色
	GetType() string  // 获取动物的种类
}

// Cat 具体的类
type Cat struct {
	color string
}

func (cat *Cat) Sleep() {
	fmt.Println("Cat.Sleep()...")
}

func (cat *Cat) GetColor() string {
	return cat.color
}

func (cat *Cat) GetType() string {
	return "Cat"
}

// Dog 具体的类
type Dog struct {
	color string
}

func (dog *Dog) Sleep() {
	fmt.Println("Dog.Sleep()...")
}

func (dog *Dog) GetColor() string {
	return dog.color
}

func (dog *Dog) GetType() string {
	return "Cat"
}
