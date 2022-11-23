package main

import (
	"bufio"
	"fmt"
	"os"
)

//获取用户输入时有空格

func useScan() {
	var s string
	fmt.Print("请输入内容：")
	fmt.Scan(&s)
	fmt.Printf("输入的内容：%s\n", s)
}

func uesBufio() {
	var s string
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("请输入内容：")
	s, _ = reader.ReadString('\n')
	fmt.Printf("输入的内容：%s\n", s)

}
func main() {
	// useScan()
	uesBufio()
}
