package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

type StuRead struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
	Test  interface{}
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//json字符中的"引号，需用\进行转义，否则编译出错
	//json字符串沿用上面的结果，但对key进行了大小的修改，并添加了sex数据
	data := "{\"name\":\"张三\",\"Age\":18,\"high\":true,\"sex\":\"男\",\"CLASS\":{\"naME\":\"1班\",\"GradE\":3}}"
	str := []byte(data)
	stu := StuRead{}
	fmt.Println(stu) //打印json解析前变量类型
	printType(&stu)
	err := json.Unmarshal(str, &stu)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("--------------json 解析后-----------")
	printType(&stu)
	fmt.Println(stu) //打印json解析后变量类型

}

// 利用反射，打印变量类型
func printType(stu *StuRead) {
	nameType := reflect.TypeOf(stu.Name)
	ageType := reflect.TypeOf(stu.Age)
	highType := reflect.TypeOf(stu.HIgh)
	sexType := reflect.TypeOf(stu.sex)
	classType := reflect.TypeOf(stu.Class)
	testType := reflect.TypeOf(stu.Test)

	fmt.Println("nameType:", nameType)
	fmt.Println("ageType:", ageType)
	fmt.Println("highType:", highType)
	fmt.Println("sexType:", sexType)
	fmt.Println("classType:", classType)
	fmt.Println("testType:", testType)
}
