package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	name := flag.String("name", "张旭", "你的名字")
	flag.Parse() //先解析 再使用
	fmt.Println(*name)
	fmt.Println(os.Args)
	fmt.Printf("%T\n%#v\n", os.Args, os.Args)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
