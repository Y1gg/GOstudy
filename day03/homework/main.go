package main

import "fmt"

func main() {

	for i := 9; i >= 1; i-- {
		for j := 1; j < 10; j++ {
			if j <= i {
				fmt.Printf("%d*%d=%d\t", j, i, (i * j))
			}

		}
		fmt.Println()
	}

	fmt.Println("\n==================================")
	s1 := "hello"
	s2 := "京东Java2班"

	for _, v := range s1 {
		fmt.Printf("%c %T \t", v, v)
	}
	fmt.Println("\n==================================")
	for _, v := range s2 {
		fmt.Printf("%c %T \t", v, v)
	}
	fmt.Println("\n==================================")
}
