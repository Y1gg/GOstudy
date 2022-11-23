package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string `json:"name" zhoulin:"嘿嘿嘿"`
	Score int    `json:"score" zhoulin:"哈哈哈"`
}

// 给student添加两个方法 Study和Sleep(注意首字母大写)
func (s student) Study() string {
	msg := "好好学习，天天向上。"
	fmt.Println(msg)
	return msg
}

func (s student) Sleep() string {
	msg := "好好睡觉，快快长大。"
	fmt.Println(msg)
	return msg
}

func printMethod(x interface{}) {
	t := reflect.TypeOf(x)
	v := reflect.ValueOf(x)

	fmt.Println(t.NumMethod()) //方法的个数	2
	for i := 0; i < v.NumMethod(); i++ {
		methodType := v.Method(i).Type()
		fmt.Printf("method name:%s\n", t.Method(i).Name)
		fmt.Printf("method:%s\n", methodType)
		// 通过反射调用方法传递的参数必须是 []reflect.Value 类型
		var args = []reflect.Value{}
		v.Method(i).Call(args)
	}
}

func printField(s interface{}) {
	t := reflect.TypeOf(s)
	//任意值通过reflect.TypeOf()获得反射对象信息后，

	fmt.Println(t.Name(), t.Kind())
	//student	struct

	//通过for循环遍历结构体的所有字段信息
	fmt.Println(t.NumField()) //字段数量	2
	// tag	标签
	for i := 0; i < t.NumField(); i++ {
		filed := t.Field(i)
		fmt.Printf("name:%s\tindex:%d\ttype:%v\t json tag: %v\n", filed.Name, filed.Index, filed.Type, filed.Tag.Get("zhoulin"))
	}

	//通过字段名获得指定结构体字段信息
	if scoreFiled, ok := t.FieldByName("Score"); ok {
		fmt.Printf("name:%s\tindex:%d\ttype:%v\t json tag:%v\n", scoreFiled.Name, scoreFiled.Index, scoreFiled.Type, scoreFiled.Tag.Get("json"))
	}
}

func main() {
	stu1 := student{
		Name:  "虹猫",
		Score: 100,
	}
	fmt.Println("打印字段信息")
	printField(stu1)
	fmt.Println("打印方法信息")
	printMethod(stu1)

}
