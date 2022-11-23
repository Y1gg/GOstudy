package main

import (
	"fmt"
	"github/Y1gg/day41/splitstring"
)

func main() {
	slice1 := splitstring.Split("bcdbefgbhijk", "b")
	fmt.Println(slice1)
	fmt.Printf("%#v\n", slice1)
}
