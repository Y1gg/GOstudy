package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup
var notify bool

func f() {
	defer wg.Done()
	for {
		fmt.Println("陆小栎")
		time.Sleep(time.Millisecond * 500)
		if notify {
			break
		}
	}
}

func main() {
	wg.Add(1)
	go f()
	time.Sleep(time.Second * 5)
	notify = true
	wg.Wait()
}
