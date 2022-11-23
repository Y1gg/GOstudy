package main

import (
	"fmt"
	"os"
)

//学生管理系统

var smr studentMgr //声明一个全局的变量，学生管理对象smr

// 菜单函数
func showMenu() {
	fmt.Println("welcome sms!")
	fmt.Println(`
	1、查看所有学生
	2、添加学生信息
	3、修改学生信息
	4、删除学生信息
	5、退出学生系统
	`)

}

type student struct {
	id   int64
	name string
}

// 造一个学生管理者
type studentMgr struct {
	allStudent map[int64]student
}

// 查看学生
func (s studentMgr) showStudents() {
	for _, stu := range s.allStudent {
		fmt.Printf("学号:%d\t姓名:%s\n", stu.id, stu.name)
	}
}

// 增加学生
func (s studentMgr) addStudent() {
	var (
		stuID   int64
		stuName string
	)
	fmt.Print("请输入学号：")
	fmt.Scanln(&stuID)
	fmt.Print("请输入姓名：")
	fmt.Scanln(&stuName)
	newStu := student{
		id:   stuID,
		name: stuName,
	}
	s.allStudent[newStu.id] = newStu
	fmt.Println("添加成功！")
}

func (s studentMgr) editStudent() {
	var stuID int64
	fmt.Print("请输入学号")
	fmt.Scanln(&stuID)
	//展示学号对应学生的信息，如果没有提示“查无此人”
	stuObj, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	fmt.Printf("你要修改的学生信息如下：\n学号:%d,姓名:%s\n", stuObj.id, stuObj.name)
	fmt.Print("请输入更改后的名字：")
	var newName string
	fmt.Scanln(&newName)
	stuObj.name = newName
	s.allStudent[stuID] = stuObj
}

func (s studentMgr) deleteStudent() {
	var stuID int64
	fmt.Print("请输入学号")
	fmt.Scanln(&stuID)
	//展示学号对应学生的信息，如果没有提示“查无此人”
	_, ok := s.allStudent[stuID]
	if !ok {
		fmt.Println("查无此人")
		return
	}
	delete(s.allStudent, stuID)
	fmt.Println("删除成功！")

}

func main() {
	smr = studentMgr{
		allStudent: make(map[int64]student, 100),
	}
	for {
		showMenu()
		fmt.Print("请输入序号:")
		var choice int
		fmt.Scanln(&choice)
		fmt.Printf("你输入的序号是%d\n", choice)

		switch choice {
		case 1:
			smr.showStudents()
		case 2:
			smr.addStudent()
		case 3:
			smr.editStudent()
		case 4:
			smr.deleteStudent()
		case 5:
			fmt.Println("系统正在退出.................")
			os.Exit(1)
		default:
			fmt.Println("你输入的有误，请重新输入:")
		}
	}
}
