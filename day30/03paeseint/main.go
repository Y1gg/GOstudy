package main

import (
	"fmt"
	"strconv"
)

func main() {
	// func ParseInt(s string, base int, bitSize int) (i int64, err error)
	// ParseInt()返回字符串表示的整数值，接受正负号。

	str1 := "256456787980033456761"
	num1, err := strconv.ParseInt(str1, 10, 64)
	if err != nil {
		fmt.Printf("有错误,err:%v\n", err)
		return
	}
	str2 := "-567"
	num2, err := strconv.ParseInt(str2, 10, 64)
	if err != nil {
		fmt.Printf("有错误,err:%v\n", err)
		return
	}
	fmt.Printf("num1\ttype:%T\tvalue:%v\n", num1, num1)
	fmt.Printf("num2\ttype:%T\tvalue:%v\n", num2, num2)
}
