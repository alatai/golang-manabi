package main

import "fmt"

// Hero 类名首字母大写表示其他包也能够访问
// Go语言的封装针对包，类名、属性名、方法名首字母大写表示对外（其他包）可以访问，否则只能够在本包访问
type Hero struct {
	// 类的属性名首字母大写表示该属性时对外能够访问
	Name  string
	Ad    int
	Level int
}

// GetName (hero Hero)表示当前方法绑定到指定结构体
func (hero *Hero) GetName() string {
	return hero.Name
}

func (hero *Hero) SetName(name string) {
	hero.Name = name
}

func (hero *Hero) Show() {
	fmt.Println("hero = ", hero)
}
