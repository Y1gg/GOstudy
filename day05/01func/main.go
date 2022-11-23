package main

import "fmt"

//函数
//函数是一段代码的封装
//把一段逻辑抽象出来封装到一个函数中，给它起一个名字，每次用到它的时候直接用函数名调用
//使用函数能够让代码结构更清晰、更简洁、

//函数的定义
//func 函数名(形参 类型)(返回值 类型){ return 返回值}
func sum(x int, y int) (ret int) {
	return x + y
}

//没有返回值
func f1(x int, y string) {
	fmt.Printf("%d,%s", x, y)
}

//没有参数，但有返回值
// func f2() (ret string) {
func f2() string {
	return "helloworld"
}

//没有参数没有返回值
func f3() {
	fmt.Println("业精于勤荒于嬉，行成于思毁于随")
}

//返回值可以命名，也可以不命名
//命名的返回值，就相当于在函数中声明了一个变量
func f4(x int, y int) (ret int) {
	ret = x + y
	return
}

//多个返回值
func f5() (int, string) {
	return 8, "奥利给"
}

//参数 的类型简写：
//当参数中连续多个参数的类型一致时，我们可以将非最后一个参数的类型省略
func f6(x, y, z int, a, b, c string, n, m bool) int {
	return x + y + z
}

//可变长参数
//可变长参数必须放在函数参数的最后
func f7(x, y string, a ...int) {
	fmt.Println(x + y)
	fmt.Println(a) //a的类型是切片，int类型的切片,[]int

}

//go语言中函数没有默认参数 这个概念

func main() {
	fmt.Println("\n====================================")
	sum := sum(21, 34)
	fmt.Println(sum)
	fmt.Println("\n====================================")
	f1(5488, "你是儿")
	fmt.Println("\n====================================")
	fmt.Println(f2())
	fmt.Println("\n====================================")
	f3()
	fmt.Println("\n====================================")
	fmt.Println(f4(12, 34))
	fmt.Println("\n====================================")
	m, n := f5()
	fmt.Println(m, n)
	fmt.Println("\n====================================")
	fmt.Println(f6(1, 2, 3, "1", "2", "3", true, false))
	fmt.Println("\n====================================")
	f7("下雨了", "打雷了")
	fmt.Println("\n====================================")
	f7("我喜", "欢你", 5, 2, 1)
	fmt.Println("\n====================================")

}
