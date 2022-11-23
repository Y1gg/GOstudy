package main

import (
	"fmt"
	"reflect"
)

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("Type:\t%v\n", v)
}

func main() {
	var num1 = 1222
	var num2 int64 = 56
	var num3 float64 = 12
	reflectType(num1)
	reflectType(num2)
	reflectType(num3)

}
