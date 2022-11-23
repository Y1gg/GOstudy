package main

import (
	"fmt"
	"os"
)

// 函数版学生管理系统(查看、新增、删除)
var (
	allStudent map[int64]*student //变量声明
)

type student struct {
	id   int64
	name string
}

// 构造方法，便于创建学生对象
// newStudent是student类型的构造函数
func newStudent(id int64, name string) *student {
	return &student{
		id:   id,
		name: name,
	}
}

func showAllStudent() {
	//遍历——把所有的学生的都打印出来
	for k, v := range allStudent {
		fmt.Printf("学号:%d\t姓名:%s", k, v.name)
		fmt.Println()
	}
}

func addStudent() {
	var (
		id   int64
		name string
	)
	fmt.Print("请输入学生学号：")
	fmt.Scanln(&id)
	fmt.Print("请输入学生姓名：")
	fmt.Scanln(&name)

	newstu := newStudent(id, name)

	allStudent[id] = newstu
}

func deleteStudent() {
	var (
		deleteID int64
	)
	fmt.Print("请输入你想删除的学号:")
	fmt.Scanln(&deleteID)
	delete(allStudent, deleteID)

}

func main() {
	allStudent = make(map[int64]*student, 50) //初始化，开辟内存空间
	for {
		fmt.Println("欢迎使用学生管理系统")
		fmt.Println(`
		1、查看学生
		2、新增学生
		3、删除学生
		4、退出系统
		`)
		fmt.Print("请问你想进行什么操作(1——4):")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你选择了%d选项\n", choice)
		switch choice {
		case 1:
			fmt.Println("=======学生信息如下=======")
			showAllStudent()
		case 2:
			fmt.Println("=======可以新增学生=======")
			addStudent()
		case 3:
			fmt.Println("=======可以删除学生 =======")
			deleteStudent()
		case 4:
			fmt.Println("=======可以退出系统=======")
			os.Exit(1) //退出
		default:
			fmt.Println("=======输入错误！请重新输入！")
		}

	}

}
