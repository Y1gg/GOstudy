package main

import "fmt"

//类型断言
//x.(T)
//x表示类型为interface{}的变量
//T表示断言x可能的类型

//我想知道空接口中接受的值具体是什么

//类型断言1
func assign(a interface{}) {
	fmt.Printf("%T\n", a)
	str, ok := a.(string)
	if !ok {
		fmt.Println("猜错了")
	} else {
		fmt.Println("传进来的是一个字符串", str)
	}
}

//类型断言2
func assign1(a interface{}) {
	fmt.Printf("%T\n", a)
	switch t := a.(type) {
	case string:
		fmt.Println("是一个字符串", t)
	case int:
		fmt.Println("是一个int", t)
	case bool:
		fmt.Println("是一个bool", t)
	}
}

func main() {
	assign(100)
	fmt.Println("\n====================================")
	assign("刘越")
	fmt.Println("\n====================================")
	assign1(false)
}
