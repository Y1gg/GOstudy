package main

import "fmt"

//接口的实现

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) move() {
	fmt.Println("走猫步")
}

func (c cat) eat(food string) {
	fmt.Printf("猫吃%s\n", food)
}

func main() {
	var a1 animal //定义一个接口类型的变量
	var c1 = cat{ //定义一个cat类型的变量
		"米米",
		4,
	}

	a1 = c1
	fmt.Println(a1)
	a1.eat("小黄鱼")

}
