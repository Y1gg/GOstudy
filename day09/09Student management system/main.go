package main

import (
	"fmt"
	"os"
)

/*
学生管理系统——函数版
写一个系统能够查看、新增、删除学生
*/
var (
	allStudent map[int64]*student //变量声明
)

type student struct {
	id   int64
	name string
}

// newStudent是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	//把所有学生都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号：%d\t姓名：%s\n", k, v.name)
	}

}
func addStudent() {
	//向allAtudent中添加一个新的学生
	//1、创建一个新学生
	//1.1、获取用户输入

	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)

	//1.2、造学生(调用student的构造函数)
	newStu := newStudent(id, name)
	//2、追加到 allStudent这个map中
	allStudent[id] = newStu
}

func deleteStudent() {
	//1、请用户输入要删除的学生的学号
	var (
		deleteID int64
	)
	fmt.Println("请输入学生学号：")
	fmt.Scanln(&deleteID)
	//2、去allStudent这个map中根据学号删除对应的键值对
	delete(allStudent, deleteID)
}

func main() {
	allStudent = make(map[int64]*student, 50) //初始化，开辟内存空间

	for {
		//1、打印菜单
		fmt.Println("欢迎使用学生管理系统Student management system！")
		fmt.Println(`
	1、查看学生信息
	2、新增学生信息
	3、删除学生信息
	4、退出学生系统
	`)
		fmt.Print("请输入你要进行的操作编号(1~3):")
		//2、等待用户选择要做什么
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择执行的操作是：%d选项\n", choice)
		//3、执行对应的函数
		switch choice {
		case 1:
			showAllStudent()
		case 2:
			addStudent()
		case 3:
			deleteStudent()
		case 4:
			os.Exit(1) //退出
		default:
			fmt.Println("您输入的选项有误，请重新输入")
		}
	}
}
