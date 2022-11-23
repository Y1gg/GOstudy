package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

//go语言中函数参数永远是拷贝
func f(x person) {
	x.gender = "女"
}

func f2(x *person) {

	//根据内存地址找到那个原变量，修改的就是原来的变量
	x.gender = "女" //(*x).gender = "女"	语法糖，自动根据指针找对应的变量

}

func main() {
	var p person
	p.name = "狼来了"
	p.gender = "公"
	f(p)
	fmt.Println(p.gender)
	fmt.Println("\n====================================")
	fmt.Println(&p)
	f2(&p)
	fmt.Println(p.gender)
	fmt.Println("\n====================================")
	var p2 = new(person)
	fmt.Printf("%T\n", p2)

}
