package main

import "fmt"

func main() {
	ss := "黄山落叶松叶落山黄"
	f1(ss)
	s1 := "山西运煤车煤运山西"
	f1(s1)
	s2 := "1234567890987654321"
	f1(s2)
	s3 := "火力少年王力火"
	f1(s3)
	s := "Go语言编程"
	fmt.Println(s[0:7])
}
func f1(s string) {
	fmt.Println("\n====================================")
	r := make([]rune, 0, len(s))
	for _, c := range s {
		r = append(r, c)
	}

	fmt.Println("[]rune", r)

	for i := 0; i < len(r)/2; i++ {
		if r[i] != r[len(r)-1-i] {
			fmt.Println("不是回文")
			return
		}
	}
	fmt.Println("是回文")

}
