package main

import (
	"fmt"
	"sort"
)

func main() {
	//切片的练习题
	var q = make([]int, 5, 10)
	fmt.Println(q)
	for i := 0; i < 10; i++ {
		q = append(q, i)
	}
	fmt.Println(q)
	fmt.Println(len(q), cap(q))
	fmt.Println("\n====================================")
	//切片的练习题
	var a = [...]int{3, 5, 7, 3, 133, 55, 22, 8, 9, 00, 2}
	sort.Ints(a[:]) //对切片进行排序
	fmt.Println(a)
}
