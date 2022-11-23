package main

import "fmt"

type animal interface {
	eat(string)
}

type cat struct {
	name string
	feet int8
}

func (c cat) eat(food string) {
	fmt.Printf("%s吃%s", c.name, food)
}

func da(x animal) {
	x.eat("小黄鱼")
}

func main() {
	c1 := cat{
		"花花",
		4,
	}
	da(c1)

}
