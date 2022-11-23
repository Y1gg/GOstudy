package main

import (
	"fmt"
	"github.com/hpcloud/tail"
	"time"
)

func main() {
	//指定文件名
	filename := "./my.log"
	//配置文件
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}

	//以config的配置打开filename文件
	tails, err := tail.TailFile(filename, config)
	//处理错误
	if err != nil {
		fmt.Println("taillog file failed  ,err:", err)
		return
	}

	var (
		line *tail.Line
		ok   bool
	)

	//一行一行的读取日志
	for {
		line, ok = <-tails.Lines
		if !ok {
			fmt.Printf("taillog file close reopen, filename:%s\n", tails.Filename)
			time.Sleep(time.Second)
			continue
		}
		fmt.Println("line:", line.Text)
	}
}
