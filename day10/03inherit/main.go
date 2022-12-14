package main

import "fmt"

//结构体模拟实现其他语言中的“继承”

type animal struct {
	name string
}

//给animal实现一个移动的方法
func (a animal) move() {
	fmt.Printf("%s会动！！", a.name)
}

//狗类
type dog struct {
	feet   uint8
	animal //animal拥有的方法，dog此时也有了
}

//给dog实现一个汪汪汪的方法
func (d dog) wang() {
	fmt.Printf("%s在汪汪汪地叫 ", d.name)
}

func main() {
	d1 := dog{
		feet:   4,
		animal: animal{name: "花花"},
	}
	fmt.Println(d1)
	d1.wang()
	d1.move()
}
