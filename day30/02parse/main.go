package main

import (
	"fmt"
	"strconv"
)

func main() {
	/*  ParseBool()
	func ParseBool(str string) (value bool, err error)
	返回字符串表示的bool值。它接受1、0、t、f、T、F、true、false、True、False、TRUE、FALSE；否则返回错误。*/

	str1 := "1"
	bool2, err := strconv.ParseBool(str1)
	if err != nil {
		fmt.Printf("The value passed in is wrong,err:%v", err)
		return
	}
	fmt.Printf("type:%T\tvalue:%#v\n", bool2, bool2)
	fmt.Println("\n====================================")
	var str string
	fmt.Println("请输入你想转换的布尔值：")
	fmt.Scan(&str)
	bool1, err := strconv.ParseBool(str)
	if err != nil {
		fmt.Printf("The value passed in is wrong,err:%v", err)
		return
	}
	fmt.Printf("type:%T\tvalue:%#v\n", bool1, bool1)

}
