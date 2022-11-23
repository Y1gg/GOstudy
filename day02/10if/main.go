package main

import "fmt"

//if条件判断

func main() {
	//单个条件
	age := 20
	if age >= 18 {
		fmt.Println("该上大学了")
	} else {
		fmt.Println("还未成年！")
	}
	fmt.Println("===================================================")
	//多个条件
	if age <= 18 {
		fmt.Println("少年")
	} else if age > 18 && age < 35 {
		fmt.Println("青年")
	} else {
		fmt.Println("中老年")
	}
	fmt.Println("===================================================")
	//变形
	if a := 99; a > 10 {
		fmt.Println("比10大")
	} else {
		fmt.Println("比10小")
	}
}
