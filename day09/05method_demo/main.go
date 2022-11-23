package main

import "fmt"

/*
func (接收者变量 接收者类型) 方法名(参数列表) (返回参数) {
    函数体
}
其中，
1、接收者变量：接收者中的参数变量名在命名时，官方建议使用接收者类型名称首字母的小写，
	而不是self、this之类的命名。例如，Person类型的接收者变量应该命名为 p，
	Connector类型的接收者变量应该命名为c等。
2、接收者类型：接收者类型和参数类似，可以是指针类型和非指针类型。
3、方法名、参数列表、返回参数：具体格式与函数定义相同。
*/

/* //方法
//go语言中如果标识符首字母是大写的。就是表示对外部包可见的（暴露的、公共的）

//Dog 这是一个狗的结构体
type Dog struct {
} */

type dog struct {
	name string
}

//构造函数
func newDog(name string) dog {
	return dog{
		name: name,
	}
}

//方法是作用于特点类型变量的函数
//接收者表示的是调用该方法的具体类型变量，多用类型名首字母小写表示
func (d dog) wang() {
	fmt.Printf("%s:汪汪汪汪~~~", d.name)
}

func main() {
	d1 := newDog("小黑")
	d1.wang()
}
