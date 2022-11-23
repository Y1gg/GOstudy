package main

import (
	"fmt"
	"strings"
	"unicode"
)

func main() {
	//多维数组，只有最外层的定义时可以使用...
	//例如：二维数组
	var a1 = [...][2]int{
		{1, 2},
		{3, 4},
		{5, 6},
	}
	fmt.Println(a1)
	fmt.Println("\n====================================")

	var arr = [3]int{1, 2, 3}
	fmt.Println(arr)
	f1(arr)
	fmt.Println(arr)

	fmt.Println("\n====================================")

	var s1 = []int{1, 2, 34, 6}
	fmt.Println(s1)
	s3 := make([]int, 12, 20)
	fmt.Println(s3)
	fmt.Println("\n====================================")
	var s2 []int
	s2 = append(s2, 1)
	fmt.Println(s2)
	fmt.Println(s2[0])
	s2[0] = 123
	fmt.Println(s2)
	fmt.Println("\n====================================")

	//go里面的指针只能读取不能修改，不能修改指针变量指向的地址
	addr := "武清天狮"
	addrP := &addr
	fmt.Println(addrP)
	fmt.Printf("%T\n", addrP)
	addrV := *addrP
	fmt.Print(addrV)
	fmt.Printf("%T\n", addrV)
	fmt.Println("\n====================================")

	f2()
	f3()
}

func f1(a [3]int) {
	a[1] = 100
}

func f2() {
	//map
	var m1 map[string]int
	fmt.Println(m1 == nil)
	fmt.Println(m1)

	//无序
	m1 = make(map[string]int, 10)
	m1["津南区"] = 1
	m1["西青区"] = 2
	m1["武清区"] = 3
	fmt.Println(m1)

	fmt.Println(m1["武清区"])
	fmt.Println(m1["天津市"])

	//如果返回的是布尔值，我们通常用ok接收它
	num, ok := m1["天津市"]
	if ok {
		fmt.Println("没有天津市")

	} else {
		fmt.Println("天津市的代号 是", num)
	}

	fmt.Println("\n====================================")

}

func f3() {
	//判断一串字符中汉字的数量
	//难点：判断一个字符是汉字
	//1、依次拿到每个汉字
	//2、判断该字符是不是汉字
	//3、把汉字得到的次数累加

	str := "小时不识月my name is 陆小栎 ご飯食べた？Tu as mangé?"
	count := 0
	//1、依次拿到每个汉字
	for _, c := range str {
		//2、判断该字符是不是汉字
		if unicode.Is(unicode.Han, c) {
			count++
		}
	}
	//3、把汉字得到的次数累加
	fmt.Print(count)

	fmt.Println("\n====================================")

	//2.how do you do 单词出现的 次数
	s2 := "how do you do ? how do you do"
	//2.1、把字符串按照空格切割，得到切片
	s3 := strings.Split(s2, " ")
	//2.2、遍历切片存储到一个map
	m1 := make(map[string]int, 20)
	for _, w := range s3 {
		//1、如果原来map中不存在w这个key，那么出现次数=1
		if _, ok := m1[w]; !ok {
			m1[w] = 1
		} else {
			m1[w]++
		}
		//2、如果map中存在 w这个key，那么出现次数+1

	}
	//2.3、累加出现的次数
	for key, value := range m1 {
		fmt.Println(key, value)
	}

}
