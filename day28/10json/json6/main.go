package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  interface{} `json:"name"`
	Age   interface{}
	HIgh  interface{}
	sex   interface{}
	Class interface{} `json:"class"`
}

type Class struct {
	Name  string
	Grade int
}

func main() {
	//实例化一个数据结构，用于生成json字符串
	stu := Student{
		Name: "张三",
		Age:  18,
		HIgh: true,
		sex:  "男",
	}

	//指针变量
	cla := new(Class)
	cla.Name = "1班"
	cla.Grade = 3
	stu.Class = cla

	//Marshal失败时err!=nil
	jsonStu, err := json.Marshal(stu)
	if err != nil {
		fmt.Println("生成json字符串错误")
	}

	//jsonStu是[]byte类型，转化成string类型便于查看
	fmt.Println(string(jsonStu))
}
