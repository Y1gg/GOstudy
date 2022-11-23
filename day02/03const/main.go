package main

import "fmt"

//常量 const

const pi = 3.1415926
const (
	statusOK = 200
	noFound  = 404
)

//批量声明常量时，如果某一行没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

//iota:类似枚举
//常量计数器：
//1.在const关键字出现时将被重置为0
//2.const中每新增一行常量声明将使iota计数一次

const (
	a1 = iota
	a2
	a3
)

//插队
const (
	b1 = iota
	b2 = 100
	b3
	b4 = iota
)

const (
	c1, c2 = iota + 1, iota + 2
	c3, c4 = iota + 1, iota + 2
)

//定义数量级
const (
	_  = iota
	KB = 1 << (10 * iota)
	MB = 1 << (10 * iota)
	GB = 1 << (10 * iota)
	TB = 1 << (10 * iota)
	PB = 1 << (10 * iota)
)

func main() {
	fmt.Println("n1:", n1)
	fmt.Println("n2:", n2)
	fmt.Println("n3:", n3)
	fmt.Println("====================================")
	fmt.Println(statusOK)
	fmt.Println(noFound)
	fmt.Println(pi)
	fmt.Println("====================================")
	fmt.Println("a1:", a1)
	fmt.Println("a2:", a2)
	fmt.Println("a3:", a3)
	fmt.Println("====================================")
	fmt.Println("b1:", b1)
	fmt.Println("b2:", b2)
	fmt.Println("b3:", b3)
	fmt.Println("b4:", b4)
	fmt.Println("====================================")
	fmt.Println("c1:", c1)
	fmt.Println("c2:", c2)
	fmt.Println("c3:", c3)
	fmt.Println("c4:", c4)

}
