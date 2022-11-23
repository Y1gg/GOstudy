package main

import "fmt"

//结构体是值类型

type person struct {
	name, gender string
}

func f1(x person) {
	x.name = "喜欢谁"
}
func f2(x *person) {
	x.name = "石凯"
}

//创建指针类型的结构体
//还可以通过new关键字对结构体进行实例化，得到的是结构体的地址
func main() {
	var p2 = new(person)
	fmt.Printf("%T\n", p2)
	fmt.Printf("p2=%#v\n", p2)
	//p2是结构体指针
	var p person
	p.name = "陆小栎"
	p.gender = "男"
	f1(p)
	fmt.Println(p.name)
	f2(&p)
	fmt.Println(p.name)
	fmt.Println("\n====================================")
	var p3 = new(person)
	p3.name = "何畅"
	fmt.Printf("%T\n", p3)
	fmt.Printf("%x\n", *p3)
	fmt.Printf("%p\n", p3)

}
