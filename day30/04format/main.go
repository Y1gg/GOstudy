package main

import (
	"fmt"
	"strconv"
)

// Format将给定的类型转换为字符串
func main() {
	// FormatBool()
	// func FormatBool(b bool) string
	// 根据b的值返回”true”或”false”。
	n1 := true
	str1 := strconv.FormatBool(n1)
	fmt.Printf("str1\ttype:%T\tvalue:%v\n", str1, str1)

	// FormatInt()
	// func FormatInt(i int64, base int) string
	// 返回i的base进制的字符串表示。base 必须在2到36之间，结果中会使用小写字母’a’到’z’表示大于10的数字。
	n2 := -32
	str2 := strconv.FormatInt(int64(n2), 8)
	fmt.Printf("str2\ttype:%T\tvalue:%v\n", str2, str2)

	// FormatUint()
	// func FormatUint(i uint64, base int) string
	// 是FormatInt的无符号整数版本。
	n3 := 32
	str3 := strconv.FormatInt(int64(n3), 16)
	fmt.Printf("str3\ttype:%T\tvalue:%v\n", str3, str3)
	// FormatFloat()
	// func FormatFloat(f float64, fmt byte, prec, bitSize int) string
	// 函数将浮点数表示为字符串并返回。
	n4 := 12.34
	str4 := strconv.FormatFloat(n4, 'E', -1, 64)
	fmt.Printf("str4\ttype:%T\tvalue:%v\n", str4, str4)
}
