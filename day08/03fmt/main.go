package main

import "fmt"

func main() {
	fmt.Print("直接打印，不会换行")
	fmt.Println("\n====================================")
	fmt.Println("直接打印，然后换行")
	fmt.Println("====================================")
	//fmt.Printf("格式化字符串",值)
	//%T：查看类型
	//%d：十进制数	%b：二进制数	%o：八进制数	%x：十六进制数
	//%c：字符	该值对应的unicode码值
	//%s：字符串		%p：指针		%v：值
	//%f：浮点数	%t：布尔值		%%:百分号
	//%q 整数=>字符	该值对应的双引号括起来的go语法字符串字面值

	/* 	%v 按默认格式输出
	%+v 在%v的基础上额外输出字段名
	%#v 在%+v的基础上额外输出类型名 	*/

	var m1 = make(map[string]int, 1)
	m1["李响 "] = 12
	fmt.Printf("%v\n", m1)
	fmt.Printf("%#v\n", m1)
	fmt.Println("\n====================================")
	fmt.Printf("%q\n", 65)

	//宽度标识符
	//%f	默认宽度，默认精度
	//%9f：宽度9，默认精度
	//%9.3f:宽度9，精度3
	//%9.0f:宽度9，精度0
	//%.3f:默认宽度，精度3
	n := 123.345
	fmt.Printf("%f\n", n)
	fmt.Printf("%9f\n", n)
	fmt.Printf("%.2f\n", n)
	fmt.Printf("%9.0f\n", n)
	fmt.Printf("%9.2f\n", n)
	fmt.Println("\n====================================")

	//Scan系列
	//Scan\Scanf\Scanln
	// var s string
	// fmt.Scan(&s)
	// fmt.Print(s)
	// fmt.Println()
	// var s1 string
	// fmt.Scanf("%s", &s1)
	// fmt.Printf("获得了s为=%s", s1)
	var (
		name  string
		age   int
		isBoy bool
	)
	fmt.Scanln(&name, &age, &isBoy)
	fmt.Println("姓名年龄是男的吗", name, age, isBoy)

}
