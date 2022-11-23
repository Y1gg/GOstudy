package main

import "fmt"

//类型断言

func main() {
	var a interface{} //定义一个空接口
	fmt.Println("\n====================================")
	//判断a保存的值的具体的类型
	//类型断言
	//1、	x.(T)
	a = 100
	v1, ok := a.(float32)
	if ok {
		fmt.Println("猜对了,a是float32", v1)
	} else {
		fmt.Println("猜错了,a不是float32")
	}

	//2、switch()
	switch v2 := a.(type) {
	case int8:
		fmt.Println("int8", v2)
	case int16:
		fmt.Println("int16", v2)
	case string:
		fmt.Println("string", v2)
	case int:
		fmt.Println("int", v2)
	default:
		fmt.Println("你好,我不知道")
	}

}
