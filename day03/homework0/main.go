package main

import "fmt"

func main() {
	count := 0
	str := "你好啊hello沙河小王子"
	for _, i := range str {
		length := len(string(i)) //获取每个字符串的长度
		if length >= 3 {         //判断字符长度
			fmt.Printf("%s\t", string(i))
			count++
		}
	}
	fmt.Printf("\nchinese number is %d\n", count)
}
