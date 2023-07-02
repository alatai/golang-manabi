package main

/* Golang面向对象 */

func main() {
	book02 := Book02{author: "jjj"}
	MyFunc(book02)
	MyFunc(123)
	MyFunc("abc")
}

func testClass01() {
	// 创建对象
	hero := Hero{Name: "ttt", Ad: 100, Level: 3}
	hero.Show()
	hero.SetName("777")
	hero.Show()
}

func testClass02() {
	human := Human{"zhang3", "male"}
	human.Eat()

	// 声明时继承的写法
	superman := Superman{Human{"li4", "male"}, 3}
	superman.Eat()
	superman.Walk()
	superman.Fly()
	superman.Print()
}

func testInterface01() {
	// Golang的多态
	// 接口数据类型，父类指针
	var animal IAnimal
	animal = &Cat{color: "gray"}
	// 调用Cat的方法
	animal.Sleep()
	animal = &Dog{color: "white"}
	// 调用Dog的方法
	animal.Sleep()
}
