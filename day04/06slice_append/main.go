package main

import "fmt"

func main() {
	//关于append删除切片中的某个元素
	a1 := [...]int{12, 44, 35, 761, 39}
	s1 := a1[:]
	//删除索引为2的那个值
	s1 = append(s1[:2], s1[3:]...)
	fmt.Println(s1) //[12,44,761,39]
	fmt.Println(a1) //[12,44,761,39,39]
	fmt.Printf("%d %d ", len(a1), cap(a1))
	fmt.Printf("%d %d ", len(s1), cap(s1))

}
