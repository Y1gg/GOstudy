package main

import (
	"fmt"

	liuyue "github/Y1gg/day18/04package/pkg" //起别名
)

func main() {
	ret := liuyue.Add(10, 20)
	fmt.Println(ret)
}
func init() {
	fmt.Println("我会被执行吗")
}
