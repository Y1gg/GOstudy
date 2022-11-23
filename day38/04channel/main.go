package main

import (
	"fmt"
	"math/rand"
	"time"
)

func f1() {
	ch := make(chan int, 1)
	ch <- 100
	<-ch
	x, ok := <-ch //再取阻塞
	fmt.Println(x, ok)
	//fatal error: all goroutines are asleep - deadlock!   死锁
}

func sendNum(ch chan<- int) {
	for {
		num := rand.Intn(10)
		ch <- num
		time.Sleep(5 * time.Second)
	}
}
func main() {
	//f1()
	ch := make(chan int, 1)
	go sendNum(ch)
	for {
		x, ok := <-ch
		fmt.Println(x, ok)
		time.Sleep(time.Second)
	}
}
