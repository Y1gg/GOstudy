package main

import "fmt"

const (
	eat   int = 4
	sleep int = 2
	play  int = 1
)

//111
//左边的1表示吃饭100
//中间的1表示睡觉010
//右边的1表示打豆豆001

func f(arg int) {
	fmt.Printf("%b\n", arg)
}

func main() {
	f(eat)
	f(eat | play)
	f(eat | sleep | play)
	fmt.Println()
}
