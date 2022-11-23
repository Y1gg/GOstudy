/*
func SetOutput
func SetOutput(w io.Writer)
SetOutput设置标准logger的输出目的地，默认是标准错误输出。

var (

	Stdin  = NewFile(uintptr(syscall.Stdin), "/dev/stdin")
	Stdout = NewFile(uintptr(syscall.Stdout), "/dev/stdout")
	Stderr = NewFile(uintptr(syscall.Stderr), "/dev/stderr")

)
Stdin、Stdout和Stderr是指向标准输入、标准输出、标准错误输出的文件描述符。
*/
package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

func main() {
	fileObj, err := os.OpenFile("./xx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Printf("open file failed,err:%v\n", err)
		return
	}
	// log.SetOutput(os.Stdout)
	log.SetOutput(fileObj)

	for {
		log.Println("这是一条测试日志")
		time.Sleep(time.Second * 3)

	}
}
