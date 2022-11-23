package main

import (
	"fmt"
)

//切片就是一个框，框住了一块连续的内存
//切片属于引用类型，真正的数据都是保存在底层数组里的

func main() {
	//make()函数创造切片
	//make([]T, size, cap)
	//T是指切片内元素的类型
	//size是指切片的长度，切片内元素的数量
	//cap是指切片的容量
	s0 := make([]int, 5, 10)
	s00 := make([]int, 8)
	fmt.Printf("%d,s0的len为:%d,s0的cap为%d\n", s0, len(s0), cap(s0))
	fmt.Printf("%d,s00的len为:%d,s00的cap为%d\n", s00, len(s00), cap(s00))

	fmt.Println("====================================")
	//切片不能直接不比较
	//切片唯一合法的比价操作是和nil比较
	//一个nil值的切片并没有底层数组
	//一个nil值的切片长度和容量都是0
	//但我们不能说一个长度和容量都是 0 的切片一定是nil

	var s1 []int         //0,0,等于nil
	s2 := []int{}        //0,0，不等于nil
	s3 := make([]int, 0) //0,0，不等于nil
	fmt.Printf("%d,s1的len为:%d,s1的cap为%d\n", s1, len(s1), cap(s1))
	fmt.Println(s1 == nil)
	fmt.Printf("%d,s2的len为:%d,s2的cap为%d\n", s2, len(s2), cap(s2))
	fmt.Println(s2 == nil)
	fmt.Printf("%d,s3的len为:%d,s3的cap为%d\n", s3, len(s3), cap(s3))
	fmt.Println(s3 == nil)
	//所以要判断一个切片是否为空，要使用len(s)==nil来判断，不应该使用s==nil来判断
	fmt.Println("====================================")

	s4 := []int{1, 2, 3, 4}
	s5 := s4
	//s4和s5都指向了同一个底层数组
	fmt.Println(s4, s5)
	s5[3] = 100
	fmt.Println(s4, s5)

	fmt.Println("====================================")
	//切片的遍历：索引遍历和range遍历
	s6 := []int{1, 2, 3, 4, 56, 78, 9}
	for i := 0; i < len(s6); i++ {
		fmt.Printf("%d\t", s6[i])
	}
	fmt.Println()

	for i, v := range s6 {
		fmt.Println(i, v)
	}
	//用户补充代码配置
	//点击ctrl+shift+p、输入snippet、点击用户相关user、选择go进行配置

}
