package main

import "fmt"

//构造函数
type person struct {
	name  string
	age   int
	isBoy bool
}

//构造(结构体变量的)函数,返回值是对应的结构体类型
func newPerson(n string, a int, sex bool) person {

	return person{
		name:  n,
		age:   a,
		isBoy: sex,
	}

}

//方法
//1、方法的接收者使用对应类型的首字母小写
//2、指定了接收者之后，只有接收者才能调用这个方法
//3、保持一致性：如果有一个方法使用指针接收者，其他的方法为了统一也要使用指针接收者
func (p person) dream(str string) {
	fmt.Printf("%s的梦想是%s", p.name, str)
}

func main() {
	p1 := newPerson("刘猛", 22, true)
	fmt.Println(p1)
	p1.dream("程序员赚钱")

}
