package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

/*
在Go语言中，
	使用reflect.TypeOf()函数可以获得任意值的类型对象（reflect.Type），
	程序通过类型对象可以访问任意值的类型信息。
*/

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("类型是：%v\n", v)
	fmt.Printf("type name:%v\ttyepe kind:%v\n", v.Name(), v.Kind())
}

func main() {
	test()
}

/*	类型是：float32
	类型是：int64
	int32		*/
// 当需要处理中文、日文或者其他复合字符时，则需要用到rune类型。
// rune类型实际是一个int32。

// Go语言的反射中像数组、切片、Map、指针等类型的变量，它们的.Name()都是返回空。
func test() {
	var a = [...]int{1, 23, 4, 5, 6}
	fmt.Print("数组：")
	reflectType(a)
	var b = []string{"你好", "我叫", "陆小栎", "你好吗"}
	fmt.Print("切片:")
	reflectType(b)
	var c = make(map[string]int, 10)
	c["山东"] = 12
	c["天津"] = 23
	c["上海"] = 34
	c["河北"] = 65
	fmt.Print("map:")
	reflectType(c)
	var d = 678
	var e = &d
	fmt.Print("指针:")
	reflectType(e)
}

/*
====================================
数组：类型是：[5]int
type name:      tyepe kind:array
切片:类型是：[]string
type name:      tyepe kind:slice
map:类型是：map[string]int
type name:      tyepe kind:map
指针:类型是：*int
type name:      tyepe kind:ptr
*/
