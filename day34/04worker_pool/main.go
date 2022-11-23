package main

import (
	"fmt"
	"time"
)

func worker(id int, jobs <-chan int, results chan<- int) {
	/*jobs是只读的通道，results是只写的通道*/
	for j := range jobs {
		fmt.Printf("worker:%d start job :%d\n", id, j)
		time.Sleep(time.Second)
		fmt.Printf("worker:%d end job :%d\n", id, j)
		results <- j * 2
	}
}

func main() {
	/*定义了两个int的通道，并初始化，jobs和results*/
	jobs := make(chan int, 100)
	results := make(chan int, 100)
	//开启三个goroutine（线程）
	for w := 1; w <= 3; w++ {
		/*调用woker函数*/
		go worker(w, jobs, results)
	}
	//五个任务
	for j := 1; j <= 5; j++ {
		jobs <- j
	}
	close(jobs)
	//输出结果
	for a := 1; a <= 5; a++ {
		<-results
	}
}
