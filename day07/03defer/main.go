package main

import "fmt"

func f1() int {
	x := 12
	y := 23

	defer fmt.Println(x, y, x+y)
	x++
	defer fmt.Println(x, y, x+y)
	x++
	return 22

}

func calc(index string, a, b int) int {
	ret := a + b
	fmt.Println(index, a, b, ret)
	return ret
}
func main() {
	f1()
	a := 1
	b := 2
	defer calc("1", a, calc("10", a, b))
	a = 0
	defer calc("2", a, calc("20", a, b))
	b = 1
}

//a=1  b=2			1 1 3 4
//10 1 2 3 		a=0 	2 0  2 2
//20 0 2 2

/* 10 1 2 3
20 0 2 2
2 0 2 2
1 1 3 4   */

//压栈

// defer虽然是延迟执行，但第一遍进行的时候，遇到未完全解析的变量或者表达式，都会先进行
// 计算或者是简化，比如需要输出x+y，此时值为5，将把x+y替换为5
// 就算后面又发生了x和y 的变化，等到执行defer时 ，输出的仍然是5
