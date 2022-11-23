package main

import "fmt"

func main() {
	//copy

	a1 := []int{1, 3, 5, 7}
	a2 := a1 //赋值
	var a3 = make([]int, 4)
	copy(a3, a1)
	fmt.Println(a1, a2, a3)
	a1[0] = 100
	fmt.Println(a1, a2, a3)

	//从切片中删除元素
	a := []int{11, 12, 13, 14, 15, 16, 17}
	//要删除元素为13的值
	a = append(a[:2], a[3:]...)
	fmt.Println(a)
	fmt.Println(cap(a))
	fmt.Println("\n====================================")
	x1 := [...]int{1, 3, 5, 7, 9}
	x2 := x1[:]
	fmt.Println(x2, len(x1), cap(x1))

	x2 = append(x2[:2], x2[3:]...)
	fmt.Println(x2, len(x2), cap(x2))

	fmt.Println(x1)
	fmt.Println("\n====================================")

}
