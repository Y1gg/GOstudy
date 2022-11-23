package main

import "fmt"

//数组
//存放元素的容器
//必须指定存放的元素的类型和容量（长度）
//数组的长度是数组类型的一部分

func main() {
	var ar1 [3]bool
	var ar2 [4]bool

	fmt.Printf("ar1:%T\t%v\nar2:%T\t%v\n", ar1, ar1, ar2, ar2)

	//数组的初始化
	//如果不初始化，默认都是零值。布尔是flase整型和浮点型都是0，字符串是""

	//初始化方式1
	ar1 = [3]bool{false, true, true}
	fmt.Println(ar1)

	//初始化方式2
	//根据初始值自动判断数组的长度是多少
	ar3 := [9]int{1, 2, 3, 4, 5, 67, 8}
	ar4 := [...]int{1, 2, 34, 5, 5, 667, 9, 8, 78, 78, 56, 4}
	fmt.Printf("ar3:%T\t%v\n", ar3, ar3)
	fmt.Printf("ar4:%T\t%v\n", ar4, ar4)

	//初始化方式3：根据索引初始化
	var ar5 = [5]int{0: 1, 4: 5}
	fmt.Printf("ar4:%T\t%v\n", ar5, ar5)

	fmt.Println("========1.根据索引遍历数组===================")

	//数组的遍历
	//1.根据索引遍历数组
	ar6 := [...]string{"小何", "小刘", "小倩", "小王", "小颖", "小月"}
	for i := 0; i < len(ar6); i++ {
		fmt.Println(i, ar6[i])
	}
	fmt.Println("\n=======2.for range遍历====================")
	//2.for range遍历
	for i, v := range ar6 {
		fmt.Print(i, v)
		fmt.Printf("\t")
	}
	fmt.Println("\n=========多维数组==================")

	//多维数组
	var ar7 = [3][2]int{
		{1, 2},
		{3, 4},
		{5, 6}}

	fmt.Println(ar7)
	fmt.Println("\n===========多维数组的遍历================")
	//多维数组的遍历
	for i := 0; i < len(ar7); i++ {
		for j := 0; j < len(ar7[0]); j++ {
			fmt.Printf("%d\t", ar7[i][j])
		}
		fmt.Println()
	}
	fmt.Println("\n===========多维数组的遍历================")
	for _, v := range ar7 {
		for _, i := range v {
			fmt.Printf("%d\t", i)
		}
		fmt.Println()
	}
	fmt.Println("\n=========数组是值类型，赋值不会改变原来的==================")
	//数组是值类型，赋值不会改变原来的
	ar8 := [3]int{1, 2, 3}
	ar9 := ar8
	ar9[0] = 99
	fmt.Println(ar8)
	fmt.Println(ar9)
	//数组支持“==”、“！=”操作符，因为内存总是被初始化过的，因为是值类型
	//[n]*T	表示指针数组
	//*[n]T	表示数组指针
	fmt.Println("\n=========求数组内所有元素的和==================")
	ar10 := [...]int{1, 3, 5, 7, 9, 11, 13, 15, 17, 19}
	sum := 0
	for _, v := range ar10 {
		sum += v
	}
	fmt.Printf("数组ar10的元素和为：%d", sum)

	fmt.Println("\n=========找出数组中和为指定值的两个元素的下标==================")
	num := 8
	ar11 := [...]int{1, 3, 5, 4, 7, 8}
	for i := 0; i < len(ar11); i++ {
		for j := i + 1; j < len(ar11); j++ {
			if ar11[i]+ar11[j] == num {
				fmt.Printf("%d+%d=%d,(%d,%d)\n", ar11[i], ar11[j], ar11[i]+ar11[j], i, j)
			}

		}

	}
}
