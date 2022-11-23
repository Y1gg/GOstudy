package main

import "fmt"

//recover()必须搭配defer使用
//defer一定要在可能引发panic的语句之前定义

func funcA() {
	fmt.Println("A")
}
func funcB() {

	defer func() {
		//如果程序出现了panic错误，可以通过recover恢复过来
		err := recover()
		fmt.Println(err)
		fmt.Println("释放数据库连接")
	}()
	panic("出现了严重的错误！！！")
	// fmt.Println("B")
}

func funcC() {
	fmt.Println("C")
}

func main() {
	funcA()
	funcB()
	funcC()
}
