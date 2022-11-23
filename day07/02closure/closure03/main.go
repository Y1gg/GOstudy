package main

import "fmt"

func f13(f func()) {
	fmt.Println("This is f13")
	f()
}

func f14(x, y int) {
	fmt.Println("This is f14")
	fmt.Println(x + y)
}

func f16(a func(int, int), b, c int) func() {

	ret := func() {
		a(b, c)
	}
	return ret
}
func main() {
	//闭包
	ret := f16(f14, 1, 2) //调用f16
	f13(ret)
}
