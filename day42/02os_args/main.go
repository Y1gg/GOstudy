package main

import (
	"fmt"
	"os"
)

//os.Args
//获取命令行参数

func main() {
	fmt.Println(os.Args)
	fmt.Printf("%T\n%#v\n", os.Args, os.Args)
}
