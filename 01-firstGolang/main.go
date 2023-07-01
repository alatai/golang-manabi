package main // 表示当前程序的包，main函数为main包（包名和文件夹名无关)

/*
import "fmt"
import "time"
*/

// 导包的几种写法（一下推荐）
import (
	"fmt"
	"time"
)

// main函数
func main() { // “{”必须紧跟函数名在同一行，否则编译错误
	// golang中的表达式，加“;”或不加“;”都可以（建议不加）
	fmt.Println("Hello Golang!")
	time.Sleep(1 * time.Second)
}
