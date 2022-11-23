package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

type job struct {
	value int64
}

type result struct {
	job    *job
	result int64
}

var jobChan = make(chan *job, 100)
var resultChan = make(chan *result, 100)
var wg sync.WaitGroup

func f1(zl chan<- *job) {
	//循环生成int64类型的随机数，发送到jobChan
	for {
		x := rand.Int63()
		newJob := &job{
			value: x,
		}
		zl <- newJob
		time.Sleep(time.Millisecond * 500)
	}
}

func f2(zl <-chan *job, resultChan chan<- *result) {
	//从jobChan中取出随机数计算个位数的和，将结果发送到resultChan
	for {
		job := <-zl
		sum := int64(0)
		n := job.value
		for n > 0 {
			sum += n % 10
			n /= 10
		}
		newResult := &result{
			job:    job,
			result: sum,
		}
		resultChan <- newResult
	}
}
func main() {
	wg.Add(1)
	go f1(jobChan)
	//开启24个goroutine执行f2
	wg.Add(24)
	for i := 0; i < 24; i++ {
		go f2(jobChan, resultChan)
	}
	for result := range resultChan {
		fmt.Printf("value:%v\tsum:%d\n", result.job.value, result.result)
	}
	wg.Wait()
}
