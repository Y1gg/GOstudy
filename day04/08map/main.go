package main

import "fmt"

func main() {
	//map、无序、key-value的数据结构，map是引用类型，必须初始化才能使用

	var m1 map[string]int
	//还没有初始化，没有在内存里开辟空间
	fmt.Println(m1)
	fmt.Println(m1 == nil)
	//估算好该map的容量，避免在程序运行期间再动态扩容
	m1 = make(map[string]int, 10)
	m1["孙少安"] = 35
	m1["孙少平"] = 22
	m1["田润叶"] = 30
	m1["田晓霞"] = 22

	fmt.Println(m1)

	fmt.Println("\n====================================")
	fmt.Println(m1["孙少安"])
	fmt.Println(m1["孙玉厚"]) //如果不存在这个key，得到对应值类型的零值
	//约定俗成用ok接收返回的布尔值
	v, ok := m1["孙兰花"]
	if !ok {
		fmt.Println("查无此key")
	} else {
		fmt.Println(v)

	}
	fmt.Println("\n====================================")
	//遍历map
	for i, v := range m1 {
		fmt.Println(i, v)
	}
	fmt.Println("\n====================================")
	//只遍历key
	for k := range m1 {
		fmt.Println(k)
	}
	fmt.Println("\n====================================")
	//只遍历value
	for _, v := range m1 {
		fmt.Println(v)
	}
	fmt.Println("\n====================================")
	//使用delete()函数删除键值对
	delete(m1, "田润叶")
	fmt.Println(m1)
	delete(m1, "沙河") //删除不存在的键值，没有返回值，

	//查看在线中文文档：http://studygolang.com/pkgdoc
	//查看内置函数：go doc builtin.delete
	/*
			D:\Gostudy\src\github.com\Y1gg\day04\08map>go doc builtin.delete
			package builtin // import "builtin"

			func delete(m map[Type]Type1, key Type)
		    The delete built-in function deletes the element with the specified key
		    (m[key]) from the map. If m is nil or there is no such element, delete is a
		    no-op.

	*/

}
