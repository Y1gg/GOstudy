package main

import "fmt"

//结构体嵌套
type address struct {
	city     string
	province string
}

type person struct {
	name string
	age  int
	add  address
}

//匿名嵌套结构体
type teacher struct {
	name    string
	worknum uint64
	address
}

func main() {
	p1 := person{
		name: "刘斌",
		age:  55,
		add: address{
			city:     "深圳",
			province: "广东省",
		},
	}
	fmt.Println(p1)

	t1 := teacher{
		name:    "王其秋",
		worknum: 12332,
		address: address{
			"西青区",
			"天津市",
		},
	}
	fmt.Println(t1)

}
