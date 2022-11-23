package main

import "fmt"

func main() {
	var n = 200
	fmt.Printf("%T\n", n)
	fmt.Printf("%v\n", n)
	fmt.Printf("%b\n", n)
	fmt.Printf("%o\n", n)
	fmt.Printf("%d\n", n)
	fmt.Printf("%x\n", n)

	var str = "花面郎"
	fmt.Printf("%s\n", str)
	fmt.Printf("%T\n", str)
	fmt.Printf("%v\n", str)
	fmt.Printf("%#v\n", str)
	fmt.Println("\n====================================")
	var num1 = 20
	var num2 = 12.34
	var isBoy = true
	var str1 = "你好 北京"
	fmt.Printf("整型:类型:%T\t\t值:%v\n", num1, num1)
	fmt.Printf("浮点型:类型:%T\t值:%v\n", num2, num2)
	fmt.Printf("布尔型:类型:%T\t值:%v\n", isBoy, isBoy)
	fmt.Printf("字符串类型:类型:%T\t值:%v\n", str1, str1)

}
