package main

import "fmt"

func main() {

	//append()为切片追加元素
	s1 := []string{"北京", "天津", "山东", "河北"}
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n", s1, len(s1), cap(s1))

	//调用append函数必须用原来的切片变量接收返回值
	//append追加元素，原来的底层数组放不下的时候，go底层就会把底层数组换一个
	//必须用变量接收append的返回值

	s1 = append(s1, "云南")
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n", s1, len(s1), cap(s1))
	s1 = append(s1, "甘肃", "四川", "上海")
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n", s1, len(s1), cap(s1))
	s2 := []string{"山西", "陕西", "河南", "湖南", "湖北"}
	s1 = append(s1, s2...) //...表示拆开
	fmt.Printf("s1=%v\tlen(s1)=%d\tcap(s1)=%d\n", s1, len(s1), cap(s1))

}
