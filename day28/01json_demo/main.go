package main

import (
	"encoding/json"
	"fmt"
)

//json
/*
Json(Javascript Object Nanotation)是一种数据交换格式，
常用于前后端数据传输。任意一端将数据转换成json 字符串，
另一端再将该字符串解析成相应的数据结构，如string类型，strcut对象等。
*/

type person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	str := `{"name":"周琳","age":590}`
	var p person
	json.Unmarshal([]byte(str), &p)
	fmt.Println(p.Name, p.Age)

}
