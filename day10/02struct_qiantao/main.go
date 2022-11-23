package main

import "fmt"

//结构体嵌套

/*
字段	field
嵌套	nesting
省份	province
公司	company
*/
type address struct {
	province string
	city     string
}
type person struct {
	name string
	age  int
	add  address
}

type company struct {
	//匿名嵌套结构体
	name    string
	email   string
	address //这是个匿名的
}

func main() {
	p1 := person{
		name: "阿林",
		age:  22,
		add: address{
			province: "山东",
			city:     "泰安",
		},
	}
	fmt.Println(p1)
	fmt.Println(p1.name, p1.age)
	fmt.Println(p1.add.province, p1.add.city)
	fmt.Println("\n====================================")
	c1 := company{
		name:  "山安药业",
		email: "1713194910@qq.com",
		address: address{
			province: "河北",
			city:     "石家庄",
		},
	}
	fmt.Println(c1)
	fmt.Println(c1.name, c1.email)
	fmt.Println(c1.address)
	fmt.Println(c1.address.city, c1.address.province)
	fmt.Println(c1.city, c1.province)
	//先在自己的结构体找这个字段，找不到就去匿名嵌套的结构体中查找该字段
	//如果发生“匿名嵌套结构的字段冲突”，就老老实实写全了

}
