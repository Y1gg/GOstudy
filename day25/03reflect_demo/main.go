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

/*
type name和type kind
在反射中关于类型还划分为两种：类型（Type）和种类（Kind）。
因为在Go语言中我们可以使用type关键字构造很多自定义类型，
而种类（Kind）就是指底层的类型，======================
但在反射中，当需要区分指针、结构体等大品种的类型时，就会用到种类（Kind）。
举个例子，我们定义了两个指针类型和两个结构体类型，通过反射查看它们的类型和种类。
*/
func main() {
	var a float32 = 3.14
	reflectType(a)
	var b int64 = 1000
	reflectType(b)
	var c rune = '李'
	fmt.Println(reflect.TypeOf(c))
	fmt.Println("\n====================================")
	var cc = Cat{}
	reflectType(cc)
	fmt.Println("\n====================================")
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
