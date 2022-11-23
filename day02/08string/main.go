package main

import (
	"fmt"
	"strings"
)

func main() {
	str1 := "\"D:\\Gostudy\\src\\github.com\\Y1gg\\day02\\08string\""
	fmt.Println(str1)
	str2 := "I'm ok !"
	fmt.Println(str2)

	fmt.Println("===================================================")

	//多行的字符串:原样输出，包括空格换行
	str3 := `
世情薄
	人情恶
		雨送黄昏花易落	
	`
	str4 := `D:\Gostudy\src\github.com\Y1gg\day02\08string`
	fmt.Println(str3)
	fmt.Println(str4)

	fmt.Println("===================================================")

	//求长度 len()
	s0 := "abcdefghijklmnopqrstuvwxyz"
	s1 := "陆小栎到此一游！哈哈哈"
	fmt.Println(len(s0))
	fmt.Println(len(s1))

	fmt.Println("===================================================")

	//拼接，可以用+号。也可以用fmt.Sprintf，但需要将值赋给变量
	s2 := "理想"
	s3 := "大帅哥"
	fmt.Println(s2 + s3)
	s4 := s2 + s3
	fmt.Println(s4)
	fmt.Printf("%s%s", s2, s3)
	fmt.Println()
	s5 := fmt.Sprintf("%s%s", s2, s3)
	fmt.Println(s5)

	fmt.Println("===================================================")
	//分割 string.Split

	str5 := "\"D:\\Gostudy\\src\\github.com\\Y1gg\\day02\\08string\""
	fmt.Println(strings.Split(str5, "\\"))

	fmt.Println("===================================================")

	//包含 strings.Contains

	fmt.Println(strings.Contains(s5, s3))
	fmt.Println(strings.Contains(s5, "liu"))

	fmt.Println("===================================================")

	//前缀strings.HasPrefix和后缀strings.HasSuffix

	fmt.Println(strings.HasPrefix(s5, s2))
	fmt.Println(strings.HasSuffix(s5, "大帅哥"))

	fmt.Println("===================================================")

	//出现的位置，最先出现的位置strings.Index和最后出现的位置strings.LastIndex
	fmt.Println(strings.Index(s0, "d"))
	fmt.Println(strings.LastIndex(s0, "xyz"))

	fmt.Println("===================================================")

	//拼接 string.Join()
	str6 := strings.Split(str5, "\\")
	fmt.Println(strings.Join(str6, "|||||"))

}
