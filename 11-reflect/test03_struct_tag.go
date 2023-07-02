package main

import (
	"fmt"
	"reflect"
)

// 结构体标签：增加说明

type Resume struct {
	Name   string `info:"name" doc:"名字"`
	Gender string `info:"gender"`
}

func FindTag(in interface{}) {
	t := reflect.TypeOf(in).Elem()

	for i := 0; i < t.NumField(); i++ {
		tagInfo := t.Field(i).Tag.Get("info")
		tagDoc := t.Field(i).Tag.Get("Doc")
		fmt.Println("info: ", tagInfo, "doc: ", tagDoc)
	}
}
