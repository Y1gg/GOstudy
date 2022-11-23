package main

import "fmt"

/*
	空接口:没有必要起名字，通常定义成下面的格式：
	interface{}
	所有的类型都实现了空接口
	 任意类型的变量都能实现空接口
*/

//空接口作为函数参数
func show(a interface{}) {
	fmt.Printf("type:%T\t\tvalue:%v\n", a, a)
}

func main() {
	var m1 map[string]interface{}
	fmt.Println("\n====================================")
	m1 = make(map[string]interface{}, 22)
	m1["name"] = "范倩倩"
	m1["age"] = 19
	m1["sex"] = "女"
	m1["isBoy"] = true
	m1["hobby"] = [...]string{"打代码", "跑步", "游泳"}
	fmt.Println(m1)
	fmt.Println("\n====================================")
	show(1234567)
	show(true)
	show(m1)
	show("苏州桥")

}
