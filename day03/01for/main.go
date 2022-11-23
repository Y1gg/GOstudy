package main

//break和continue
//switch

import "fmt"

func main() {

	switch num := 0; num {
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	case 3:
		fmt.Println(3)
	case 4:
		fmt.Println(4)
	case 5:
		fmt.Println(5)
	default:
		fmt.Println(666)
	}

	n := 7
	switch n {
	case 1, 3, 5, 7, 9:
		fmt.Println("奇数")
	case 2, 4, 6, 8, 10:
		fmt.Println("偶数")
	}

	age := 12
	switch {
	case age%2 == 0:
		fmt.Print("偶数")

	case age%2 == 1:
		fmt.Println("奇数")
	}
}
