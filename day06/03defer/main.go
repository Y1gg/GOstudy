package main

import "fmt"

//go语言中函数的return不是原子操作，在底层是分为两步来执行
//第一步：返回值赋值
//第二步：真正的ret返回
//函数中如果存在defer，那么defer执行的时机是在第一步和第二步之间

func f1() int {
	x := 5
	defer func() {
		x++
	}()
	return x
	//第11行x值为5
	//第15行将x的值返回，到了第10行int对应的
	//执行第13行，x的值发生改变，但int对应的并不是x，所以返回值没有发生变化.还是5
}

func f2() (x int) {
	defer func() {
		x++
	}()
	return 5
	//第25行返回的值是5，将5给到第21行的int对应的x变量
	//执行defer里的第23行，x+1，x即为第21行的x，5+1
	//所以最终的结果是返回值x为6
}

func f3() (y int) {
	x := 5
	defer func() {
		x++
	}()
	return x
	//第32行x的值为5
	//第36行将x的值返回给第31行的int对应的y，目前返回值y为5
	//第34行，x自增，x值为6
	//x的变化和y没有关系，所以返回值y为5
}
func f4() (x int) {
	defer func(x int) {
		x++
	}(x)
	return 5
	//先执行第46行，返回值5，返回给第42行的int对应的x。
	//对于整个函数f4来说 ，第42行的x为全局变量
	//再执行defer里的第44行，x自增。但此x是在43行新的函数里面定义的x，为局部变量
	//局部变量不影响第42行的全局变量，不影响返回值，
	//所以返回值还是5

}
func main() {
	fmt.Println(f1())
	fmt.Println(f2())
	fmt.Println(f3())
	fmt.Println(f4())
}

// a=0;b=1

// 201
// 20201

// 101
// 10101
