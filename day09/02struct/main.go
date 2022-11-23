package main

import "fmt"

type person struct {
	name, gender string
}

func main() {
	//结构体指针
	//1、key-value初始化
	var p3 = &person{
		name:   "天蓬",
		gender: "男",
	}
	fmt.Printf("%#v\n", p3)
	//2、使用值列表的形式初始化
	var p4 = &person{
		"大泼猴",
		"男",
	}
	fmt.Printf("%#v\n", p4)

	/*
		注意：p4值列表的形式初始化，不仅值的顺序要和结构体保持一致，
		另外值要全部填上，否则会无法保存代码，如果想要填入空值，只能用“”符号占位
	*/
}
