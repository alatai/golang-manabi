package lib01

import (
	"fmt"
	_ "golangManabi/05-init/lib02"
)

// init函数，golang包调用过程中先会执行import包中init函数
// 初始化操作可以在init函数值执行
func init() {
	fmt.Println("lib01 do init...")
}

// TestLib01 （模块中要导出的函数，函数名首字母大写。函数名首字母大写就表示公有）
func TestLib01() {
	fmt.Println("test by lib01...")
}
