package main

import (
	"fmt"
	"strings"
)

func a(name string) func() {
	return func() {
		fmt.Println("你好，" + name)
	}
}

func makeSuffixFunc(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func calc(base int) (func(int) int, func(int) int) {
	add := func(i int) int {
		base += i
		return base
	}

	sub := func(i int) int {
		base -= i
		return base
	}

	return add, sub
}

// 闭包 = 函数 ＋ 外部环境(外层变量的引用)
func main() {
	c := a("天软彭于晏") //c此时就是一个闭包
	c()             //相当于执行了a函数内部的匿名函数
	r := makeSuffixFunc("玉盘")
	ret := r("小时不识月呼作白玉盘")
	fmt.Println(ret)
	fmt.Println("\n====================================")
	x, y := calc(100)
	ret1 := x(200)
	fmt.Println(ret1)
	ret2 := y(200)
	fmt.Println(ret2)
}
