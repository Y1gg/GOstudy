package main

import (
	"github/Y1gg/day23/mylogger"
)

//测试我们自己写的日志库

func main() {
	// log := mylogger.NewLog("Info")
	log := mylogger.NewFileLogger("Info", "./", "lili.log", 10*1024*1024)
	for {
		log.Debug("这是一条debug日志")
		log.Trace("这是一条trace日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条warning日志")
		id := 10010
		name := "沙河"
		log.Error("这是一条error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条fatal日志")
		// time.Sleep(time.Second * 2)
	}
}
