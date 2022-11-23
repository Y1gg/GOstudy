package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("main函数")
	for i := 0; i < 100; i++ {
		go func(i int) {
			fmt.Println(i)
		}(i)
	}

	fmt.Println("main函数")
	time.Sleep(time.Second)

}
