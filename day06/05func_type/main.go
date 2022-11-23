package main

import "fmt"

//函数类型

func f1() {
	fmt.Println("hello，清风")
}
func f2() int {
	fmt.Println("hello，秋天")
	return 1
}

//函数也可以作为参数的类型
func f3(x func() int) {
	ret := x()
	fmt.Println(ret)
}
func f4(x, y int) int {
	return x + y
}

//函数还可以作为返回值
// func f5(x func() int) func() {

// }

func ff(a, b int) int {
	return a + b
}
func f6(x func() int) func(int, int) int {
	return ff
}
func main() {
	a := f1
	fmt.Printf("%T\n", a)

	b := f2
	fmt.Printf("%T\n", b)

	f3(f2)
	f3(b)
	fmt.Printf("%T\t", f3)
	fmt.Printf("%T\t\n", f4)

	f7 := f6(f2)
	fmt.Printf("%T\n", f7)

}
