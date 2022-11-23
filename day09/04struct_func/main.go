package main

import "fmt"

//构造函数：返回一个结构体变量的函数

type person struct {
	name string
	age  int
}

//构造函数：约定成俗用new开头
//返回的是结构体还是结构体指针
//当结构体比较大的时候尽量使用结构体指针，减少程序的内存开销
func newPerson(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}

func main() {
	p1 := newPerson("刘龙", 22)
	p2 := newPerson("石玉帅", 23)
	fmt.Println(p1, p2)

}
