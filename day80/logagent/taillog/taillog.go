package taillog

import (
	"fmt"
	"github.com/Y1gg/day80/logagent/kafka"
	"github.com/hpcloud/tail"
)

//专门从日志文件收集日志的模块

// TailTask :一个日志收集的任务
type TailTask struct {
	path     string
	topic    string
	instance *tail.Tail
}

func NewTailTask(path, topic string) (tailObj *TailTask) {
	tailObj = &TailTask{
		path:  path,
		topic: topic,
	}
	tailObj.Init() //根据路径去打开对应的日志
	return

}

func (t *TailTask) Init() {
	//配置文件
	config := tail.Config{
		ReOpen:    true,                                 //重新打开
		Follow:    true,                                 //是否跟随
		Location:  &tail.SeekInfo{Offset: 0, Whence: 2}, //从文件的哪个地方开始读
		MustExist: false,                                //文件不存在不报错
		Poll:      true,
	}
	var err error
	t.instance, err = tail.TailFile(t.path, config)
	if err != nil {
		fmt.Println("tail file failed ,err:", err)
	}
}

func (t *TailTask) run() {
	for {
		select {
		case line := <-t.instance.Lines: //从tailObj的通道中一行一行的读取日志数据
			//发往kafka
			kafka.SendToKafka(t.topic, line.Text)
		}
	}
}
