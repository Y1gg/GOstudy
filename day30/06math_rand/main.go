package main

import (
	"fmt"
	"math/rand"
)

func f() {
	for i := 0; i < 10; i++ {
		r1 := rand.Int()
		r2 := rand.Intn(10)
		fmt.Println(r1, r2)
	}
}

func main() {
	f()
}
