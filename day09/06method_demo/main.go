package main

import "fmt"

/*
什么时候应该使用指针类型接收者?
	1、需要修改接收者中的值
	2、接收者是拷贝代价比较大的大对象
	3、保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

*/

type person struct {
	name string
	age  int
}

func newPerson(n string, a int) person {
	return person{
		name: n,
		age:  a,
	}
}

//值接收者：传拷贝进去
func (p person) year() {
	fmt.Println(p.name, "过年了！")
	// p.age++
}

//指针接收者：传内存地址进去
func (p *person) birthday() {
	fmt.Println(p.name, "过生日啦")
	p.age++
}

func main() {
	per1 := newPerson("苗磊雨", 24)
	fmt.Println("过年之前", per1.age)
	per1.year()
	fmt.Println("过年之后", per1.age)
	fmt.Println("\n====================================")
	per2 := newPerson("张小艳", 22)
	fmt.Println("过生日之前", per2.age)
	per2.birthday()
	fmt.Println("过生日之后", per2.age)

}
