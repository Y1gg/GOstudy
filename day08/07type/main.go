package main

import "fmt"

type myInt int     //自定义类型
type yourInt = int //类型别名

func main() {
	var num1 myInt = 100
	fmt.Println(num1)
	fmt.Printf("%T", num1)
	fmt.Println("\n====================================")
	var num2 yourInt = 200
	fmt.Println(num2)
	fmt.Printf("%T", num2)
	fmt.Println("\n====================================")
	var c rune = '月'
	fmt.Println(c)
	fmt.Printf("%T", c)
	//rune就是int32的别名

}
