package main

import "fmt"

//结构体占用一块连续的内存空间

type x struct {
	a, b, c int8
}

func main() {
	m := x{
		int8(10), //8bit-->1byte
		int8(20),
		int8(30),
	}

	fmt.Printf("%p\n%p\n%p\n", &(m.a), &(m.b), &(m.c))
}
