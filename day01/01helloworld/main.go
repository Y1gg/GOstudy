package main

//导入语句
import "fmt"

// var student_nname string
// var studentName string
var StudentName string

//函数外只能放置标识符(变量、常量、函数、类型)的声明

//单独声明
var s1 string
var s2 int
var s3 bool

var name3 = "yui"

//批量声明
var (
	name string = "liu" //""
	age  int    = 22    // 0
	isok bool   = false // false
)

//程序的入口
func main() {
	name = "理想"
	age = 22
	isok = true

	fmt.Println("Hello world!")
	fmt.Print("s1=" + s1 + "\n")

	fmt.Printf("s1=%s,s2=%d", s1, s2)

	fmt.Print("\n")
	fmt.Print(s3)
	fmt.Print("\n")
	fmt.Print("\n")

	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(isok)

	//声明变量同时赋值
	var name1 string = "yue"
	fmt.Println(name1)
	//类型推导
	var name2 = "liu"
	fmt.Println(name2)
	//简短变量声明（只能在函数里面使用）
	fmt.Print(name3)
	name3 := "yue"
	fmt.Println(name3)

	//常用批量声明和简短变量声明
	//同一个作用域中{}不能重复声明同名变量

	//匿名变量: 	_

	//常量 把var换成了const

}
