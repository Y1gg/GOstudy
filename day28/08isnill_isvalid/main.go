package main

import (
	"fmt"
	"reflect"
)

func main() {
	// IsNil()常被用于判断指针是否为空；IsValid()常被用于判定返回值是否有效。
	var a *int //*int类型空指针
	fmt.Println("var a *int IsNil:", reflect.ValueOf(a).IsNil())
	fmt.Println("nil IsValid:", reflect.ValueOf(nil).IsValid())
	fmt.Println("\n====================================")
	b := struct{}{}
	//尝试从结构体查找"abc"字段
	fmt.Println("不存在的结构体成员：", reflect.ValueOf(b).FieldByName("abc").IsValid())
	//尝试从结构体查找"abc"方法
	fmt.Println("不存在的结构体方法：", reflect.ValueOf(b).MethodByName("abc").IsValid())
	c := map[string]int{}
	//尝试从map中查找一个不存在的键
	fmt.Println("查找map中不存在的键:", reflect.ValueOf(c).MapIndex(reflect.ValueOf("娜扎")).IsValid())
}
