package main

import (
	"fmt"
	"reflect"
)

// 通过反射设置变量的值=========================================
// 想要在函数中通过反射修改变量的值，需要注意函数参数传递的是值拷贝，
// 必须传递变量地址才能修改变量值。而反射中使用专有的Elem()方法来获取指针对应的值。

func reflectSetValue1(x interface{}) {
	v := reflect.ValueOf(x)
	if v.Kind() == reflect.Int64 {
		v.SetInt(200)
		//修改的是副本，reflect包会引发painc
	}
}

func reflectSetValue2(x interface{}) {
	v := reflect.ValueOf(x)
	//反射中使用Elem()方法获取指针对应的值
	if v.Elem().Kind() == reflect.Int64 {
		v.Elem().SetInt(300)
	}
}

func main() {
	var a int64 = 1
	// reflectSetValue1(a)
	// panic: reflect: reflect.Value.SetInt using unaddressable value
	// 恐慌：反射：反射。价值。SetInt使用无法寻址的值

	reflectSetValue2(&a)
	fmt.Println(a)
}
