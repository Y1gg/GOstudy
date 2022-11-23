package main

import (
	"flag"
	"fmt"
	"time"
)

func main() {
	var name string

	flag.StringVar(&name, "name", "张旭", "姓名")
	var age int
	flag.IntVar(&age, "age", 23, "年龄")
	var married bool
	flag.BoolVar(&married, "married", false, "婚否")
	var delay time.Duration
	flag.DurationVar(&delay, "d", 0, "时间间隔")
	flag.Parse()
	fmt.Println(name, age, married, delay)
}
