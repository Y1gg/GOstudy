package main

import (
	"fmt"
)

func main() {
	//匿名函数
	//函数内部没有办法声明带名字的函数

	f1 := func(x, y int) {
		fmt.Println("x+y:", x+y)
	}
	f1(100, 220)

	//如果只是调用一次的函数，可以简写成立即执行函数
	func(x, y int) {
		fmt.Println("你好啊")
		fmt.Println(x * y)
	}(100, 111)
	func() {
		fmt.Println("你好啊,没有参数的立即执行函数")
	}()

}
