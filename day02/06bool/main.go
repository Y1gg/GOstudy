package main

import "fmt"

func main() {
	b1 := true
	var b2 bool
	fmt.Printf("%T\n", b1)
	fmt.Printf("%T values:%v\n", b2, b2)
	var c1 complex64
	c1 = 1 + 2i
	var c2 complex128
	c2 = 2 + 3i
	fmt.Println(c1)
	fmt.Println(c2)
}
