package main

import "fmt"

func main() {
	//1.取地址& 2.根据地址取值*
	n := 12
	fmt.Println(&n)
	p := &n
	fmt.Println(p)
	fmt.Printf("%T\n", p) //*int，int类型的指针
	t := *p
	fmt.Println(t)

	fmt.Println("\n====================================")
	//1.取地址& 2.根据地址取值*
	//new函数申请一个内存地址
	var a1 *int
	fmt.Println(a1) //nil
	var a2 = new(int)
	fmt.Println(a2)  //地址
	fmt.Println(*a2) //0
	*a2 = 100
	fmt.Println(*a2) //100

	//make和new的区别：都是用来申请内存的
	//new很少用 ，一般用来给基本数据类型申请内存，string\int，返回的是对应类型的指针
	//make是用来给slicp、map、chan申请内存的，make函数返回的是对应的这三个类型的本身

}
