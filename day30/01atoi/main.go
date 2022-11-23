package main

import (
	"fmt"
	"strconv"
)

func main() {
	//`Atoi()`函数用于将字符串类型的整数转换为int类型，函数签名如下。
	// func Atoi(s string) (i int, err error)
	str1 := "1345634"
	num1, err := strconv.Atoi(str1)
	if err != nil {
		fmt.Println("can't strconv to int")
	}
	fmt.Printf("num1的类型:%T\tnum1的值:%#v\n", num1, num1)

	//Itoa()
	// `Itoa()`函数用于将int类型数据转换为对应的字符串表示，具体的函数签名如下。
	// func Itoa(i int) string
	num2 := 100
	str2 := strconv.Itoa(num2)
	fmt.Printf("str2的类型:%T\tstr2的值:%#v\n", str2, str2)

}
