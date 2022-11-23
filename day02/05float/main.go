package main

import (
	"fmt"
	"math"
)

//浮点数

func main() {
	//math.MaxFloat32()  float32最大值

	f1 := 1.23456
	fmt.Printf("%T\n", f1) //默认go语言中的小数都是float64位
	f2 := float32(1.234)
	fmt.Printf("%T\n", f2)

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.4f\n", math.Pi)

}
