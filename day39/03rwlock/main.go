package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	x      = 0
	wg     sync.WaitGroup
	lock   sync.Mutex
	rwlock sync.RWMutex
)

func read() {
	defer wg.Done()
	//lock.Lock()
	rwlock.Lock()
	fmt.Println(x)
	//lock.Unlock()
	rwlock.Unlock()
	time.Sleep(time.Millisecond)
}
func write() {
	defer wg.Done()
	//lock.Lock()
	rwlock.RLock()
	x++
	rwlock.RUnlock()
	//lock.Unlock()
	time.Sleep(time.Millisecond * 10)
}
func main() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		go write()
		wg.Add(1)
	}
	for i := 0; i < 1000; i++ {
		go read()
		wg.Add(1)
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(start))
}
