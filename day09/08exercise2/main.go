package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	//声明切片slice和map，并初始化
	var s1 = []int{1, 2, 34, 5, 78, 9}
	var m1 = map[string]int{
		"小刘": 1,
		"小丁": 2,
		"小郑": 3,
		"小张": 4,
		"小邬": 5,
	}
	fmt.Println(s1)
	fmt.Println(m1)
	fmt.Println("\n====================================")
	//对结构体进行初始化

	//方法1:按字段初始化
	var p person
	p.name = "刘金涛"
	p.age = 29
	fmt.Println(p)
	var pp person
	pp.name = "卢志超"
	pp.age = 23
	fmt.Println(pp)
	//方法2：键值对初始化
	var p2 = person{
		name: "李腾龙",
		age:  23,
	}
	fmt.Println(p2)
	//方法3：值列表初始化
	var p3 = person{
		"王少博",
		24,
	}
	fmt.Println(p3)
	fmt.Println("\n====================================")
	//调用一下构造函数
	np1 := newPerson("梁成贇", 22)
	np2 := newPerson1("王雪鑫", 23)
	fmt.Println(np1, np2)

}

//返回值类型
func newPerson(name string, age int) person {
	return person{
		name: name,
		age:  age,
	}
}

//返回指针类型
func newPerson1(name string, age int) *person {
	return &person{
		name: name,
		age:  age,
	}
}
