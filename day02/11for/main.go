package main

import "fmt"

//for循环

func main() {

	//基本形式
	for i := 0; i < 10; i++ {
		fmt.Printf("%d\t", i)
	}
	fmt.Println("\n===================================================")
	//变种1：省略初始语句
	var a = 0
	for ; a < 10; a++ {
		fmt.Printf("%d\t", a)
	}
	fmt.Println("\n===================================================")
	//变种2：省略结束语句
	for i := 0; i < 10; {
		fmt.Printf("%d\t", i)
		i++
	}
	fmt.Println("\n===================================================")
	//变种3：全部省略：无限循环：死循环
	for {
		fmt.Print(123)
		break
	}
	fmt.Println("\n===================================================")

	s := "流沙河123liu"
	for t, v := range s {
		fmt.Printf("%d————————%c\n", t, v)
	}

	fmt.Println("\n=======================九九乘法表练习===================")

	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d*%d=%d\t", j, i, (i * j))
		}
		fmt.Println()

	}
}
