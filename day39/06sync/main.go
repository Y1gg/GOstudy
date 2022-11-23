package main

import "sync"

var m = sync.Map{}

func main() {
	wg := sync.WaitGroup{}
	for i := 0; i < 21; i++ {
		wg.Add(1)
		go func(n int) {
			//key:=
		}(i)
	}
}
