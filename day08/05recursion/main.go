package main

import "fmt"

func f1(num uint64) uint64 {
	if num > 0 {
		return num * f1(num-1)
	} else {
		return 1
	}
}

//上台阶面试题
//n个台阶，一次可以走1步，也可以走2步,有多少种走法
func f2(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	return f2(n-1) + f2(n-2)
}

func main() {
	// fmt.Println("请输入想要计算阶乘的数字")
	var num uint64 = 5
	// fmt.Scanln(&num)
	fmt.Printf("%d!=%d", num, f1(num))
	fmt.Println("\n====================================")
	ret := f2(5)
	fmt.Println(ret)

}
