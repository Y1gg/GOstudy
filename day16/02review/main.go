package main

import (
	"encoding/json"
	"fmt"
)

//json序列化与反序列化
//1、架构体内部的字段数字母要大写，不大写别人访问不到
//2、反序列化时要传递指针！

func main() {
	type point struct {
		X int `json:"lin"`
		Y int `json:"lu"`
	}

	//序列化
	p1 := point{100, 200}
	b, err := json.Marshal((p1))
	//如果出错了
	if err != nil {
		fmt.Printf("marshal failed ,err:%v\n", err)
		return
	}
	fmt.Println(string(b))
	//反序列化
	str := `{"lin":10,"lu":20}`
	var po point
	err = json.Unmarshal([]byte(str), &po)
	if err != nil {
		fmt.Printf("unmarshal failed  ,err:%v\n", err)
		return
	}
	fmt.Println(po)

	fmt.Println("\n====================================")

	//map

	m1 := map[int64]string{
		10086: "移动移不动",
		10010: "联通联不通",
		10000: "电信信不信",
	}

	n1 := m1[2000]
	fmt.Println(n1)
	//取不到就返回value类型的零值
	n2, ok := m1[3000]
	fmt.Println(n2, ok)
	//ok=true表示有这个key，
	//ok=false表示没有这个key。
	for k, v := range m1 {
		fmt.Println(k, v)
	}
	for k := range m1 {
		fmt.Println(k)
	}
	for _, v := range m1 {
		fmt.Println(v)
	}

}
