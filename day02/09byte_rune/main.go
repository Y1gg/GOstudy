package main

import "fmt"

//byte和rune类型
//go语言为了处理非ASCII码类型的字符，定义了新的rune类型

func main() {
	s := "流沙河hello"
	//len()求的是byte字节的数量
	fmt.Println(len(s))
	fmt.Println("===================================================")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v,", s[i])
		// fmt.Printf("%c,", s[i])
	}
	fmt.Println("\n===================================================")
	for _, c := range s {
		fmt.Printf("%c\t", c)
	}
	fmt.Println("\n===================================================")
	//字符串修改
	s2 := "白萝卜"
	s3 := []rune(s2)
	s3[0] = '红'
	fmt.Println(string(s3))
	fmt.Println("\n====================================")
	traversalString()
}

// 遍历字符串
func traversalString() {
	s := "hello沙河"
	for i := 0; i < len(s); i++ { //byte
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()
	for _, r := range s { //rune
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()
}
