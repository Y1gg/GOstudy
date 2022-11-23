package main

import "fmt"

//使用值接收者实现接口与使用指针接收者实现接口的区别？
//使用值接收者实现接口，结构体类型和结构体指针类型的变量都能存
//指针接收者实现接口，只能存结构体指针类型的变量

type animal interface {
	move()
	eat(string)
}

type cat struct {
	name string
	feet int8
}

type dog struct {
	name, sex string
}

//方法使用值接收者
func (c cat) move() {
	fmt.Println("我爱你中国，亲爱的母亲")
}

func (c cat) eat(food string) {
	fmt.Printf("%s有%d条腿，爱吃%s\n", c.name, c.feet, food)
}

//使用指针接收者实现接口的所有方法
func (d *dog) move() {
	fmt.Println("国庆快乐！！！")
}

func (d *dog) eat(food string) {
	fmt.Printf("%s是%s的，爱吃%s\n", d.name, d.sex, food)
}

func main() {
	var a1 animal
	fmt.Println(a1)

	c1 := cat{"淘气", 4}  //cat
	c2 := &cat{"咖喱", 6} //*cat

	a1 = c1
	fmt.Println(a1)
	a1.eat("纸包鱼")
	a1.move()

	a1 = c2
	fmt.Println(a1)
	a1.eat("大盘鸡面")
	fmt.Println("\n====================================")
	d1 := &dog{"小黑", "公"}
	d2 := &dog{"小佛", "母"}
	a1 = d1
	fmt.Println(a1)
	a1.eat("骨头")
	a1.move()
	a1 = d2
	fmt.Println(a1)
	a1.eat("鸡头")
	a1.move()

}
