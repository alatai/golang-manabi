package main

import (
	"fmt"
	"io"
	"os"
)

// 类型pair

func testPair01() {
	var a string
	// pair<type:string, value:"abcd">
	a = "abcd"

	var allType interface{}
	// pair<type:string, value:"abcd>
	allType = a
	str, _ := allType.(string)
	fmt.Println(str)
}

func testPair02() {
	// tty: pair<type:*os.File, value:/dev/tty>
	tty, err := os.OpenFile("/dev/tty", os.O_RDWR, 0)
	if err != nil {
		fmt.Println("open file error", err)
	}

	// reader: pair<type:, value:>
	var reader io.Reader
	// reader: pair<type:*os.File, value:/dev/tty>
	reader = tty
	// writer: pair<type:, value:>
	var writer io.Writer
	// writer: pair<type:*os.File, value:/dev/tty>
	writer = reader.(io.Writer)
	_, _ = writer.Write([]byte("HELLO THIS IS A TEST!!!\n"))
}

type Reader interface {
	ReadBook()
}

type Writer interface {
	WriteBook()
}

type Book struct {
}

func (book *Book) ReadBook() {
	fmt.Println("Read a book...")
}

func (book *Book) WriteBook() {
	fmt.Println("Write a book...")
}
