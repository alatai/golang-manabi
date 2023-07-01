package main

import (
	// 匿名导包？解决导入但不使用编译错误
	_ "golangManabi/05-init/lib01"
	// 别名导包
	// myLib02 "golang-manabi/05-init/lib02"
	// “.”表示导入包的函数全部“转移”到当前go文件，调用时可以直接使用函数名
	// 建议不使用，若出现同名函数会产生歧义
	. "golangManabi/05-init/lib02"
)

func main() {
	// 导入包后若未使用相关函数，则会报错
	// lib01.TestLib01()
	// lib02.TestLib02()
	// 使用别名
	// myLib02.TestLib02()
	// 直接使用函数名调用
	TestLib02()
}
