package taillog

import (
	"fmt"
	"github.com/hpcloud/tail"
)

//专门从日志文件中收集日志的模块

var (
	tailObj *tail.Tail
)

func Init(fileName string) (err error) {
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}

	//以config的配置打开filename文件
	tailObj, err = tail.TailFile(fileName, config)
	//处理错误
	if err != nil {
		fmt.Println("taillog file failed, err:", err)
		return
	}
	return
}

func ReadChan() <-chan *tail.Line {
	return tailObj.Lines
}
