package main

import "fmt"

//map和slice组合
func main() {

	//元素类型为map的切片
	var s1 = make([]map[int]string, 10, 20)
	//没有对内部的map进行初始化

	s1[0] = make(map[int]string, 1)
	s1[0][10] = "路遥"
	s1[0][11] = "路远"
	s1[0][12] = "路远"
	s1[0][13] = "路远"
	fmt.Println(s1)

	fmt.Println("\n====================================")
	//值为切片类型的map
	var m1 = make(map[string][]int, 10)
	m1["北京"] = []int{10, 20, 30}

	fmt.Println(m1)

	//统计字符串中每个单词出现的次数
	//how do you do

	fmt.Println("\n====================================")
	// str = "how do you do"

}
