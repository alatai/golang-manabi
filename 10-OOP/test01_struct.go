package main

import "fmt"

// Book 定义一个Book结构体
type Book struct {
	title  string
	author string
}

// 结构体基本定义与使用
func testStruct() {
	var book01 Book
	book01.title = "Golang"
	book01.author = "ttt"

	fmt.Println(book01) // {Golang ttt}
	changBook(book01)
	fmt.Println(book01) // {Golang ttt}
	changeBook02(&book01)
	fmt.Println(book01) // {Golang 777}
}

func changBook(book Book) {
	// 传递一个book的副本
	book.author = "666"
}

func changeBook02(book *Book) {
	// 指针传递
	book.author = "777"
}
