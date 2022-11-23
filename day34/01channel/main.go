package main

import (
	"fmt"
	"sync"
)

var b chan int //需要指定通道内元素的类型
var wg sync.WaitGroup

func nobufChannel() {
	fmt.Println(b)     //nil
	b = make(chan int) //不带缓冲区通道的初始化
	wg.Add(1)
	go func() {
		defer wg.Done()
		x := <-b
		fmt.Println("从后台goroutine从通道b中取到了", x)
	}()

	b <- 10
	fmt.Println("10发送到通道b中")
	wg.Wait()
}
func bufChannel() {
	fmt.Println(b)         //nil
	b = make(chan int, 16) //带缓冲区的通道的初始化

	b <- 10
	fmt.Println("10发送到通道b中")
	fmt.Println(b)
}
func main() {
	nobufChannel()
	fmt.Println("=======================================")
	bufChannel()

}
