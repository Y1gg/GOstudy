package main

import "fmt"

var a []int
var b chan int //需要指定通道中元素的类型

func main() {
	fmt.Println(b)
	b = make(chan int) //通道的初始化
	fmt.Println(b)

}
