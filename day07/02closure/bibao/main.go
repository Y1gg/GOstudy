package main

import (
	"fmt"
)

func f1(f func()) {
	fmt.Println("这是f1")
	f()
}
func f2(x, y int) {
	fmt.Println("这是f2")
	fmt.Println("\n====================================")
	fmt.Println(x + y)
}

func f3(f func(x, y int), x, y int) func() {
	temp := func() {
		f(x, y)
	}
	return temp

}

func main() {
	r1 := f3(f2, 100, 200)
	f1(r1)
	//把原来需要传递两个int类型 的参数包装成一个不需要传参的函数

}
