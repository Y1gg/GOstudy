package main

import "fmt"

//同一个结构体可以实现多个接口

//接口还可以嵌套

type animal interface {
	eater
	mover
}

type eater interface {
	eat()
}
type mover interface {
	move()
}

type dragon struct {
	name string
	feet int8
}

//dragon实现了move接口
func (d *dragon) move() {
	fmt.Println("龙飞凤舞")
}

//dragon实现了eat接口
func (d *dragon) eat() {
	fmt.Println("吞天食地")
}

func main() {
	var a1 animal
	d1 := dragon{
		"小神龙",
		12,
	}
	a1 = &d1
	fmt.Println(a1)

}
