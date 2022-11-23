package main

import "fmt"

/*
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
类型名：标识自定义结构体的名称，在同一个包内不能重复。
字段名：表示结构体字段名。结构体中的字段名必须唯一。
字段类型：表示结构体字段的具体类型。
*/

func main() {
	//匿名结构体:多用于临时场景
	var s struct {
		x string
		y int
	}
	s.x = "一二三"
	s.y = 123
	fmt.Println(s)
	fmt.Printf("结构体的类型为：%T", s)
}
