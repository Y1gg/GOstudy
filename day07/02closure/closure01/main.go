package main

import "fmt"

//closure 闭包

//函数名：getSequence
//没有参数
//返回值是一个函数：func() int
//返回值是一个：没有参数，返回值为int类型的函数
func getSequence() func() int {
	i := 0              //定义i
	return func() int { //匿名函数
		i += 1
		return i
	}
}

func main() {
	/* nextNumber 为一个函数，函数 i 为 0 */
	nextNumber := getSequence()

	/* 调用 nextNumber 函数，i 变量自增 1 并返回 */
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())
	fmt.Println(nextNumber())

	/* 创建新的函数 nextNumber1，并查看结果 */
	nextNumber1 := getSequence()
	fmt.Println(nextNumber1())
	fmt.Println(nextNumber1())
}
