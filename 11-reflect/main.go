package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	var num float64 = 1.2345
	reflectNum(num)

	user := User{Id: 1, Name: "AceId", Age: 18}
	DoFiledAndMethod(user)

	var re Resume
	FindTag(&re)
}

func jsonStructTag() {
	movie := Movie{
		Title: "喜剧之王",
		Year:  2000,
		Price: 10,
		Actor: []string{"zhouxingchi", "zhangbozhi"},
	}

	// 编码：结构体 -> json
	jsonStr, err := json.Marshal(movie)
	if err != nil {
		fmt.Println("json marshal error", err)
		return
	} else {
		fmt.Printf("jsonStr = %s\n", jsonStr)
	}

	myMovie := Movie{}
	// 解码：json -> 结构体
	err = json.Unmarshal(jsonStr, &myMovie)
	if err != nil {
		fmt.Println("json unmarshal error ", err)
		return
	} else {
		fmt.Println(myMovie)
	}
}

func reflectNum(arg interface{}) {
	fmt.Println("value = ", reflect.ValueOf(arg), ", type = ", reflect.TypeOf(arg))
}

func testBookPair() {
	book := &Book{}
	// pair: <type: , Value:>
	var r Reader
	// pair: <type: Book, Value:book{}地址>
	r = book
	r.ReadBook()

	// pair: <type: , Value:>
	var w Writer
	// pair: <type: Book, Value:book{}地址>
	// 此处的断言为什么会成功，因为w、r具体的type是一致的
	w = r.(Writer)
	w.WriteBook()
}
