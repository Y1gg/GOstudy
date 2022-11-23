package main

import "fmt"

//只有当有两个或两个以上的具体类型必须以相同的方式进行处理时才需要定义接口。
//切记不要为了使用接口类型而增加不必要的抽象，导致不必要的运行时损耗。
type animal interface {
	mover
	eater
}

type mover interface {
	move()
}

type eater interface {
	eat()
}

type people struct {
	name    string
	sex     string
	age     uint8
	address string
}

func (p people) eat() {
	fmt.Println("吃好吃的")
}
func (p people) move() {
	fmt.Println("蹦蹦跳跳")
}
func main() {
	var a1 animal
	p1 := people{
		"刘猛",
		"男",
		22,
		"天津市西青区",
	}
	a1 = p1
	a1.eat()
	a1.move()
}
