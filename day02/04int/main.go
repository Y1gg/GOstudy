package main

import "fmt"

//整数

func main() {
	i1 := 101
	fmt.Printf("%d\n", i1)
	fmt.Printf("%b\n", i1)
	fmt.Printf("%o\n", i1)
	fmt.Printf("%x\n", i1)
	i2 := 077
	fmt.Printf("%d\n", i2)
	i3 := 0x1234567
	fmt.Printf("%d\n", i3)

	fmt.Println("=================")
	//查看变量类型
	i4 := 123
	fmt.Printf("%T\n", i4)
	//声明int8类型的变量
	i5 := int8(12)
	fmt.Printf("%T\n", i5)

}
