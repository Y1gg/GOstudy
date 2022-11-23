package main

import (
	"flag"
	"fmt"
	"time"
)

// flag获取命令行参数
func main() {
	//创建一个标志位参数
	name := flag.String("name", "张旭", "你的名字")
	age := flag.Int("age", 23, "你的芳龄")
	married := flag.Bool("married", false, "婚否")
	cTime := flag.Duration("ct", time.Second, "结婚时间")
	//使用flag
	flag.Parse()
	fmt.Println(*name, *age, *married, *cTime)

	fmt.Println(flag.Args())
	fmt.Println(flag.NArg())
	fmt.Println(flag.NFlag())
}
