package main

import (
	"fmt"
	"strings"
)

//作业：网友写的

/*
	你有50枚金币，需要分配给以下几个人：
	Matthew,Sarah,Augustus,Heidi,Emilie,Peter,Giana,Adriano,Aaron,Elizabeth。
	分配规则如下：
	a. 名字中每包含1个'e'或'E'分1枚金币
	b. 名字中每包含1个'i'或'I'分2枚金币
	c. 名字中每包含1个'o'或'O'分3枚金币
	d: 名字中每包含1个'u'或'U'分4枚金币
	写一个程序，计算每个用户分到多少金币，以及最后剩余多少金币？
	程序结构如下，请实现 ‘dispatchCoin’ 函数
*/

var (
	coins = 50
	users = []string{
		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie",
		"Peter", "Giana", "Adriano", "Aaron", "Elizabeth",
	}
	distribution = make(map[string]int, len(users))
)

func main() {
	left := dispatchCoin()
	fmt.Println("剩下：", left)
}

// 模拟分发金币的函数、没有参数、返回值是剩余的金币数
func dispatchCoin() (left int) {
	for _, name := range users {
		//for循环遍历每个人的名字

		//调用countCoin函数，得到每个人对应的金币数
		coin := countCoin(name)

		//将其存到map中
		distribution[name] = coin

		//并打印map，打印对应的名字和金币数
		fmt.Printf("%s 的金币为：%d\n", name, coin)
	}

	result := 0
	for _, value := range distribution {
		result += value
	}
	fmt.Printf("总使用的金币： %d\n", result)
	return coins - result
}

// 传入名字计算金币数的函数
// 参数是string类型的名字
// 返回值是金币的个数，int类型
func countCoin(name string) (result int) {

	name = strings.ToUpper(name)
	//先利用string.ToUpper函数将名字统一为大写的
	//或者全部转为小写的

	//创建一个byte类型的切片来存储名字的每个字符：bs
	bs := []byte(name)
	result = 0

	//遍历bs切片的每个字符，如果符合，+
	for _, value := range bs {
		if value == 'E' {
			result++
		} else if value == 'I' {
			result += 2
		} else if value == 'O' {
			result += 3
		} else if value == 'U' {
			result += 4
		}
	}
	return

}

// var (
// 	coins = 50
// 	users = []string{
// 		"Matthew", "Sarah", "Augustus", "Heidi", "Emilie", "Peter",
// 		"Giana", "Adriano", "Aaron", "Elizabeth",
// 	}
// 	distribution = make(map[string]int, len(users))
// )

// func main() {
// 	for i := 0; i < len(users); i++ {
// 		name := []string{users[i]}
// 		count := 0
// 		for j := 0; j < len(name); j++ {
// 			if name[j] == 'e' {
// 				count++
// 			}
// 		}
// 	}
// 	// left := dispatchCoin()
// 	// fmt.Println("剩下：", left)
// }

// //遍历每个人的名字
// //遍历名字的每个个字母
// //用变量统计遇到eiou的次数
// //进行乘法计算，
