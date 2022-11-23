package main

import "fmt"

func main() {

	var flag bool
	for i := 1; i < 5; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			fmt.Printf("%d ,%c\n", i, j)
			if i == 3 && j == 'C' {
				flag = true
				break
			}
		}
		if flag {
			break
		}
	}
	fmt.Println("over")
	fmt.Println("===========================")
	for i := 1; i < 5; i++ {
		for j := 'A'; j <= 'Z'; j++ {
			fmt.Printf("%d ,%c\n", i, j)
			if i == 2 && j == 'B' {
				goto XX
			}
		}
	}
XX:
	fmt.Println("over")
}
