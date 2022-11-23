package main

import (
	"encoding/json"
	"fmt"
)

//结构体和json
//1、序列化：把go语言的结构体变量——>json格式的字符串
//2、反序列化：json格式的字符串——>go语言中能够实现别的结构体变量

/*
	在Go语言中，没有特别的关键字来声明一个方法、函数或者类型是否为公开的，
	Go语言提供的是以大小写的方式进行区分的，
	如果一个类型的名字是以大写开头，那么其他包就可以访问；
	如果以小写开头，其他包就不能访问。
*/

type person struct {
	Name string `json:"name" db:"name" ini:"name"`
	Age  int    `json:"age"`
}

func main() {
	p1 := person{
		Name: "粒粒",
		Age:  18,
	}

	//序列化
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("marshal failed,err:%v", err)
		return
	}
	fmt.Printf("%#v\n", string(b))
	fmt.Println("\n====================================")
	//反序列化
	str := `{"name":"李响","age":18}`
	var p2 person
	json.Unmarshal([]byte(str), &p2)
	//传指针是为了能在	json.Unmarshal()函数内部修改p2的值
	fmt.Printf("%#v\n", p2)

}
