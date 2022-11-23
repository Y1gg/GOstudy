package main

import "fmt"

//defer多用于函数结束之前，释放资源（文件句柄、数据库连接、socket连接）

//defer把它后面的语句延迟到 函数即将返回的时候再执行
//一个函数中可以有多个defer语句
//多个defer语句按照先进后出的顺序延迟执行
func f0() {
	fmt.Println("这是第一步")
	defer fmt.Println("这是第3步")
	defer fmt.Println("这是第2步")
	defer fmt.Println("这是第1步")
	defer fmt.Println("这是第二步")
	fmt.Println("这是第三步")
}

func main() {
	f0()
}
