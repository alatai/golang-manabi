package main

import "fmt"

type Human struct {
	name   string
	gender string
}

func (human *Human) Eat() {
	fmt.Println("Human.Eat()...")
}

func (human *Human) Walk() {
	fmt.Println("Human.walk()...")
}

// Superman 继承Human
type Superman struct {
	Human // 继承了Human的方法和属性
	level int
}

// Eat 方法重载
func (superman *Superman) Eat() {
	fmt.Println("Superman.Eat()...")
}

// Fly 子类的新方法
func (superman *Superman) Fly() {
	fmt.Println("Superman.fly()...")
}

func (superman *Superman) Print() {
	fmt.Println(superman)
}
