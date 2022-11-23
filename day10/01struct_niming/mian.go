package main

import "fmt"

//匿名字段：没有名字的字段
//默认把类型当成名字了。
//缺点：相同类型只能写一个。
//字段比较少，也比较简单。不常用。
type person struct {
	string
	int
}

func main() {
	p := person{
		"鹿小瓯",
		12,
	}
	fmt.Println(p)
	fmt.Println(p.string)
	fmt.Println(p.int)

}
