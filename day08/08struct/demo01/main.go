package main

import "fmt"

//结构体
type person struct {
	name   string
	age    int
	gender string
	hobby  []string
}

func main() {
	//声明一个person类型的变量p
	var p person
	//通过字段赋值
	p.name = "陆小栎"
	p.age = 18
	p.gender = "男"
	p.hobby = []string{
		"篮球", "羽毛球", "象棋",
	}
	fmt.Println(p)
	fmt.Println("\n====================================")
	//访问变量p的字段
	fmt.Printf("%T\n", p)
	fmt.Println(p.name)
	var p2 person
	p2.name = "张明阳 "
	p2.age = 23
	fmt.Printf("类型：%T\t值为%v", p2, p2)
}
