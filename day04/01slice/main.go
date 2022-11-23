package main

import "fmt"

func main() {
	//切片的定义
	var s1 []int    //定义一个存放int类型元素的切片
	var s2 []string //定义一个存放string类型元素的切片

	fmt.Println(s1, s2)
	//为空，所以是true
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//初始化
	s1 = []int{1, 2, 3, 4, 5, 6}
	s2 = []string{"C语言", "go语言", "java语言"}
	fmt.Println(s1)
	fmt.Println(s2)
	//初始化完不为空，所以是flase
	fmt.Println(s1 == nil)
	fmt.Println(s2 == nil)

	//长度和容量
	fmt.Printf("len(s1):%d\tcap(s1):%d\n", len(s1), cap(s1))
	fmt.Printf("len(s2):%d\tcap(s2):%d\n", len(s2), cap(s2))

	//2.由数组得到切片
	ar1 := [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	//基于一个数组切割，左包含右不包含（左闭右开）
	s3 := ar1[0:4] //[0, 1, 2, 3]从第0个元素到第3个元素
	fmt.Println(s3)
	s4 := ar1[1:7] //[ 1, 2, 3, 4, 5, 6]从第1个元素到第6个元素
	fmt.Println(s4)
	s5 := ar1[4:] //[ 4, 5, 6, 7, 8, 9]从第4个元素到最后，相当于ar1[0:len(ar1)]
	s6 := ar1[:5] //[0, 1, 2, 3, 4]从第0个元素到第5个元素，相当于ar1[0:5]
	s7 := ar1[:]  //[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]从第0个元素到最后，相当于ar1[0:len(ar1)]

	fmt.Println(s5, s6, s7)

	//切片的容量是指底层数组的容量
	//底层数组从切片的第一个元素到最后一个元素的数量

	fmt.Printf("len(ar1):%d\tcap(ar1):%d\n", len(ar1), cap(ar1))
	fmt.Printf("len(s4):%d\tcap(s4):%d\n", len(s4), cap(s4))
	fmt.Printf("len(s5):%d\tcap(s5):%d\n", len(s5), cap(s5))

	//切片再切割
	s8 := s6[3:] //[3,4]
	fmt.Println(s8)
	fmt.Printf("len(s8):%d\tcap(s8):%d\n", len(s8), cap(s8))

	//切片是引用类型，都指向了低层的一个数组
	//修改了底层数组的值，切片的值会发生改变

	fmt.Println(s6) //[0,1,2,3,4]
	ar1[2] = 1000   //[0, 1, 2, 3, 4, 5, 6, 7, 8, 9]
	fmt.Println(s6) //[0,1,1000,3,4]

	fmt.Println("====================================")

}
