package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	fmt.Printf("current time:%v\n", now)

	year := now.Year()
	month := now.Month()
	day := now.Day()
	hour := now.Hour()
	minute := now.Minute()
	second := now.Second()
	fmt.Println(year)
	fmt.Println(month)
	fmt.Println(day)
	fmt.Println(hour)
	fmt.Println(minute)
	fmt.Println(second)
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Println("\n====================================")
	fmt.Println(now.Date())
}
