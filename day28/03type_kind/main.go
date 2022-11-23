package main

import (
	"fmt"
	"reflect"
)

type Cat struct {
}

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Type name:\t%v\nType kind:\t%v\n", v.Name(), v.Kind())
}

func main() {
	var c1 = Cat{}
	reflectType(c1)
	fmt.Println("====================================")
	var n1 rune = '刘'
	reflectType(n1)

}
