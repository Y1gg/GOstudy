package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	rwlock sync.RWMutex
	wg     sync.WaitGroup
)

func writeLock() {
	rwlock.RLock()
	x++
	rwlock.RUnlock()
	wg.Done()
}
func readLock() {
	rwlock.Lock()
	fmt.Println(x)
	rwlock.Unlock()
	wg.Done()
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go writeLock()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go readLock()
		wg.Add(1)
	}
	fmt.Println(time.Since(start))
}
