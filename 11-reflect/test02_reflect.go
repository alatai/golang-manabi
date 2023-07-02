package main

import (
	"fmt"
	"reflect"
)

// 反射

type User struct {
	Id   int
	Name string
	Age  int
}

func (user *User) Call() {
	fmt.Println("User is called..")
	fmt.Println(user)
}

func DoFiledAndMethod(in interface{}) {
	// 获取in的type
	t := reflect.TypeOf(in)
	fmt.Println("inType is ", t.Name())
	// 获取in的value
	v := reflect.ValueOf(in)
	fmt.Println("inValue is ", v)

	// 分别通过type获取内部的字段
	// 1.获取interface的reflect.Type，通过Type得到NumField，进行遍历
	for i := 0; i < t.NumField(); i++ {
		// 2.得到每个field，数据类型
		field := t.Field(i)
		// 3.通过field的Interface()方法获取对应的vale
		value := v.Field(i).Interface()
		fmt.Println("field = ", field, ", value = ", value)
	}

	// 分别通过type获取内部的方法
	for i := 0; i < t.NumField(); i++ {
		method := t.Method(i)
		fmt.Println("method = ", method)
	}
}
