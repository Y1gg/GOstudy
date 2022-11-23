package main

import "fmt"

func f() func() int {
	x := 10
	// 匿名函数(返回的匿名函数可以看成是一个闭包), 闭包对它作用域上部的变量可以进行修改，修改引用的变量会对变量进行实际修改
	a := func() int {
		x += 1
		fmt.Println(x)
		return x
	}
	return a
}

func main() {
	a := f()
	a()
	a()
	a()
}
